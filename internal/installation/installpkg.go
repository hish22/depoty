package installation

import (
	"depoty/internal/util/common"
	"runtime"
)

func InstallPkg(pkgName string) bool {
	switch runtime.GOOS {
	case "windows":
		return installwithchoco(pkgName)
	case "linux":
		return installwithapt(pkgName)
	default:
		return false
	}
}

func installwithchoco(pkgName string) bool {
	success, err := common.ExecutePrevScript("choco install", pkgName+" -y")
	if err != nil {
		return false
	}
	return success
}

func installwithapt(pkgName string) bool {
	success, err := common.ExecutePrevScript("apt install", pkgName+" -y")
	if err != nil {
		return false
	}
	return success
}
