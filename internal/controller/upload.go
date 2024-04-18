package controller

import (
	"context"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"goframe-shop/api/backend"
	upload2 "goframe-shop/utility/upload"
)

var Upload = cUpload{}

type cUpload struct {
}

func (c *cUpload) Upload(ctx context.Context, req *backend.UploadImgToCloudReq) (res *backend.UploadImgToCloudRes, err error) {
	if req.File == nil {
		return nil, gerror.NewCode(gcode.CodeMissingParameter, "请上传文件")
	}

	url, err := upload2.UploadImgToCloud(ctx, req.File)
	if err != nil {
		return nil, err
	}

	return &backend.UploadImgToCloudRes{Url: url}, nil
}
