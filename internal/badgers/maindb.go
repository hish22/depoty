package badgers

import (
	"fmt"
	"log"

	badger "github.com/dgraph-io/badger/v4"
)

func MainDb(location string) *badger.DB {
	// We open defalut badger with Caching.
	opts := badger.DefaultOptions(location)
	// Empty the default badger logger for production.
	opts.Logger = nil
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

func Read(db *badger.DB, key []byte) ([]byte, error) {
	var value []byte
	err := db.View(func(txn *badger.Txn) error {
		item, err := txn.Get(key)

		if err != nil {
			return err
		}

		val, err := item.ValueCopy(nil)

		if err != nil {
			log.Fatal("Unexpected behavior happen while copying the value")
		}

		value = val

		return err

	})

	return value, err

}
