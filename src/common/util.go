package common

import (
	"strings"
	"os"
)

//the /home/root/file formart is "home" "root" "file", not include the root "/"dir
func SplistPath(fname string) []string{
	temp := strings.Split(string(fname), string(os.PathSeparator))
	parts := []string{}
	if strings.HasPrefix(fname, string(os.PathSeparator)){
		parts = append(parts, string(os.PathSeparator))
	}
	for _, v := range temp{
		if v != ""{
			parts = append(parts, v)
		}
	}
	return parts
}
