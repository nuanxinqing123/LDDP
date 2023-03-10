package main

import (
	"fmt"
	"regexp"
)

func main() {
	fmt.Println(VerifyEmailFormat("123@123.com"))
}

func VerifyEmailFormat(email string) bool {
	pattern := `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*` //匹配电子邮箱
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(email)
}
