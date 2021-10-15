package main

import (
	"context"
	"fmt"
	"github.com/jacobsa/fuse"
	"github.com/jacobsa/fuse/fuseops"
	"github.com/jacobsa/fuse/fuseutil"
	"io"
	_ "io"
	"log"
	"os"
	"strings"
	_ "strings"
)

type helloFS struct {
	fuseutil.NotImplementedFileSystem
}

func NewHelloFS() (fuse.Server, error) {
	fs := &helloFS{}
	return fuseutil.NewFileSystemServer(fs), nil
}

func (fs *helloFS) LookUpInode(
	ctx context.Context,
	op *fuseops.LookUpInodeOp) error {
	fmt.Println("Lookup INode with name ", op.Name, " and parent inode ", op.Parent)

	op.Entry.Child = 2
	op.Entry.Attributes = fuseops.InodeAttributes{
		Size:  13,
		Nlink: 1,
		Mode:  0555,
	}

	return nil
}

func (fs *helloFS) GetInodeAttributes(
	ctx context.Context,
	op *fuseops.GetInodeAttributesOp) error {
	fmt.Println("Get Inode attribute of inode", op.Inode)

	var mode os.FileMode = 0
	if op.Inode == 1 { // root node is always a directory
		mode = os.ModeDir
	}

	op.Attributes = fuseops.InodeAttributes{
		Nlink: 1,
		Mode:  0555 | mode,
	}
	return nil
}

func (fs *helloFS) OpenDir(
	ctx context.Context,
	op *fuseops.OpenDirOp) error {
	fmt.Println("Allow to open dir of inode ", op.Inode)
	return nil
}

func (fs *helloFS) ReadDir(
	ctx context.Context,
	op *fuseops.ReadDirOp) error {
	fmt.Println("Read dir of inode ", op.Inode)
	fmt.Println(op.Offset)

	var entries []fuseutil.Dirent
	dirent := fuseutil.Dirent{
		Offset: 1,
		Inode:  2,
		Name:   "Hello",
		Type:   fuseutil.DT_File,
	}
	entries = append(entries, dirent)

	if op.Offset > fuseops.DirOffset(len(entries)) {
		return fuse.EIO
	}

	entries = entries[op.Offset:]

	if op.Offset > 1 {
		return fuse.EIO
	}

	for _, e := range entries {
		n := fuseutil.WriteDirent(op.Dst[op.BytesRead:], e)
		if n == 0 {
			break
		}
		op.BytesRead += n
	}
	return nil
}

func (fs *helloFS) OpenFile(
	ctx context.Context,
	op *fuseops.OpenFileOp) error {
	fmt.Println("Allow to open file ", op.Inode)
	return nil
}

func (fs *helloFS) ReadFile(
	ctx context.Context,
	op *fuseops.ReadFileOp) error {
	fmt.Println("Read file ", op.Inode)
	reader := strings.NewReader("Hello, world!")

	var err error
	op.BytesRead, err = reader.ReadAt(op.Dst, op.Offset)

	// Special case: FUSE doesn't expect us to return io.EOF.
	if err == io.EOF {
		return nil
	}

	return err
}

func main() {

	// Create an appropriate file system.
	server, err := NewHelloFS()
	if err != nil {
		log.Fatalf("makeFS: %v", err)
	}

	cfg := &fuse.MountConfig{
		ReadOnly: true,
	}

	mfs, err := fuse.Mount(os.Args[1], server, cfg)
	if err != nil {
		log.Fatalf("Mount: %v", err)
	}

	// Wait for it to be unmounted.
	if err = mfs.Join(context.Background()); err != nil {
		log.Fatalf("Join: %v", err)
	}

}
