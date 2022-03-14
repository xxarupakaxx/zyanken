package util

import (
	"github.com/xxarupakaxx/zyanken/game"
	pb "github.com/xxarupakaxx/zyanken/gen/proto"
)

func PBRoom(r *game.Room) *pb.Room {
	return &pb.Room{
		Id:    int32(r.ID),
		Host:  PBPlayer(r.Host),
		Guest: PBPlayer(r.Guest),
	}
}

func PBPlayer(p *game.Player) *pb.Player {
	return &pb.Player{
		Id: int32(p.ID),
	}
}

func PBTe(te game.Zyanken) pb.Te {
	switch te {
	case game.Gu:
		return pb.Te_Gu
	case game.Choki:
		return pb.Te_Choki
	case game.Pa:
		return pb.Te_Pa
	}

	return 0
}
