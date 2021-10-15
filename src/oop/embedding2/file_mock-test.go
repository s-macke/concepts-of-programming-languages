package main

import (
	"fmt"
	"io/fs"
	"time"
)

type myFs string
type myFile string
type myInfo string

// A FileInfo describes a file and is returned by Stat.
type FileInfo interface {
	Name() string       // base name of the file
	Size() int64        // length in bytes for regular files; system-dependent for others
	Mode() fs.FileMode  // file mode bits
	ModTime() time.Time // modification time
	IsDir() bool        // abbreviation for Mode().IsDir()
	Sys() interface{}   // underlying data source (can return nil)
}

func (file myInfo) Name() string {
	return "myFile"
}

func (file myInfo) Size() int64 {
	return 0
}

func (file myInfo) Mode() fs.FileMode {
	return 0
}

func (file myInfo) IsDir() bool {
	return true
}

func (file myInfo) Sys() interface{} {
	return nil
}

func (file myInfo) ModTime() time.Time {
	return time.Now()
}

func (file myFile) Stat() (fs.FileInfo, error) {
	fmt.Println("Stat")
	var info myInfo
	return info, nil
}

func (file myFile) Read([]byte) (int, error) {
	fmt.Println("Read")
	return 0, nil
}

func (file myFile) Close() error {
	fmt.Println("Close")
	return nil
}

func (fs myFs) Open(name string) (fs.File, error) {
	fmt.Println("Open", name)
	var file myFile
	return file, nil
}

func main() {
	//dat, err := os.ReadFile("/tmp/dat")
	var fsroot myFs
	/*
		fsroot := os.DirFS("")
	*/
	//fmt.Println(reflect.TypeOf(fsroot).Kind())

	err := fs.WalkDir(fsroot, "usr/include", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			fmt.Println("Walk", err)
		}
		fmt.Println("Walk", path)
		return nil
	})
	fmt.Println(err)
	/*
		f, _ := fsroot.Open("test")
		defer f.Close()
	*/
}
