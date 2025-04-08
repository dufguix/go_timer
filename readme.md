# Simple timer for GO and TinyGO

## Example
```go
//Create timer and reset it
var myTimer = Timer{}
myTimer.Reset()

//Check if timer is elapsed
//method 1
if myTimer.RefreshAndCheck(time.Second * 2) {
    //do something
    myTimer.Reset()
}

//method 2
myTimer.Refresh()
if myTimer.Check(time.Second * 2) {
    //do something
    myTimer.Reset()
}
```

## Install 

```
go get github.com/dufguix/go_timer

import (
	"time"
	timer "github.com/dufguix/go_timer"
)
```