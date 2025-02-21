package finding

import "depoty/internal/util/common"

func FindPkgInfo(pkgName string) string {
	info := common.ExecuteScript("choco info", pkgName)

	return info

}
