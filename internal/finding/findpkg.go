package finding

import (
	"depoty/internal/util/common"
	"runtime"
	"strings"
)

func FindPkg(pkgName string) []string {
	switch runtime.GOOS {
	case "windows":
		return findwithchoco(pkgName)
	case "linux":
		return findwithapt(pkgName)
	default:
		return make([]string, 0)
	}
}

func findwithchoco(pkgName string) []string {
	text := common.ExecuteScript("choco find -r", pkgName)

	textSlice := strings.Split(text, "\n")

	return textSlice
}

func findwithapt(pkgName string) []string {
	text := common.ExecuteScript("apt search", pkgName)

	textSlice := strings.Split(text, "\n")

	var searchlist []string
	var i int
	for i = 2; i < len(textSlice)-1; i += 3 {
		name := strings.Split(textSlice[i], "/")
		searchlist = append(searchlist, name[0])
	}
	return searchlist
}
