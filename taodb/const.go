package taodb


import (
	"sync"
)
const (
	ERROR_STRORAGE_NOT_ACTIVE = "storage is not active"
	ERROR_META_NOT_SET = "meta is not set"
	KB = 1024
	MB = 1024 * KB
	GB = 1024 * MB
	
)


var (
	TransactionId = uint32(0)
	TransactionMutex = &sync.Mutex{}
	
)
