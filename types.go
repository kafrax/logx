package logx

import (
	"os"
	"bufio"
	"sync"
	"io"
)

type Logger struct {
	look           uint8
	link           string
	fileName       string
	file           *os.File
	fileBuf        *bufio.Writer
	timestamp      int
	fileMaxSize    int
	fileActualSize int
	bucket         chan *bytesBuffer
	bucketLen      int
	lock           sync.RWMutex
	out            io.Writer //out is file os.Stdout or kafaka
	//queue          chan *bytesBuffer
}
