package backend

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

type FileUploadReq struct {
	g.Meta `path:"/upload/img" method:"post" mime:"mutipart/form-data" tags:"tools" dc:"上传文件"`
	File   *ghttp.UploadFile `json:"file" type:"file" dc:"选择上传文件"`
}

type FileUploadRes struct {
	Name string `json:"name" dc:"文件名称"`
	Url  string `json:"url" dc:"图片描述"`
}
