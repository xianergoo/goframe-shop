package upload

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"os"
	"path"
)

/*
1. 获取本地文件
2. 指定上传目录
3. 上传服务器临时目录
4. 上传到云平台
4. 删除服务器临时文件
*/

func UploadImgToCloud(ctx context.Context, file *ghttp.UploadFile) (url string, err error) {
	dirPath := "/tmp/"
	fileName, err := file.Save(dirPath, true)
	if err != nil {
		return "", err
	}

	localFile := path.Join(dirPath, fileName)
	g.Log().Debug(ctx, "LocalFile", localFile)

	//获得七牛云配置文件参数
	bucket := g.Cfg().MustGet(ctx, "qiniu.bucket").String()
	accessKey := g.Cfg().MustGet(ctx, "qiniu.accessKey").String()
	secretKey := g.Cfg().MustGet(ctx, "qiniu.secretKey").String()
	//对接七牛云的SDK
	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}
	//生成token
	mac := qbox.NewMac(accessKey, secretKey)
	//上传token
	upToken := putPolicy.UploadToken(mac)
	//七牛上传的配置
	cfg := storage.Config{}
	cfg.Zone = &storage.ZoneHuabei
	cfg.UseHTTPS = true
	cfg.UseCdnDomains = false //根据自己需求去灵活配置
	//构建表单上传对象
	formUploader := storage.NewFormUploader(&cfg)
	//上传结果的结构体
	ret := storage.PutRet{}
	//可选配置
	putExtra := storage.PutExtra{
		Params: map[string]string{},
	}
	//七牛云表单上传
	key := fileName
	err = formUploader.PutFile(ctx, &ret, upToken, key, localFile, &putExtra)
	g.Dump(err)
	if err != nil {
		return "", err
	}
	fmt.Println(ret.Key, ret.Hash, ret.PersistentID)
	//删除本地临时文件
	err = os.RemoveAll(localFile)
	if err != nil {
		return "", err
	}
	//返回数据
	url = g.Cfg().MustGet(ctx, "qiniu.url").String() + ret.Key
	return

}
