package get

import "github.com/deepakr199/go/basic/dbstorage/storage"

func Get(key string) string {
	m := storage.Storage()
	value := m[key]
	return value
}