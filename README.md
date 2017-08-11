# logx
 log tool,easy to use,performance,handy,availability

# tps
```
 //OutPut
 //tps is : 703401/s
```
# simple to use
## install

```
go get github.com/kafrax/logx
```
## start
```
package main

import (
    "github.com/kafrax/logx"
)

func main(){
    logx.Debugf("module=test |message=%s","logx is a lightweight log to use")
}

```

# config logx
- let logx.json  or config.json in your project root dir.
- if there is not config.json,where execute default for logx.
```
{
    "llevel":1,         //log level,1debug,2info,3warn,4error,5fatal
    "lmaxsize":102400  //bit
    "lout":"stdout",   //file|stdout
    "lbucketlen":1024, //memory cache size
    "lfilename":"logx",//log file name eg. logx2006-01-02.04.05.000.log
    "lfilepath":"./",  //log file path
    "lpollerinterval": //500 millisecond
}
```

# future
 - data queue send to kafaka

# @me
 - kafrax.go@gmail.com