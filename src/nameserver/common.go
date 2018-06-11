package nameserver

import (
	"strings"
	"os"
)

func SplistPath(fname string) []string{
	temp := strings.Split(string(f), string(os.PathSeparator))
	parts := []string{}
	for _, v := range temp{
		if v != ""{
			parts = append(parts, v)
		}
	}
	return parts
}
