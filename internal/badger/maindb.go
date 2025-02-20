package badger

import (
	"fmt"
	"log"

	badger "github.com/dgraph-io/badger/v4"
)

func MainDb() *badger.DB {
	opts := badger.DefaultOptions("/tmp/badger")
	// opts.Logger = nil
	db, err := badger.Open(opts)

	if err != nil {
		log.Fatal(err)
	}

	return db

}

func Insert(db *badger.DB, key []byte, value []byte) {
	err := db.Update(func(txn *badger.Txn) error {
		err := txn.Set(key, value)
		return err
	})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Value inserted")

}

func Delete(db *badger.DB, key []byte) {
	err := db.Update(func(txn *badger.Txn) error {
		return txn.Delete(key)
	})

	if err != nil {
		log.Fatal(err)
	}
}

func Read(db *badger.DB, key []byte) []byte {
	var value []byte
	err := db.View(func(txn *badger.Txn) error {
		item, err := txn.Get(key)

		if err != nil {
			log.Fatal(err)
		}

		val, err := item.ValueCopy(nil)

		value = val

		return err

	})

	if err != nil {
		log.Fatal(err)
	}

	return value

}
