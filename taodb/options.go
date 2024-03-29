package taodb

import (
	"os"
)
type Options struct {
	// 数据库路径
	DirPath string
	MaxStorageSize uint32
}


var DefaultOptions = Options{
	DirPath: tempDBDir(),
	MaxStorageSize: 1 * GB,
}

func tempDBDir() string {
	dir, _ := os.MkdirTemp("", "taodb-temp")
	return dir
}