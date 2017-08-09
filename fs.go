package logx

import (
	"os"
	"syscall"
	"fmt"
	"path/filepath"
)

func openFile(name string, do func()) error {
	file, err := os.OpenFile(name, os.O_WRONLY|os.O_CREATE|os.O_APPEND, os.ModePerm)
	if err != nil {
		return err
	}
	syscall.Syscall(syscall.O_SYNC, file.Fd(), 0, 0)

	do()

	return err
}

func closeFile(file *os.File) {
	if file != nil {
		file.Close()
	}
}

func IsLinkFile(filename string) (name string, ok bool) {
	fi, err := os.Lstat(filename)
	if err != nil {
		return
	}
	s, ok := fi.Sys().(*syscall.Stat_t)
	if !ok {
		return
	}
	ino := uint64(s.Ino)
	nlink := uint32(s.Nlink)
	if nlink > 1 {
		fmt.Printf("link info Inode %v has %v other hardlinks with %v.\n", ino, nlink, filename)
	}

	if fi.Mode()&os.ModeSymlink != 0 {
		name, err = os.Readlink(filename)
		if err != nil {
			return
		}
		return name, true
	} else {
		return
	}

	return
}

func pathIsExist(path string) bool {
	if _, err := os.Stat(path); err == nil {
		return true
	} else {
		if os.IsNotExist(err) {
			return false
		}
	}
	return false
}

func createLinkFile(path, filename, linkname string) {
	os.Remove(filepath.Join(path, linkname))
	os.Symlink(filepath.Join(path, filename), filepath.Join(path, linkname))

}
