package main

import (
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"

	_ "goframe-shop/internal/logic"

	_ "goframe-shop/internal/packed"

	_ "github.com/gogf/gf/contrib/nosql/redis/v2"

	"github.com/gogf/gf/v2/os/gctx"

	"goframe-shop/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
