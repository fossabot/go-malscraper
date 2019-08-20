package helper

import (
	"encoding/json"
	_ "fmt"
	_ "reflect"
	"time"

	bolt "go.etcd.io/bbolt"
)

const MalBucket string = "MalBucket"

func InitBolt() *bolt.DB {
	db, err := bolt.Open("mal-cache.db", 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		panic(err)
	}

	err = db.Update(func(tx *bolt.Tx) error {
		_, e := tx.CreateBucketIfNotExists([]byte(MalBucket))
		return e
	})

	if err != nil {
		panic(err)
	}

	return db
}

func SaveToCache(key string, value interface{}) {
	db := InitBolt()
	defer db.Close()

	jsonValue, err := json.Marshal(value)
	if err != nil {
		panic(err)
	}

	err = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(MalBucket))
		return b.Put([]byte(key), []byte(string(jsonValue)))
	})

	if err != nil {
		panic(err)
	}
}

func GetFromCache(key string) ([]byte, error) {
	db := InitBolt()
	defer db.Close()

	value := []byte("")

	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(MalBucket))
		value = b.Get([]byte(key))
		return nil
	})

	return value, err
}
