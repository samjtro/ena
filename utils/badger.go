package utils

import (
	"errors"
	"log"

	badger "github.com/dgraph-io/badger/v3"
)

func Start() *badger.DB {
	db, err := badger.Open(badger.DefaultOptions("/tmp/badger"))

	if err != nil {
		log.Fatal(err)
	}

	return db
}

func AddKeyValue(db *badger.DB, k, v string) {
	err := db.Update(func(txn *badger.Txn) error {
		e := badger.NewEntry([]byte(k), []byte(v))
		err := txn.SetEntry(e)
		return err
	})

	if err != nil {
		log.Fatal(err)
	}
}

func CheckSimilarity(db *badger.DB, keyword, hash string) error {
	err := db.View(func(txn *badger.Txn) error {
		it := txn.NewIterator(badger.DefaultIteratorOptions)

		defer it.Close()

		for it.Rewind(); it.Valid(); it.Next() {
			item := it.Item()

			var key, val []byte
			item.KeyCopy(key)
			item.ValueCopy(val)

			if string(key) == keyword {
				if string(val) == hash {
					return nil
				} else {
					return errors.New("")
				}
			} else {
				return errors.New("")
			}
		}

		return nil
	})

	return err
}
