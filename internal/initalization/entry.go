package initalization

import (
	"depoty/internal/badgers"
	"fmt"

	badgerdb "github.com/dgraph-io/badger/v4"
)

func EntryPoint() {

	db := badgers.MainDb("/tmp/badger/config")

	defer db.Close()

	item, err := badgers.Read(db, []byte("initDone"))

	// Install Related dependencies
	InstallChoco()

	if err != nil {
		fmt.Println("The configuration process is not initialized yet.")
		if err == badgerdb.ErrKeyNotFound {
			fmt.Println("Starting Initalization Process..")

			badgers.Insert(db, []byte("initDone"), []byte("done"))

		}
	} else {
		if string(item) == "done" {
			fmt.Println("Initalization Process is already done.")
		}

	}

}
