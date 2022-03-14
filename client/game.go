package client

import (
	"bufio"
	"context"
	"fmt"
	"github.com/xxarupakaxx/zyanken/game"
	pb "github.com/xxarupakaxx/zyanken/gen/proto"
	"github.com/xxarupakaxx/zyanken/util"
	"google.golang.org/grpc"
	"os"
	"sync"
	"time"
)

type Zyanken struct {
	sync.RWMutex
	started  bool
	finished bool
	me       *game.Player
	room     *game.Room
	game     *game.Game
}

func NewZyanken() *Zyanken {
	return &Zyanken{}
}

func (z *Zyanken) Run() error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	conn, err := grpc.Dial(":50051", grpc.WithInsecure())
	if err != nil {
		return fmt.Errorf("failed to connect to grpc server :%w", err)
	}

	defer conn.Close()

	if err = z.matching(ctx, pb.NewMatchingServiceClient(conn)); err != nil {
		return err
	}

	z.game = game.NewGame()

	return z.play(ctx, pb.NewZyankenServiceClient(conn))
}

func (z *Zyanken) matching(ctx context.Context, client pb.MatchingServiceClient) error {
	stream, err := client.JoinRoom(ctx, &pb.JoinRoomRequest{})
	if err != nil {
		return err
	}

	defer stream.CloseSend()

	fmt.Println("マッチング相手を探しています...")

	for true {
		response, err := stream.Recv()
		if err != nil {
			return err
		}

		if response.GetStatus() == pb.JoinRoomResponse_MATCHED {
			z.room = util.Room(response.GetRoom())
			z.me = util.Player(response.GetMe())
			fmt.Printf("マッチしたルームのIDは %d\n", response.GetRoom().GetId())
			return nil
		} else if response.GetStatus() == pb.JoinRoomResponse_WAITING {
			fmt.Printf("マッチングを待っています...")
		}
	}

	return nil
}

func (z *Zyanken) play(ctx context.Context, client pb.ZyankenServiceClient) error {
	c, cancel := context.WithCancel(ctx)
	defer cancel()

	stream, err := client.Play(ctx)
	if err != nil {
		return err
	}

	defer stream.CloseSend()

	go func() {
		err = z.send(c, stream)
		if err != nil {
			cancel()
		}
	}()

	err = z.receive(c, stream)
	if err != nil {
		cancel()
		return err
	}

	return nil
}

func (z *Zyanken) send(ctx context.Context, stream pb.ZyankenService_PlayClient) error {
	for true {
		z.RLock()

		if z.finished {
			z.RUnlock()
			return nil
		} else if !z.started {
			err := stream.Send(&pb.PlayerRequest{
				RoomId: int32(z.room.ID),
				Player: util.PBPlayer(z.me),
				Action: &pb.PlayerRequest_Start{Start: &pb.PlayerRequest_StartAction{}},
			})
			z.RUnlock()
			if err != nil {
				return err
			}

			for true {
				z.RLock()
				if z.started {
					z.RUnlock()
					fmt.Println("対戦相手が見つかりました")
					break
				}
				z.RUnlock()
				fmt.Println("対戦相手が見つかりません")

				time.Sleep(time.Second * 1)
			}
		} else {
			z.RUnlock()
			fmt.Println("どの手をだす？0,1,2の中で入力してね。0:グー,1:チョキ,2:パー")
			stdin := bufio.NewScanner(os.Stdin)
			stdin.Scan()

			text := stdin.Text()
			te, err := parseInput(text)
			if err != nil {
				fmt.Println(err)
				continue
			}

			go func() {
				err = stream.Send(&pb.PlayerRequest{
					RoomId: int32(z.room.ID),
					Player: util.PBPlayer(z.me),
					Action: &pb.PlayerRequest_Zyanken{
						Zyanken: &pb.PlayerRequest_ZyankenAction{
							Zyanken: &pb.Zyanken{
								Te: util.PBTe(te),
							},
						},
					},
				})
				if err != nil {
					fmt.Println(err)
				}
			}()

			ch := make(chan int)
			go func(ch chan int) {
				fmt.Println("")
				for i := 0; i < 5; i++ {
					fmt.Printf("%d秒間止まります \n", 5-i)
					time.Sleep(time.Second * 1)
				}
				fmt.Println("")
				ch <- 0
			}(ch)

			<-ch
		}

		select {
		case <-ctx.Done():
			return nil
		default:

		}
	}
	return nil
}

func (z *Zyanken) receive(ctx context.Context, stream pb.ZyankenService_PlayClient) error {
	for true {
		response, err := stream.Recv()
		if err != nil {
			return err
		}

		z.Lock()
		switch response.GetEvent().(type) {
		case *pb.PlayerResponse_Waiting:
		case *pb.PlayerResponse_Ready:
			z.started = true
		case *pb.PlayerResponse_Zyanken:
			aite := util.Te(response.GetZyanken().GetPlayer().GetTe())
			winLose := z.game.DecideWinLose(aite)
			if winLose {
				z.finished = true
			}
		case *pb.PlayerResponse_Finished:
			winner := util.Result(response.GetFinished().GetResult())
			fmt.Println("")
			if winner == game.Win {
				fmt.Println("あなたの勝ち！")
			} else if winner == game.Lose {
				fmt.Println("あなたの負け!")
			}

			z.Unlock()
			return nil
		}
		z.Unlock()

		select {
		case <-ctx.Done():
			return nil
		default:
		}

	}

	return nil
}

func parseInput(text string) (game.Zyanken, error) {
	switch text {
	case "0":
		return game.Gu, nil
	case "1":
		return game.Choki, nil
	case "2":
		return game.Pa, nil
	default:
		return 3, fmt.Errorf("入力が不正です")
	}
}
