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
		Te: PBTe(p.Te),
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

func PBResult(result game.Result) pb.Result {
	switch result {
	case game.Draw:
		return pb.Result_DRAW
	case game.Lose:
		return pb.Result_LOSE
	case game.Win:
		return pb.Result_WIN
	default:
		return 3
	}
}
