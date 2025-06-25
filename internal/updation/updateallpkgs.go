package updation

import (
	"depoty/internal/badgers"
	"depoty/internal/util/common"
	"runtime"
)

func UpdateAllPkgs() bool {

	switch runtime.GOOS {
	case "windows":
		return updateAllWithChoco()
	case "linux":
		return updateAllWithApt()
	default:
		return false
	}

	// var success bool
	// i := 0
	// for i < len(pkgs) {
	// 	success = UpdatePkg(pkgs[i])
	// 	i++
	// }
	// return success
}

func updateAllWithChoco() bool {

	db := badgers.MainDb("/tmp/choco/outdate")

	defer db.Close()

	db.DropAll()

	success, err := common.ExecutePrevScript("choco upgrade all", "-y")
	if err != nil {
		return false
	}
	return success
}

func updateAllWithApt() bool {

	db := badgers.MainDb("/tmp/apt/outdate")

	defer db.Close()

	db.DropAll()

	success, err := common.ExecutePrevScript("apt upgrade", "-y")
	if err != nil {
		return false
	}
	return success
}
