package installation

import "depoty/internal/util/common"

func InstallPkg(pkgName string) bool {

	success, err := common.ExecutePrevScript("choco install", pkgName+" -y")

	if err != nil {
		return false
	}

	return success
}
