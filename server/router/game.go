package router

import (
	"fmt"
	"github.com/xxarupakaxx/zyanken/game"
	pb "github.com/xxarupakaxx/zyanken/gen/proto"
	"github.com/xxarupakaxx/zyanken/util"
	"sync"
)

type GameHandler struct {
	sync.RWMutex
	games  map[int]*game.Game
	client map[int][]pb.ZyankenService_PlayServer
}

func NewGameHandler() *GameHandler {
	return &GameHandler{
		games:  make(map[int]*game.Game),
		client: make(map[int][]pb.ZyankenService_PlayServer),
	}
}

func (g *GameHandler) Play(server pb.ZyankenService_PlayServer) error {
	for true {
		request, err := server.Recv()
		if err != nil {
			return err
		}

		roomID := request.GetRoomId()
		player := util.Player(request.GetPlayer())
		switch request.GetAction().(type) {
		case *pb.PlayerRequest_Start:
			err = g.start(server, roomID, player)
			if err != nil {
				return err
			}
		case *pb.PlayerRequest_Zyanken:
			err = g.zyanken(roomID, player)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (g *GameHandler) start(stream pb.ZyankenService_PlayServer, id int32, player *game.Player) error {
	g.Lock()
	defer g.Unlock()

	ga := g.games[int(id)]
	if ga == nil {
		ga = game.NewGame()
		g.games[int(id)] = ga
		g.client[int(id)] = make([]pb.ZyankenService_PlayServer, 0, 2)
	}

	g.client[int(id)] = append(g.client[int(id)], stream)

	if len(g.client[int(id)]) == 2 {
		for _, server := range g.client[int(id)] {
			err := server.Send(&pb.PlayerResponse{Event: &pb.PlayerResponse_Ready{Ready: &pb.PlayerResponse_ReadyEvent{}}})
			if err != nil {
				return err
			}
		}

		fmt.Printf("じゃんけんが始まりました roomID = %v\n", id)
	} else {
		err := stream.Send(&pb.PlayerResponse{Event: &pb.PlayerResponse_Waiting{Waiting: &pb.PlayerResponse_WaitingEvent{}}})
		if err != nil {
			return err
		}
	}

	return nil
}

func (g *GameHandler) zyanken(id int32, player *game.Player) error {
	g.Lock()
	defer g.Unlock()

	ga := g.games[int(id)]

	winLose := ga.DecideWinLose(player.Te)
	for _, server := range g.client[int(id)] {
		err := server.Send(&pb.PlayerResponse{Event: &pb.PlayerResponse_Zyanken{Zyanken: &pb.PlayerResponse_ZyankenEvent{
			Player: util.PBPlayer(player),
		}}})
		if err != nil {
			return err
		}

		if winLose {
			err = server.Send(&pb.PlayerResponse{Event: &pb.PlayerResponse_Finished{
				Finished: &pb.PlayerResponse_FinishedEvent{
					Result: util.PBResult(ga.Winner(player.Te)),
				},
			}})

			delete(g.client, int(id))
			if err != nil {
				return err
			}
		}
	}

	return nil
}
