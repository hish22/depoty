package initalization

import (
	"depoty/internal/badger"
	"fmt"

	badgerdb "github.com/dgraph-io/badger/v4"
)

func EntryPoint() {

	db := badger.MainDb()

	defer db.Close()

	item, err := badger.Read(db, []byte("initDone"))

	if err != nil {
		fmt.Println(err)
		if err == badgerdb.ErrKeyNotFound {
			fmt.Println("Starting Initalization Process")

			badger.Insert(db, []byte("initDone"), []byte("done"))
			// Install Related dependencies
			InstallChoco()
		}
	} else {
		if string(item) == "done" {
			fmt.Println("Initalization Process is already done")
		}

	}

}
