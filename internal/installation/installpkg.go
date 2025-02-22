package installation

import "depoty/internal/util/common"

func InstallPkg(pkgName string) bool {
	return common.ExecutePrevScript("choco install", pkgName+" -y")
}
