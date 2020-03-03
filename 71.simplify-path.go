package main

import (
	"fmt"
	"strings"
)

func main() {
	path1 := "/home/"
	fmt.Println(simplifyPath(path1))
	path2 := "/../"
	fmt.Println(simplifyPath(path2))
	path3 := "/home//foo/"
	fmt.Println(simplifyPath(path3))
	path4 := "/a/../../b/../c//.//"
	fmt.Println(simplifyPath(path4))
	path5 := "/a//b////c/d//././/.."
	fmt.Println(simplifyPath(path5))
}

func simplifyPath(path string) string {
	list := []string{}
	arr := strings.Split(path, "/")
	for _, p := range arr {
		if p == "" || p == "." {

		} else if p == ".." {
			if len(list) > 0 {
				list = list[:len(list)-1]
			}
		} else {
			list = append(list, p)
		}
	}
	if len(list) == 0 {
		return "/"
	}

	res := ""
	for _, p := range list {
		res += "/" + p
	}
	return res
}
