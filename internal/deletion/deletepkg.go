package deletion

import (
	"depoty/internal/badgers"
	"depoty/internal/util/common"
)

func DeletePkg(pkgName string) bool {

	// Open Badger
	db := badgers.MainDb("/tmp/badger/outdate")

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
