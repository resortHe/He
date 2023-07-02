package desens

import "strings"

// DesensitizationEmail 脱敏成首字母+*+@
// 例如2333@qq.com => 2****@qq.com
func DesensitizationEmail(email string) string {
	elist := strings.Split(email, "@")
	if len(elist) != 2 {
		return ""
	}
	return elist[0][:1] + "****@" + elist[1]
}

// DesensitizationTel  12345678901脱敏成123****8901
func DesensitizationTel(tel string) string {
	if len(tel) != 11 {
		return ""
	}

	return tel[:3] + "****" + tel[7:]
}
