## Documentation specification

<br/>

#### Publisher side

```go
package main

import mormon "github.com/angadsharma1016/c2c/publisher"

func main() {
    logs := pb.LogStore{
        LogId: "1",
        Timestamp: "20/12/2018",
        Log: "successfully completed a function",
        Host: "publisher svc",
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

func main() {
    go mormon.Subscribe("logs.>")
}
```

<br/>
<br/>

#### Command line

```
./bin/publisher "[log ID]" "[Timestamp]" "[Log]" "[Host]" 
```