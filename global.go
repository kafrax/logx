package logx

import (
	"os"
)

type coreStatus uint32

var (
	coreDead    coreStatus = 2 //logger is dead
	coreBlock   coreStatus = 0  //logger is block
	coreRunning coreStatus = 1  //logger is running
)

var out = "stdout"

var maxSize int = 256 * 1024 * 1024 //256mb

var bucketLen int = 1024

var fileName string = "logx"

var filePath string = getCurrentDirectory()

type level uint8

const (
	_DEBUG    level = iota
	_INFO
	_WARN
	_ERR
	_DISASTER
)

var levelFlag level = _DEBUG