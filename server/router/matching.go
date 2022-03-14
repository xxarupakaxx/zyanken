package router

import (
	"context"
	"fmt"
	"github.com/xxarupakaxx/zyanken/game"
	pb "github.com/xxarupakaxx/zyanken/gen/proto"
	"github.com/xxarupakaxx/zyanken/util"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"sync"
	"time"
)

type MatchingHandler struct {
	sync.RWMutex
	Rooms       map[int]*game.Room
	maxPlayerID int
}

func (m *MatchingHandler) JoinRoom(request *pb.JoinRoomRequest, stream pb.MatchingService_JoinRoomServer) error {
	ctx, cancel := context.WithTimeout(stream.Context(), time.Minute*2)
	defer cancel()

	m.Lock()

	m.maxPlayerID++
	me := &game.Player{ID: m.maxPlayerID}

	for _, room := range m.Rooms {
		if room.Guest == nil {
			room.Guest = me
			err := stream.Send(&pb.JoinRoomResponse{
				Room:   util.PBRoom(room),
				Me:     util.PBPlayer(room.Guest),
				Status: pb.JoinRoomResponse_MATCHED,
			})
			if err != nil {
				return err
			}

			m.Unlock()
			fmt.Printf("マッチしました roomID = %v \n", room.ID)

			return nil
		}
	}

	room := &game.Room{
		ID:   len(m.Rooms) + 1,
		Host: me,
	}
	m.Rooms[room.ID] = room
	m.Unlock()

	err := stream.Send(&pb.JoinRoomResponse{
		Room:   util.PBRoom(room),
		Status: pb.JoinRoomResponse_WAITING,
	})
	if err != nil {
		return err
	}

	ch := make(chan int)
	go func(ch chan<- int) {
		for true {
			m.RLock()
			guest := room.Guest
			m.RUnlock()
			if guest != nil {
				err = stream.Send(&pb.JoinRoomResponse{
					Room:   util.PBRoom(room),
					Me:     util.PBPlayer(room.Host),
					Status: pb.JoinRoomResponse_MATCHED,
				})
				if err != nil {
					return
				}
				ch <- 0
				break
			}
			time.Sleep(time.Second * 1)

			select {
			case <-ctx.Done():
				return
			default:

			}
		}
	}(ch)

	select {
	case <-ch:
	case <-ctx.Done():
		return status.Errorf(codes.DeadlineExceeded, "マッチングできませんでした")
	}

	return nil
}

func (m *MatchingHandler) mustEmbedUnimplementedMatchingServiceServer() {
	panic("implement me")
}
