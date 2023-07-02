package qinniu

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"go_vue/config"
	"go_vue/global"
	"time"
)

func getToken(q config.QiNiu) string {
	accessKey := q.AccessKey
	secretKey := q.SecretKey
	bucket := q.Bucket
	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}
	mac := qbox.NewMac(accessKey, secretKey)
	uploadToken := putPolicy.UploadToken(mac)
	return uploadToken
}
func getCfg(q config.QiNiu) storage.Config {
	cfg := storage.Config{}
	//空间对应的机房
	zone, _ := storage.GetRegionByID(storage.RegionID(q.Zone))
	cfg.Zone = &zone
	//是否使用https域名
	cfg.UseHTTPS = false
	//上传是否启用CDN加速
	cfg.UseCdnDomains = false
	return cfg
}
func UploadImage(data []byte, imageName string, prefix string) (filePath string, err error) {
	if !global.Config.QiNiu.Enable {
		return "", errors.New("请启用七牛云上传")
	}
	q := global.Config.QiNiu
	if q.AccessKey == "" || q.SecretKey == "" {
		return "", errors.New("请配置AccessKey以及SecretKey")
	}
	if float64(len(data))/1024/1024 > q.Size {
		return "", errors.New("文件超过预定大小")
	}
	token := getToken(q)
	cfg := getCfg(q)
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}
	putExtra := storage.PutExtra{
		Params: map[string]string{},
	}
	dataLen := int64(len(data))
	//获取当前时间
	now := time.Now().Format("20060102150405")
	key := fmt.Sprintf("%s%s__%s", prefix, now, imageName)
	err = formUploader.Put(context.Background(), &ret, token, key, bytes.NewReader(data), dataLen, &putExtra)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s%s", q.CND, ret.Key), nil
}
