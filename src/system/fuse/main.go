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
	"math/rand"
	"os"
	"os/user"
	"strconv"
	"strings"
	_ "strings"
	"time"
)

type helloFS struct {
	fuseutil.NotImplementedFileSystem
}

func NewHelloFS() fuse.Server {
	fs := &helloFS{}
	return fuseutil.NewFileSystemServer(fs)
}

func (fs *helloFS) GetInodeAttributes(ctx context.Context, op *fuseops.GetInodeAttributesOp) error {
	fmt.Println("Get Inode attribute of inode", op.Inode)

	var mode os.FileMode = 0 // is a normal file
	if op.Inode == 1 {       // root node is always a directory
		mode = os.ModeDir
	}

	user, _ := user.Current()
	uid, _ := strconv.Atoi(user.Uid)
	gid, _ := strconv.Atoi(user.Gid)
	op.Attributes = fuseops.InodeAttributes{
		Size:  0,
		Nlink: 1,
		Mode:  0555 | mode,
		Atime: time.Now(),
		Mtime: time.Now(),
		Ctime: time.Now(),
		Uid:   uint32(uid),
		Gid:   uint32(gid),
	}
	return nil
}

func (fs *helloFS) OpenDir(ctx context.Context, op *fuseops.OpenDirOp) error {
	fmt.Println("Allow to open dir of inode ", op.Inode)
	return nil
}

func (fs *helloFS) ReadDir(ctx context.Context, op *fuseops.ReadDirOp) error {
	fmt.Println("Read dir of inode ", op.Inode, " at offset ", op.Offset)

	// We have only one directory
	var entries []fuseutil.Dirent
	for i := 0; i < 10; i++ {
		dirent := fuseutil.Dirent{
			Offset: fuseops.DirOffset(i + 1),
			Inode:  fuseops.InodeID(2), // This is wrong. An InodeID should be unique
			Name:   fmt.Sprintf("Hello%d", rand.Int()),
			Type:   fuseutil.DT_File,
		}
		entries = append(entries, dirent)
	}

	if op.Offset > fuseops.DirOffset(len(entries)) {
		return fuse.EIO
	}

	entries = entries[op.Offset:]

	for _, e := range entries {
		n := fuseutil.WriteDirent(op.Dst[op.BytesRead:], e)
		if n == 0 {
			break
		}
		op.BytesRead += n
	}
	return nil
}

func (fs *helloFS) LookUpInode(ctx context.Context, op *fuseops.LookUpInodeOp) error {
	fmt.Println("Lookup INode with name ", op.Name, " and parent inode ", op.Parent)

	op.Entry.Child = 2 // This is wrong. An InodeID should be unique
	op.Entry.Attributes = fuseops.InodeAttributes{
		Size:  13,
		Nlink: 1,
		Mode:  0555,
	}

	return nil
}

func (fs *helloFS) OpenFile(ctx context.Context, op *fuseops.OpenFileOp) error {
	fmt.Println("Allow to open file ", op.Inode)
	return nil
}

func (fs *helloFS) ReadFile(ctx context.Context, op *fuseops.ReadFileOp) error {
	fmt.Println("Read file ", op.Inode, "from", op.Offset, "for", len(op.Dst), "bytes")
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
	server := NewHelloFS()

	cfg := &fuse.MountConfig{
		ReadOnly: true,
		Options:  map[string]string{"allow_other": ""},
	}

	mfs, err := fuse.Mount(os.Args[1], server, cfg)
	if err != nil {
		log.Fatalf("Mount: %v", err)
	}

	// Wait for it to be unmounted.
	err = mfs.Join(context.Background())
	if err != nil {
		log.Fatalf("Join: %v", err)
	}
}
