package deletion

import (
	"depoty/internal/badgers"
	"depoty/internal/util/common"
	"runtime"
)

func DeletePkg(pkgName string) bool {
	switch runtime.GOOS {
	case "windows":
		return deletePkgsWithChoco(pkgName)
	case "linux":
		return deletePkgsWithApt(pkgName)
	default:
		return false
	}
}

func deletePkgsWithChoco(pkgName string) bool {
	// Open Badger
	db := badgers.MainDb("/tmp/choco/outdate")

	// Close Badger
	defer db.Close()

	// Delete the cache of outdated package if it is deleted
	if _, err := badgers.Read(db, []byte(pkgName)); err == nil {
		badgers.Delete(db, []byte(pkgName))
	}

	success, err := common.ExecutePrevScript("choco uninstall", pkgName+" -y")

	if err != nil {
		return false
	}

	return success
}

func deletePkgsWithApt(pkgName string) bool {
	// Open Badger
	db := badgers.MainDb("/tmp/apt/outdate")

	// Close Badger
	defer db.Close()

	// Delete the cache of outdated package if it is deleted
	if _, err := badgers.Read(db, []byte(pkgName)); err == nil {
		badgers.Delete(db, []byte(pkgName))
	}

	success, err := common.ExecutePrevScript("apt purge", pkgName+" -y")

	if err != nil {
		return false
	}

	return success
}
