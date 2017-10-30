package logx

import (
	"bufio"
	"bytes"
	"io"
	"os"
	"sync"
)

type Hook interface {
	Fire(writer *bufio.Writer)
	Level(level)
}

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

type config struct {
	Llevel          uint8 `json:"llevel"`
	Lmaxsize        int    `json:"lmaxsize"`
	Lout            string `json:"lout"`
	Lbucketlen      int    `json:"lbucketlen"`
	Lfilename       string `json:"lfilename"`
	Lfilepath       string `json:"lfilepath"`
	Lpollerinterval int    `json:"lpollerinterval"`
}
