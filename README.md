# logx
 lightweight log tool, dynamic configuration,automatic in write and switch log files

 ```
    go get github.com/kafrax/logx
 ```

# tps
```
    	//OutPut
    	//tps is : 703401/s
```

# simple to use
```
package main

import (
    "fmt"
	"github.com/kafrax/logx"
)

func main(){
    logx.Debugf("module=test |message=%s","logx is a lightweight log to use")
    var str string
    fmt.Scan(&str)
}

```
it will be create dir with ./logx

# feature
 - only write to file
 - at midnight auto do file switch

# future
 - dynamic config
 - os.stdout output
 - data queue send to kafaka

