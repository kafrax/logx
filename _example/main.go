package main

import (
	"github.com/kafrax/logx"
	"sync/atomic"
	"time"
	"fmt"
)

func main() {
	var tps int64 = 0

	for i := 0; i < 20; i++ {
		go func() {
			for j := 0; j < 10000000; j++ {
				logx.Disasterf("tps test |how times logx can bear |message=%v", "ahha ahhaa")
				atomic.AddInt64(&tps, 1)
			}
		}()
	}

	for i := 0; i < 20; i++ {
		time.Sleep(time.Second)
		fmt.Println("tps is :", atomic.LoadInt64(&tps))
		atomic.SwapInt64(&tps, 0)
	}
}
