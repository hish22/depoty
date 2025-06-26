package finding

import (
	"depoty/internal/util/common"
	"runtime"
)

func FindPkgInfo(pkgName string) string {
	switch runtime.GOOS {
	case "windows":
		return findPkgChocoInfo(pkgName)
	case "linux":
		return findPkgAptInfo(pkgName)
	default:
		return ""
	}
}

func findPkgChocoInfo(pkgName string) string {
	info := common.ExecuteScript("choco info", pkgName)

	return info
}

func findPkgAptInfo(pkgName string) string {
	info := common.ExecuteScript("apt show", pkgName)

	return info
}
