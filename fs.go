package logx

import (
    "os"
    "strings"
    "path/filepath"
)

func openFile(name string) (file *os.File, err error) {
    file, err = os.OpenFile(name, os.O_WRONLY|os.O_CREATE|os.O_APPEND|os.O_SYNC, os.ModePerm)
    if err != nil {
        return
    }
    //syscall.Syscall(syscall.O_SYNC, file.Fd(), 0, 0)

    return
}

func closeFile(file *os.File) {
    if file != nil {
        file.Close()
    }
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
