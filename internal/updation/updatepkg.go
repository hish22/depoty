package updation

import (
	"depoty/internal/badgers"
	"depoty/internal/util/common"
)

func UpdatePkg(pkgName string) bool {

	// Open Badger
	db := badgers.MainDb("/tmp/badger/outdate")

	// Close badger
	defer db.Close()

	//Delete package from outdated packages
	if _, err := badgers.Read(db, []byte(pkgName)); err == nil {
		badgers.Delete(db, []byte(pkgName))
	}

	return common.ExecutePrevScript("choco upgrade", pkgName+" -y")
}
