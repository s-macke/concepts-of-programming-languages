package mockfs

import (
	"io/fs"
	"time"
)

type mockFileInfo struct {
	filename string
	size     int64
}

func (info mockFileInfo) Name() string {
	//fmt.Println("FileInfo: Name")
	return info.filename
}

func (info mockFileInfo) Size() int64 {
	//fmt.Println("FileInfo: Size")
	return info.size
}

func (info mockFileInfo) Mode() fs.FileMode {
	//fmt.Println("FileInfo: Mode")
	return fs.ModePerm
}

func (info mockFileInfo) IsDir() bool {
	//fmt.Println("FileInfo: IsDir")
	return true
}

func (info mockFileInfo) Sys() interface{} {
	//fmt.Println("FileInfo: Sys")
	return nil
}

func (info mockFileInfo) ModTime() time.Time {
	return time.Now()
}

func (info mockFileInfo) Info() (fs.FileInfo, error) {
	return info, nil
}

func (info mockFileInfo) Type() fs.FileMode {
	return fs.ModePerm
}
