package initalization

import (
	"depoty/internal/badgers"
	"fmt"

	badgerdb "github.com/dgraph-io/badger/v4"
)

func EntryPoint() {

	db := badgers.MainDb()

	defer db.Close()

	item, err := badgers.Read(db, []byte("initDone"))

	if err != nil {
		fmt.Println(err)
		if err == badgerdb.ErrKeyNotFound {
			fmt.Println("Starting Initalization Process")

			badgers.Insert(db, []byte("initDone"), []byte("done"))
			// Install Related dependencies
			InstallChoco()
		}
	} else {
		if string(item) == "done" {
			fmt.Println("Initalization Process is already done")
		}

	}

}
