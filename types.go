package logx

import (
	"os"
	"bufio"
	"sync"
	"io"
)

type Logger struct {
	look           coreStatus //monitor run state with block stop or running
	link           string
	path           string
	fileName       string
	file           *os.File
	fileWriter     *bufio.Writer
	timestamp      int64
	fileMaxSize    int
	fileActualSize int
	bucket         chan *bytesBuffer
	bucketFlushLen int
	lock           *sync.RWMutex
	output         io.Writer //out is file os.Stdout or kafaka
	//queue          chan *bytesBuffer
}

type logMsg struct {

	timestamp string
	msg       string
	level     uint8
}
