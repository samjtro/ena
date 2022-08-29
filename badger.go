package main

import (
	"errors"

	badger "github.com/dgraph-io/badger/v3"
)

func CheckSimilarity(db *badger.DB, hash string) error {
	err := db.View(func(txn *badger.Txn) error {
		it := txn.NewIterator(badger.DefaultIteratorOptions)

		defer it.Close()

		for it.Rewind(); it.Valid(); it.Next() {
			item := it.Item()

			var key, val []byte
			item.KeyCopy(key)
			item.ValueCopy(val)

			if string(key) == keywordFlag {
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
