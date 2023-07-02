package config

type QiNiu struct {
	Enable    bool    `json:"enable" yaml:"enable"`
	AccessKey string  `json:"access_key" yaml:"access_key"`
	SecretKey string  `json:"secret_key" yaml:"secret_key"`
	Bucket    string  `json:"bucket" yaml:"bucket"` //存储桶的名字
	CND       string  `json:"cnd" yaml:"cnd"`       //访问图片的地址前缀
	Zone      string  `json:"zone" yaml:"zone"`     //存储的地区
	Prefix    string  `json:"prefix" yaml:"prefix"` //前缀
	Size      float64 `json:"size" yaml:"size"`     //图片大小限制 单位MB
}
