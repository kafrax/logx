# logx
 - log tool,easy to use,high performance,handy,availability
 - version v1.1
# tps
```
 //OutPut to file
 //tps is : 1400000/s on windows
 //cpu i5-7600 3.5GHZ
 //8GB
 //it will be better on better platform
```
# simple to use
## install

```
go get -u github.com/kafrax/logx
```
## start
```
package main

import (
    "github.com/kafrax/logx"
)

func main(){
    logx.Debugf("LOGX |message=%v |substring=%s", "logx is a lightweight log to use", "debugf test")
    logx.Infof("LOGX |message=%s", "logx is a lightweight log to use")
    logx.Errorf("LOGX |message=%s", "logx is a lightweight log to use")
    logx.Warnf("LOGX |message=%s", "logx is a lightweight log to use")
    logx.Fatalf("LOGX |message=%s", "logx is a lightweight log to use")
}
```
```
[DEBU][08-18.13.34.47.703][main.go|main.main|51] LOGX |message=logx is a lightweight log to use |substring=debugf test
[INFO][08-18.13.34.47.703][main.go|main.main|52] LOGX |message=logx is a lightweight log to use
[ERRO][08-18.13.34.47.703][main.go|main.main|53] LOGX |message=logx is a lightweight log to use
[WARN][08-18.13.34.47.703][main.go|main.main|54] LOGX |message=logx is a lightweight log to use
[FTAL][08-18.13.34.47.703][main.go|main.main|55] LOGX |message=logx is a lightweight log to use
```

#  write to file
## config logx.json or config.json
- let logx.json  or config.json in your project root dir.
- will be executed by default , there is no config.json or logx.json yet.
- *notice* fileWriter use memory cache ,so must have enough time to do poller to save data to log file.
```
{
    "llevel":1,        //log level,1debug,2info,3warn,4error,5fatal
    "lmaxsize":256     //256mb
    "lout":"stdout",   //file|stdout
    "lbucketlen":1024, //log message bucket cache size
    "lfilename":"logx",//log file name eg. logx2006-01-02.04.05.000.log
    "lfilepath":"./",  //log file path
    "lpollerinterval": //500 millisecond flush once
}
```
## start
```
package main

import (
    "github.com/kafrax/logx"
)

func main(){
    logx.Debugf("module=test |message=%s","logx is a lightweight log to use")
    var str string
    fmt.Scan(&str)
}
```

# future
 - data queue send to kafaka

# @me
 - kafrax.go@gmail.com