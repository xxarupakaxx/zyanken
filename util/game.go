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
	return &game.Player{
		ID: int(p.GetId()),
		Te: Te(p.GetTe()),
	}
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

func Result(p pb.Result) game.Result {
	switch p {
	case pb.Result_DRAW:
		return game.Draw
	case pb.Result_LOSE:
		return game.Lose
	case pb.Result_WIN:
		return game.Win
	}

	return 3
}
