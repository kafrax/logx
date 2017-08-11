package main

import (
	"logx"
)

//func main() {
//	var now=time.Now()
//
//	for j := 0; j < 1000000; j++ {
//		logx.Debugf("tps test |how times logx can bear |message=%v", "ahha ahhaa")
//	}
//	fmt.Println("time",time.Now().Sub(now).Seconds())
//	var str string
//	fmt.Scanln(&str)
//
//	//OutPut
//	//time 3.166870248
//}

//func main() {
//	logx.Debugf("tps test |how times logx can bear |message=%v", "ahha ahhaa")
//	var tps int64 = 0
//
//	for i := 0; i < 20; i++ {
//		go func() {
//			for j := 0; j < 10000000; j++ {
//				logx.Debugf("tps test |how times logx can bear |message=%v", "ahha ahhaa")
//				atomic.AddInt64(&tps, 1)
//			}
//		}()
//	}
//
//	for i := 0; i < 20; i++ {
//		time.Sleep(time.Second)
//		fmt.Println("tps is :", atomic.LoadInt64(&tps))
//		atomic.SwapInt64(&tps, 0)
//	}
//	//OutPut
//	//tps is : 703401
//}

func main() {
	logx.Stackf("test |message=%s", "ahhh")
}
