package updation

import (
	"depoty/internal/badgers"
	"depoty/internal/util/common"
	"runtime"
)

func UpdatePkg(pkgName string) bool {
	switch runtime.GOOS {
	case "windows":
		return updatePkgWithChoco(pkgName)
	case "linux":
		return updatePkgWithApt(pkgName)
	default:
		return false
	}
}

func updatePkgWithChoco(pkgName string) bool {
	// Open Badger
	db := badgers.MainDb("/tmp/choco/outdate")

	// Close badger
	defer db.Close()

	//Delete package from outdated packages
	if _, err := badgers.Read(db, []byte(pkgName)); err == nil {
		badgers.Delete(db, []byte(pkgName))
	}

	success, err := common.ExecutePrevScript("choco upgrade", pkgName+" -y")

	if err != nil {
		return false
	}

	return success
}

func updatePkgWithApt(pkgName string) bool {
	// Open Badger
	db := badgers.MainDb("/tmp/apt/outdate")

	// Close badger
	defer db.Close()

	//Delete package from outdated packages
	if _, err := badgers.Read(db, []byte(pkgName)); err == nil {
		badgers.Delete(db, []byte(pkgName))
	}

	success, err := common.ExecutePrevScript("apt install --only-upgrade", pkgName+" -y")

	if err != nil {
		return false
	}

	return success
}
