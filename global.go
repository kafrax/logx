package logx

import (
	"os"
)

type coreStatus uint32

var (
	coreDead    coreStatus = -1 //logger is dead
	coreBlock   coreStatus = 0  //logger is block
	coreRunning coreStatus = 1  //logger is running
)

var out = os.Stdout

var maxSize int = 256 * 1024 * 1024 //256mb

var bucketLen int = 1024

var fileName string = "kafrax-logx"

var filePath string = getCurrentDirectory()

type level uint

const (
	debug    level = iota
	info
	warn
	err
	disaster
)
