package finding

import (
	"depoty/internal/util/common"
	"strings"
)

func FindPkg(pkgName string) []string {
	text := common.ExecuteScript("choco find", pkgName)

	textSlice := strings.Split(text, "\n")

	return textSlice

}
