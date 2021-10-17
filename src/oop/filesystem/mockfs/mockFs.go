package mockfs

import (
	"io/fs"
)

type MockFs struct {
	content []byte
	info    mockFileInfo
}

func (mockFs MockFs) Open(name string) (fs.File, error) {
	//fmt.Println("File: Open", name)
	file := mockFile{
		info:    mockFs.info,
		content: mockFs.content,
		offset:  int64(0),
	}
	return &file, nil
}

func (mockFs MockFs) ReadDir(name string) ([]fs.DirEntry, error) {
	//fmt.Println("FS: Readdir in ", name)
	return []fs.DirEntry{&mockFs.info}, nil
}

func NewMockFs(filename string, content string) MockFs {
	return MockFs{
		info: mockFileInfo{
			filename: filename,
			size:     int64(len(content)),
		},
		content: []byte(content),
	}
}
