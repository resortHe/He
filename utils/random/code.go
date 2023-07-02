package random

import (
	"fmt"
	"math/rand"
	"time"
)

func Code(length int) string {
	rand.Seed(time.Now().UnixNano())
	code := ""
	for i := 0; i < length; i++ {
		digit := rand.Intn(10) // 生成0-9之间的随机数
		code += fmt.Sprint(digit)
	}

	return code

}
