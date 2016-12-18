package main

import (
	"encoding/binary"
	"fmt"
	"log"

	"github.com/boltdb/bolt"
)

func main() {
	db, err := bolt.Open("/Users/panqd/panqd.txt.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Batch(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("MyBucket"))
		id, _ := b.NextSequence()
		key := make([]byte, 8)
		binary.BigEndian.PutUint64(key, id)
		fmt.Println(key)
		return nil
	})

	db.Update(func(tx *bolt.Tx) error {
		tx.CreateBucketIfNotExists([]byte("MyBucket"))
		return nil
	})

	db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("MyBucket"))
		err := b.Put([]byte("answer"), []byte("42"))
		return err
	})

	db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("MyBucket"))
		b.Delete(nil)
		return nil
	})

	mmap := make(map[string]string)
	mmap["panqdKey"] = "panqdValue"
	mmap["kingdomKey"] = "kingdomValue"

	for k, v := range mmap {
		fmt.Println(k, v)
		mmap[v] = k
	}
}
