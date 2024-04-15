package cmd

import (
	"context"
	"strconv"

	"github.com/goflyfox/gtoken/gtoken"
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/text/gstr"
	"github.com/gogf/gf/util/gconv"
	"github.com/gogf/gf/v2/frame/g"

	"goframe-shop/api/backend"
	"goframe-shop/internal/consts"
	"goframe-shop/internal/controller"
	"goframe-shop/internal/dao"
	"goframe-shop/internal/model/entity"
	"goframe-shop/internal/service"
	"goframe-shop/utility"
	"goframe-shop/utility/response"

	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"

	_ "goframe-shop/internal/logic"
)

var gfToken *gtoken.GfToken

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()

			gfToken = &gtoken.GfToken{
				ServerName:       "shop",
				CacheMode:        2, //缓存模式 1 gcache 2 gredis 3 fileCache
				LoginPath:        "/backend/login",
				LoginBeforeFunc:  loginFunc,
				LoginAfterFunc:   loginAfterFunc,
				LogoutPath:       "/backend/logout",
				AuthPaths:        g.SliceStr{"/backend/admin/info"},
				AuthExcludePaths: g.SliceStr{"/user/info", "/system/user/info"}, // 不拦截路径 /user/info,/system/user/info,/system/user,
				MultiLogin:       true,
				AuthAfterFunc:    authAfterFunc,
			}

			s.Group("/", func(group *ghttp.RouterGroup) {
				// group.Middleware(ghttp.MiddlewareHandlerResponse) //
				group.Middleware(
					service.Middleware().CORS,
					service.Middleware().ResponseHandler,
					service.Middleware().Ctx)

				//gtoken middleware
				// err := gfToken.Middleware(ctx, group)
				// if err != nil {
				// 	panic(err)
				// }

				group.Bind(
					controller.Rotation,
					controller.Position,
					controller.Admin.Create,
					controller.Admin.Update,
					controller.Admin.Delete,
					controller.Admin.List,
					controller.Login.RefreshToken,
					controller.Role,
				)

				group.Group("/", func(group *ghttp.RouterGroup) {
					// group.Middleware(service.Middleware().Auth) //for jwt
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

func loginFunc(r *ghttp.Request) (string, interface{}) {
	name := r.Get("name").String()
	password := r.Get("password").String()
	ctx := context.TODO()

	if name == "" || password == "" {
		r.Response.WriteJson(gtoken.Fail("账号或密码错误."))
		r.ExitAll()
	}

	adminInfo := entity.AdminInfo{}
	err := dao.AdminInfo.Ctx(ctx).Where("name", name).Scan(
		&adminInfo)
	if err != nil {
		r.Response.WriteJson(err)
		r.ExitAll()
	}

	if utility.EncryptPassword(password, adminInfo.UserSalt) != adminInfo.Password {
		r.Response.WriteJson(gtoken.Fail("user or password not correct."))
		r.ExitAll()
	}

	g.Log().Info(context.Background(), "test")
	// 唯一标识，扩展参数user data
	return consts.GtokenAdminPrefix + strconv.Itoa(adminInfo.Id), adminInfo
}

// 自定义的登录之后的函数
func loginAfterFunc(r *ghttp.Request, respData gtoken.Resp) {
	if !respData.Success() {
		respData.Code = 0
		r.Response.WriteJson(respData)
		return
	} else {
		respData.Code = 1
		//获得登录用户id
		userKey := respData.GetString("userKey")
		adminId := gstr.StrEx(userKey, consts.GtokenAdminPrefix)
		//根据id获得登录用户其他信息
		adminInfo := entity.AdminInfo{}
		err := dao.AdminInfo.Ctx(context.TODO()).WherePri(adminId).Scan(&adminInfo)
		if err != nil {
			return
		}
		//通过角色查询权限
		//先通过角色查询权限id
		var rolePermissionInfos []entity.RolePermissionInfo
		err = dao.RolePermissionInfo.Ctx(context.TODO()).WhereIn(dao.RolePermissionInfo.Columns().RoleId, g.Slice{adminInfo.RoleIds}).Scan(&rolePermissionInfos)
		if err != nil {
			return
		}
		permissionIds := g.Slice{}
		for _, info := range rolePermissionInfos {
			permissionIds = append(permissionIds, info.PermissionId)
		}

		var permissions []entity.PermissionInfo
		err = dao.PermissionInfo.Ctx(context.TODO()).WhereIn(dao.PermissionInfo.Columns().Id, permissionIds).Scan(&permissions)
		if err != nil {
			return
		}
		data := &backend.LoginRes{
			Type:        consts.TokenType,
			Token:       respData.GetString("token"),
			ExpireIn:    consts.GTokenExpireIn, //单位秒,
			IsAdmin:     adminInfo.IsAdmin,
			RoleIds:     adminInfo.RoleIds,
			Permissions: permissions,
		}
		response.JsonExit(r, 0, "", data)
	}
	return
}

func authAfterFunc(r *ghttp.Request, respData gtoken.Resp) {
	g.Dump("token", respData.GetString("token"))
	g.Dump("respData", respData)
	var adminInfo entity.AdminInfo
	err := gconv.Struct(respData.GetString("data"), &adminInfo)
	if err != nil {
		response.Auth(r)
		return
	}

	// if adminInfo.CreatedAt != nil {
	// 	response.AuthBlack(r)
	// 	return
	// }
	//todo 这里可以写账号前置校验、是否被拉黑、有无权限等逻辑
	r.SetCtxVar(consts.CtxAdminId, adminInfo.Id)
	r.SetCtxVar(consts.CtxAdminName, adminInfo.Name)
	r.SetCtxVar(consts.CtxAdminIsAdmin, adminInfo.IsAdmin)
	r.SetCtxVar(consts.CtxAdminRoleIds, adminInfo.RoleIds)
	r.Middleware.Next()
}
