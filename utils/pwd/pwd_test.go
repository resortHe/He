package pwd

import (
	"fmt"
	"testing"
)

func TestHashPwd(t *testing.T) {
	fmt.Println(HashPwd("1234"))
}
func TestCheckPwd(t *testing.T) {
	fmt.Println(CheckPwd("$2a$04$oW1BnsLIqOGRdr2o4alQxeZ3cgCATZBKpd.YErxuQIRkKEZUDPwwe", "1234"))
}
