package file

import (
	"context"
	"goframe-shop/internal/consts"
	"goframe-shop/internal/dao"
	"goframe-shop/internal/model"
	"goframe-shop/internal/model/entity"
	"goframe-shop/internal/service"
	"time"

	"github.com/gogf/gf/util/gconv"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/gtime"
)

type sFile struct {
}

func init() {
	//todo
	service.RegisterFile(New())
}

func New() *sFile {
	return &sFile{}
}

// 1. 定义图片上传位置
// 2. 校验 上传文件是否正确，安全性校验，1分钟只能上传10次
// 3. 定义年月日
// 4. 入库
// 5. 返回数据
func (s *sFile) UploadFile(ctx context.Context, in model.FileUploadInput) (out *model.FileUploadOutput, err error) {
	uploadFile := g.Cfg().MustGet(ctx, "upload.path").String()

	if uploadFile == "" {
		return nil, gerror.New("读取配置文件失败 上传路径不存在")
	}

	if in.Name != "" {
		in.File.Filename = in.Name
	}

	count, err := dao.FileInfo.Ctx(ctx).
		Where(dao.FileInfo.Columns().UserId, gconv.Int(ctx.Value(consts.CtxAdminId))).
		WhereGTE(dao.FileInfo.Columns().CreatedAt, gtime.Now().Add(time.Minute)).Count()
	if count >= consts.FileMaxUploadCountMinute {
		return nil, gerror.Newf("上传过于频繁一分钟只能上传{}次", consts.FileMaxUploadCountMinute)
	}

	dateDirName := gtime.Now().Format("Ymd")
	fileName, err := in.File.Save(gfile.Join(uploadFile, dateDirName), in.RandomName)
	if err != nil {
		return nil, err
	}

	//入库
	data := entity.FileInfo{
		Name:   fileName,
		Src:    gfile.Join(uploadFile, dateDirName, fileName),
		Url:    "/upload/" + dateDirName + "/" + fileName,
		UserId: gconv.Int(ctx.Value(consts.CtxAdminId)),
	}

	id, err := dao.FileInfo.Ctx(ctx).Data(data).OmitEmpty().InsertAndGetId()
	if err != nil {
		return nil, err
	}
	return &model.FileUploadOutput{
		Id:   uint(id),
		Name: data.Name,
		Src:  data.Src,
		Url:  data.Url,
	}, nil
}
