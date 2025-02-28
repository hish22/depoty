package installation

import (
	"depoty/internal/util/common"
)

func InstallPkg(pkg []string) bool {

	success, err := common.ExecutePrevScript("choco install", pkg[0]+" --version="+pkg[1]+" -y")

	if err != nil {
		return false
	}

	return success
}
