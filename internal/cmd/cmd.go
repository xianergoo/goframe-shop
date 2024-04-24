package cmd

import (
	"context"
	"goframe-shop/internal/consts"

	"github.com/goflyfox/gtoken/gtoken"
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/frame/g"

	"goframe-shop/internal/controller"
	_ "goframe-shop/internal/logic"
	"goframe-shop/internal/service"

	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
)

var gfToken *gtoken.GfToken
var gfAdminToken *gtoken.GfToken

var (
	Main = gcmd.Command{
		Name:  consts.ProjectName,
		Usage: consts.ProjectUsage,
		Brief: consts.ProjectBrief,
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()

			gfAdminToken, err = StartBackendGToken()
			if err != nil {
				return err
			}

			s.Group("/backend", func(group *ghttp.RouterGroup) {
				group.Middleware(
					service.Middleware().CORS,
					service.Middleware().ResponseHandler,
					service.Middleware().Ctx)

				// 不需要登陆的路由组
				group.Bind(
					controller.Rotation,
					controller.Position,
					controller.Admin.Create,
					controller.Admin.Update,
					controller.Admin.Delete,
					controller.Admin.List,
					controller.Login.RefreshToken,
					controller.Role,
					controller.Permission,
				)

				//需要登陆的路由组
				group.Group("/", func(group *ghttp.RouterGroup) {
					err := gfAdminToken.Middleware(ctx, group)

					if err != nil {
						panic(err)
					}
					group.ALLMap(g.Map{
						"/admin/info": controller.Admin.Info,
					})

					group.Bind(
						controller.File,     //文件上传
						controller.Upload,   //文件上传云平台
						controller.Category, //商品分类
						controller.Coupon,
						controller.UserCoupon,
					)
				})
			})

			s.Group("/frontend", func(group *ghttp.RouterGroup) {
			})
			s.Run()
			return nil
		},
	}
)
