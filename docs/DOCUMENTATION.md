## Documentation specification

<br/>

#### Publisher side

```go
package main

import mormon "github.com/angadsharma1016/c2c/publisher"

func main() {
    logs := pb.LogStore{
        "log_id": "1";
        "timestamp": "20/12/2018";
        "log": "successfully completed a function";
        "host": "publisher svc";
    }
    go mormon.PublishLogs(&logs)
}
```


<br/>
<br/>

#### Subscriber side

```go
package main

import mormon "github.com/angadsharma1016/c2c/subscriber"

func 

func main() {
    go mormon.Subscribe("logs.>", fn)
}
```