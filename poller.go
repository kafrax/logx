package logx

import (
	"sync/atomic"
	"time"
)

func poller() {
	atomic.SwapUint32(&logger.look, uint32(coreRunning))

	if err := logger.loadCurLogFile(); err != nil {
		if err = logger.createFile(); err != nil {
			panic(err)
		}
	}

	go logger.signalHandler()

	ticker := time.NewTicker(time.Millisecond * time.Duration(pollerinterval))
	now := time.Now()
	next := now.Add(time.Hour * 24)
	next = time.Date(
		next.Year(),
		next.Month(),
		next.Day(),
		0, 0, 0, 0,
		next.Location())

	tickerPoll := time.NewTicker(next.Sub(now))

	for {
		select {
		case <-logger.closeSignal:
			ticker.Stop()
			tickerPoll.Stop()

		case <-ticker.C:
			if logger.fileWriter.Buffered() > 0 {
				logger.sync()
			}

		case n := <-logger.bucket:
			logger.fileWriter.Write(n.Bytes())
			logger.fileActualSize += n.Len()

			logger.lock.Lock()
			if logger.rotate(func() { logger.fileWriter.Flush() }) {
				logger.fileWriter.Reset(logger.file)
			}
			logger.lock.Unlock()

			logger.release(n)
		case <-tickerPoll.C:
			logger.lock.Lock()
			if logger.rotate(func() { logger.fileWriter.Flush() }) {
				logger.fileWriter.Reset(logger.file)
			}
			tickerPoll = time.NewTicker(time.Hour * 24)
		}
	}
}
