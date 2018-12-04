package utils

import (
	"math/rand"
	"os"
	"regexp"
	"strings"
	"time"
)

func GenerateRandomNumber(start int, end int) uint {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	num := r.Intn((end - start)) + start
	return uint(num)
}

func ConfigTemplateRender(source string) string {
	//去除注释的部分
	source = regexp.MustCompile(`/\*.*\*/`).ReplaceAllString(source, "")
	//优先从环境变量中读取
	reg := regexp.MustCompile(`[\\]?\${([^}]+)}`)
	strs := reg.FindAllString(source, len(source)+1)
	println(len(strs))
	source = reg.ReplaceAllStringFunc(source, func(src string) string {
		default_src := ""
		src = strings.TrimSpace(src)
		src = strings.TrimPrefix(src, "${")
		src = strings.TrimSuffix(src, "}")
		if strings.Contains(src, "|") {
			s := strings.Split(src, "|")
			default_src = strings.TrimSpace(s[1])
			src = strings.TrimSpace(s[0])
		}
		temp, exist := os.LookupEnv(src)
		if !exist {
			temp = default_src
		}
		return temp
	})
	return source

}
