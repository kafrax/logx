package logx

import (
	"os"
	"syscall"
	"fmt"
	"path/filepath"
	"strings"
)

func openFile(name string) (file *os.File, err error) {
	file, err = os.OpenFile(name, os.O_WRONLY|os.O_CREATE|os.O_APPEND, os.ModePerm)
	if err != nil {
		return
	}
	syscall.Syscall(syscall.O_SYNC, file.Fd(), 0, 0)

	return
}

func closeFile(file *os.File) {
	if file != nil {
		file.Close()
	}
}

func isLinkFile(filename string) (name string, ok bool) {
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

func createLinkFile(filename, linkName string) error {
	os.Remove(linkName)
	return os.Symlink(filepath.Base(filename), linkName)
}

func substr(s string, pos, length int) string {
	runes := []rune(s)
	l := pos + length
	if l > len(runes) {
		l = len(runes)
	}
	return string(runes[pos:l])
}

func getParentDirectory(dirctory string) string {
	return substr(dirctory, 0, strings.LastIndex(dirctory, "/"))
}

func getCurrentDirectory() string {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	return filepath.Dir(ex)
}
