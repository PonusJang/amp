package main

import (
	_ "amp/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"
	"amp/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.New())
}
