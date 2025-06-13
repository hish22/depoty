package updation

import (
	"depoty/internal/badgers"
	"depoty/internal/util/common"
)

func UpdatePkg(pkg []string) bool {

	// Open Badger
	db := badgers.MainDb("/tmp/choco/outdate")

	// Get the package name
	pkgName := pkg[0]

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
