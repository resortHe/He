package flag

import (
	sys_flag "flag"
	"github.com/fatih/structs"
)

type Option struct {
	DB   bool   `structs:""`
	User string //-u  admin  -u user
}

func Parse() Option {
	db := sys_flag.Bool("db", false, "初始化数据库")
	user := sys_flag.String("u", "", "创建用户")
	sys_flag.Parse()
	return Option{
		DB:   *db,
		User: *user,
	}
}
func IsWebStop(option Option) (f bool) {
	m := structs.Map(&option)
	for _, v := range m {

		switch val := v.(type) {
		case string:
			if val != "" {
				f = true
			}
		case bool:
			if val == true {
				f = true
			}
		}
	}
	return f
}
func SwitchOption(option Option) {
	if option.DB {
		Makemigrations()
		return
	}
	if option.User == "admin" || option.User == "user" {
		CreateUser(option.User)
		return
	}
	//sys_flag.Usage()

}
