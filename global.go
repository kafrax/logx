package logx

import "os"

type coreStatus uint8

var (
	coreDead coreStatus = -1 //logger is dead
	coreBlok coreStatus = 0  //logger is block
	coreRunx coreStatus = 1  //logger is running
)

var out = os.Stdout

var maxSize int = 256 * 1024 * 1024 //256mb

var bketLen int = 1024

var filenName string = "kafrax-logx"
