package main

import "github.com/xxarupakaxx/zyanken/client"

func main() {
	err := client.NewZyanken().Run()
	if err != nil {
		panic(err)
	}
}
