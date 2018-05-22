package main

import (
	"os"

	"github.com/liipx/gdict/engien"
)

func main() {
	if len(os.Args) > 0 {
		engien.NewYoudao(os.Args[1]).Query()
	}
}
