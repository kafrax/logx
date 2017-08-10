package logx

import (
	"time"
	"sync"
	"os"
	"errors"
	"strings"
	"fmt"
	"bufio"
	"path/filepath"
	"sync/atomic"
	"os/signal"
	"syscall"
	"runtime"
	"path"
	"strconv"
	"bytes"
)

var logger *Logger

func init() {
	newLogger()
	go poller()
}

func newLogger() *Logger {
	if logger == nil {
		return &Logger{
			look:        coreDead,
			link:        fileName + ".log",
			fileName:    fileName,
			path:        filepath.Join(filePath, fileName, "log"),
			timestamp:   time.Now().Unix(),
			fileMaxSize: maxSize,
			bucket:      make(chan *bytes.Buffer, bucketLen),
			lock:        &sync.RWMutex{},
			//output         io.Writer //out is file os.Stdout or kafaka
		}
	}
	return logger
}

func (l *Logger) loadCurLogFile() error {
	actFileName, ok := isLinkFile(l.link)
	if !ok {
		return errors.New("is not link file")
	}

	l.fileName = filepath.Join(l.path, actFileName)
	f, err := openFile(l.fileName)
	if err != nil {
		return err
	}

	info, err := os.Stat(l.fileName)
	if err != nil {
		return err
	}
	sp := strings.Split(actFileName, ".")
	t, err := time.Parse("2006-01-02", sp[1])
	if err != nil {
		fmt.Errorf("loadCurrentLogFile |err=%v", err)
		return err
	}

	l.timestamp = t.Unix()
	l.file = f
	l.fileActualSize = int(info.Size())
	l.fileWriter = bufio.NewWriterSize(f, l.fileMaxSize)

	return nil
}

func (l *Logger) createFile() (err error) {
	if !pathIsExist(l.path) {
		if err = os.MkdirAll(l.path, os.ModePerm); err != nil {
			return
		}
	}

	now := time.Now()

	l.timestamp = now.Unix()
	l.fileName = filepath.Join(
		l.path,
		filepath.Base(os.Args[0])+"."+now.Format("2006-01-02.15.04.05.000")+".log")

	f, err := openFile(l.fileName)
	if err != nil {
		return err
	}

	l.file = f
	l.fileActualSize = 0
	l.fileWriter = bufio.NewWriterSize(f, l.fileMaxSize)
	return createLinkFile(l.fileName, l.link)
}

func (l *Logger) sync() {
	if l.lookRunning() {
		l.fileWriter.Flush()
	}
}

func (l *Logger) rotate(do func()) bool {
	if l.fileActualSize <= l.fileMaxSize {
		return false
	}
	do()

	closeFile(l.file)

	if err := l.createFile(); err != nil {
		return false
	}
	return true
}

func (l *Logger) lookRunning() bool {
	if atomic.LoadUint32(&uint32(l.look)) == uint32(coreRunning) {
		return true
	}
	return false
}

func (l *Logger) lookDead() bool {
	if atomic.LoadUint32(&uint32(l.look)) == uint32(coreDead) {
		return true
	}
	return false
}

func (l *Logger) lookBlock() bool {
	if atomic.LoadUint32(&uint32(l.look)) == uint32(coreBlock) {
		return true
	}
	return false
}

func (l *Logger) signalHandler() {
	var sigChan = make(chan os.Signal)
	signal.Notify(sigChan, syscall.SIGTERM, syscall.SIGINT)

	for {
		select {
		case sig := <-sigChan:
			fmt.Println("receive os signal is ", sig)
			l.fileWriter.Flush()
			closeFile(l.file)
			atomic.SwapUint32(&uint32(l.look), uint32(coreDead))
			close(l.bucket)
			os.Exit(syscall.SYS_EXIT)
		}
	}
}

func (l *Logger) release(buf *bytes.Buffer) {
	bufferPoolFree(buf)
}

func caller() string {
	pc := make([]uintptr, 3, 3)
	cnt := runtime.Callers(6, pc)

	for i := 0; i < cnt; i++ {
		fu := runtime.FuncForPC(pc[i] - 1)
		name := fu.Name()

		if !strings.Contains(name, "github.com/kafrax/logx") {
			f, l := fu.FileLine(pc[i] - 1)
			return path.Base(f) + "|" + path.Base(name) + "|" + strconv.Itoa(l)
		}

		if pc, f, l, ok := runtime.Caller(8); ok {
			funcName := runtime.FuncForPC(pc).Name()
			return path.Base(f) + "|" + path.Base(funcName) + "|" + strconv.Itoa(l)
		}
	}
	return ""
}

func Debugf(format, msg string) {
	if levelFlag > _DEBUG {
		return
	}
	buf := bufferPoolGet()
	buf.Write(s2b("[DEBU][" + time.Now().Format("01-02.15.04.05.000") + "] " + "[" + caller() + "] "))
	buf.Write(s2b(fmt.Sprintf(format, msg)))
	logger.bucket <- buf
}

func Infof(format, msg string) {
	if levelFlag > _INFO {
		return
	}
	buf := bufferPoolGet()
	buf.Write(s2b("[INFO][" + time.Now().Format("01-02.15.04.05.000") + "] " + "[" + caller() + "] "))
	buf.Write(s2b(fmt.Sprintf(format, msg)))
	logger.bucket <- buf
}

func Warnf(format, msg string) {
	if levelFlag > _WARN {
		return
	}
	buf := bufferPoolGet()
	buf.Write(s2b("[WARN][" + time.Now().Format("01-02.15.04.05.000") + "] " + "[" + caller() + "] "))
	buf.Write(s2b(fmt.Sprintf(format, msg)))
	logger.bucket <- buf
}

func Errorf(format, msg string) {
	if levelFlag > _ERR {
		return
	}
	buf := bufferPoolGet()
	buf.Write(s2b("[ERRO][" + time.Now().Format("01-02.15.04.05.000") + "] " + "[" + caller() + "] "))
	buf.Write(s2b(fmt.Sprintf(format, msg)))
	logger.bucket <- buf
}

func Disaster(format, msg string) {
	if levelFlag > _DISASTER {
		return
	}
	buf := bufferPoolGet()
	buf.Write(s2b("[DISA][" + time.Now().Format("01-02.15.04.05.000") + "] " + "[" + caller() + "] "))
	buf.Write(s2b(fmt.Sprintf(format, msg)))
	logger.bucket <- buf
}
