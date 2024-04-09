package cmd

import (
	"context"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/frame/g"

	"goframe-shop/internal/controller"
	"goframe-shop/internal/service"

	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"

	_ "goframe-shop/internal/logic"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			s.Group("/", func(group *ghttp.RouterGroup) {
				// group.Middleware(ghttp.MiddlewareHandlerResponse) //
				group.Middleware(
					service.Middleware().ResponseHandler,
					service.Middleware().Ctx)
				group.Bind(
					controller.Rotation,
					controller.Position,
					controller.Admin.Create,
					controller.Admin.Update,
					controller.Admin.Delete,
					controller.Admin.List,
					controller.Login,
				)

				group.Group("/", func(group *ghttp.RouterGroup) {
					group.Middleware(service.Middleware().Auth)
					group.ALLMap(g.Map{
						"/backend/admin/info": controller.Admin.Info,
					})
				})
			})
			s.Run()
			return nil
		},
	}
)
