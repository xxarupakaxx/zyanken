package main

import (
	"github.com/xxarupakaxx/zyanken/client"
	"os"
)

func main() {
	os.Exit(client.NewZyanken().Run())
}
