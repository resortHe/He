package pwd

import (
	"golang.org/x/crypto/bcrypt"
	"log"
)

func HashPwd(pwd string) string {
	password, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	return string(password)
}

// CheckPwd hashpwd hash后的密码 pwd输入的密码
func CheckPwd(hashPwd string, pwd string) bool {
	bytes := []byte(hashPwd)
	err := bcrypt.CompareHashAndPassword(bytes, []byte(pwd))
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
