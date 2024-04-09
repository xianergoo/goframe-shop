package login

import (
	"context"
	"goframe-shop/utility"

	"goframe-shop/internal/service"

	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/util/gutil"

	"goframe-shop/internal/dao"
	"goframe-shop/internal/model"
	"goframe-shop/internal/model/entity"
)

type sLogin struct {
	avatarUploadPath      string // 头像上传路径
	avatarUploadUrlPrefix string // 头像上传对应的URL前缀
}

func init() {
	login := New()
	// 启动时创建头像存储目录
	if !gfile.Exists(login.avatarUploadPath) {
		if err := gfile.Mkdir(login.avatarUploadPath); err != nil {
			g.Log().Fatal(gctx.New(), err)
		}
	}
	service.RegisterLogin(login)
}

func New() *sLogin {
	return &sLogin{
		avatarUploadPath:      g.Cfg().MustGet(gctx.New(), `upload.path`).String() + `/avatar`,
		avatarUploadUrlPrefix: `/upload/avatar`,
	}
}

// 获得头像上传路径
func (s *sLogin) GetAvatarUploadPath() string {
	return s.avatarUploadPath
}

// 获得头像上传对应的URL前缀
func (s *sLogin) GetAvatarUploadUrlPrefix() string {
	return s.avatarUploadUrlPrefix
}

// 执行登录
func (s *sLogin) Login(ctx context.Context, in model.UserLoginInput) error {
	adminInfo := entity.AdminInfo{}
	err := dao.AdminInfo.Ctx(ctx).Where("name", in.Name).Scan(
		&adminInfo)
	if err != nil {
		return err
	}
	gutil.Dump(utility.EncryptPassword(in.Password, adminInfo.UserSalt))
	if utility.EncryptPassword(in.Password, adminInfo.UserSalt) != adminInfo.Password {
		return gerror.New("user or password not correct")
	}

	if err := service.Session().SetUser(ctx, &adminInfo); err != nil {
		return err
	}
	// 自动更新上线
	service.BizCtx().SetUser(ctx, &model.ContextUser{
		Id:      uint(adminInfo.Id),
		Name:    adminInfo.Name,
		IsAdmin: uint(adminInfo.IsAdmin),
	})
	return nil
}

// 注销
func (s *sLogin) Logout(ctx context.Context) error {
	return service.Session().RemoveUser(ctx)
}

// 将密码按照内部算法进行加密
func (s *sLogin) EncryptPassword(passport, password string) string {
	return gmd5.MustEncrypt(passport + password)
}
