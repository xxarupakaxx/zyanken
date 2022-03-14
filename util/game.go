package util

import (
	"github.com/xxarupakaxx/zyanken/game"
	pb "github.com/xxarupakaxx/zyanken/gen/proto"
)

func Room(r *pb.Room) *game.Room {
	return &game.Room{
		ID:    int(r.GetId()),
		Host:  Player(r.GetHost()),
		Guest: Player(r.GetHost()),
	}
}

func Player(p *pb.Player) *game.Player {
	return &game.Player{ID: int(p.GetId())}
}

func Te(p pb.Te) game.Zyanken {
	switch p {
	case pb.Te_Gu:
		return game.Gu
	case pb.Te_Choki:
		return game.Choki
	case pb.Te_Pa:
		return game.Pa
	}

	return 0
}
