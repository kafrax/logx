// +build linux
// +build darwin

package logx

import (
    "fmt"
    "os"
    "path/filepath"
    "syscall"
)

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

func createLinkFile(filename, linkName string) error {
    os.Remove(linkName)
    return os.Symlink(filepath.Base(filename), linkName)
}
