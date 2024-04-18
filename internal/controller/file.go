package controller

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"goframe-shop/api/backend"
	"goframe-shop/internal/model"
	"goframe-shop/internal/service"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
)

var File = cFile{}

type cFile struct {
}

func (c *cFile) Upload(ctx context.Context, req *backend.FileUploadReq) (res *backend.FileUploadRes, err error) {
	g.Log().Debug(ctx, "test")
	if req.File == nil {
		return nil, gerror.NewCode(gcode.CodeMissingParameter, "请上传文件")
	}

	upload, err := service.File().UploadFile(ctx, model.FileUploadInput{
		File:       req.File,
		RandomName: true,
	})
	if err != nil {
		return nil, err
	}

	return &backend.FileUploadRes{Name: upload.Name, Url: upload.Url}, nil
}
