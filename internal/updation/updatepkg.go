package updation

import (
	"depoty/internal/badgers"
	"depoty/internal/util/common"
)

func UpdatePkg(pkgName string) bool {

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
