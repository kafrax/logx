package logx

import (
    "os"
    "syscall"
    "unsafe"
)

func isLinkFile(filename string) (name string, ok bool) {
    fi, err := os.Lstat(filename)
    if err != nil {
        return
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

//https://gist.github.com/bradleypeabody/10572010
var (
    k32                = syscall.MustLoadDLL("kernel32.dll")
    createSymbolicLink = k32.MustFindProc("CreateSymbolicLinkW")
)

func createLinkFile(oldpath, newpath string) error {
    os.Remove(oldpath)
    st, err := os.Stat(oldpath)
    if err != nil {
        return err
    }


    linkType := 0
    if st.Mode() == os.ModeDir {
        linkType = 1
    }

    _, _, callErr := createSymbolicLink.Call(
        uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(newpath))),
        uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(oldpath))),
        uintptr(linkType),
    )

    errno, _ := callErr.(syscall.Errno)
    if errno != 0 {
        return callErr
    }

    return nil

}
