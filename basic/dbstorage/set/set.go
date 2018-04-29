package set

import "github.com/deepakr199/go/basic/dbstorage/storage"

func Set(key string, value string) string {
	m := storage.Storage()
	m[key] = value
	return "ok"
}
