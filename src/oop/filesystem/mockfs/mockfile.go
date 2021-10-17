package mockfs

import (
	"io"
	"io/fs"
)

type mockFile struct {
	info    mockFileInfo
	content []byte
	offset  int64
}

func (file *mockFile) Stat() (fs.FileInfo, error) {
	//fmt.Println("File: Stat")
	return file.info, nil
}

func (file *mockFile) Read(bytes []byte) (int, error) {
	var err error = nil
	//fmt.Println("File: Read of size", len(bytes))

	// Determine Read size dependent on the current offset of the file pointer
	size := int64(len(bytes))
	if file.info.Size() <= file.offset+size {
		size = file.info.Size() - file.offset
		err = io.EOF // end of file
	}

	copy(bytes, file.content[file.offset:file.offset+size])

	file.offset += size
	return int(size), err
}

func (file *mockFile) Close() error {
	//fmt.Println("File: Close")
	return nil
}
