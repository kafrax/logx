package logx

import (
	"os"
	"bufio"
	"sync"
	"io"
	"bytes"
)

type Logger struct {
	look           uint32 //monitor run state with block stop or running
	link           string
	path           string
	fileName       string
	file           *os.File
	fileWriter     *bufio.Writer
	timestamp      int
	fileMaxSize    int
	fileActualSize int
	bucket         chan *bytes.Buffer
	bucketFlushLen int
	lock           *sync.RWMutex
	output         io.Writer //out is file os.Stdout or kafaka
	closeSignal    chan string
	//queue          chan *bytesBuffer
}
