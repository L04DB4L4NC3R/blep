# Mormon
A high availability log sourcing tool for modern DevOps

<br />

### Our implementation
---

- [X] Decentralizing calls to a distributed log-store, `rqlite`
- [X] Fast logs sourcing using `NATS` and `gRPC`
- [X] Low bottlenecking
- [X] Abstraction between persistent (subscriber) and application (publisher) layer
- [X] Easily deployable and testable image build
- [X] Easy integration with an existing ELK stack
- [X] Seamless integration across different codebases
- [X] Self contained binaries
- [X] Inbuilt analytics and benchmarks
- [X] `Mormon` CLI interface

<br />

### Why NATS?
---
[NATS](https://github.com/nats-io/go-nats.git) is an event sourcing tool which we will be using to publish logs and distribute related data between different services. The reason for NATS is:

* RabbitMQ is limited to HTTP and HTTPS natively

* NATS supports gRPC and works fast due to being type safe as well as removing extra overhead of marshalling and unmarshalling

* Publishing events on a different thread and subscribing from a different thread allows non-blocking log sourcing.

* NATS is very high throughput when it comes to requests per second

<br />

![NATS](./images/nats.png)

<br />
<br />

### The need for a distributed DB
---
Different containers consuming from a single DB lead to network bottlenecking. distributed databases give us the advantage in terms of reducing request overhead. In this implementation we have used [rqlite](https://github.com/rqlite/rqlite.git), because:

* It is very lightweight, and written in go

* Forming a cluster is very easy.    

* It is fast and fault tolerant

* It has an inbuilt backup and recovery mechanism


<br />

![rqlite](./images/rqlite.png)


<br />
<br />


## Introducing the mormon CLI

```
usage: mormon [subcommand] [args]
mormon build <----------------------------------------------------------> build from source, go needed
mormon run-env <--------------------------------------------------------> run the nats and rqlite env
mormon run-safe-env <---------------------------------------------------> run-env for the case docker-compose is not present
mormon run-env stop <---------------------------------------------------> stop the nats and rqlite env
mormon start -p [port] <------------------------------------------------> stop the nats and rqlite env
mormon publisher [logId] [timestamp] [log] [host] <---------------------> publish an event
mormon subscriber logs.[host] [callback executible] <-------------------> subscribe for an event
mormon mock-logs <------------------------------------------------------> generate logs for testing
```

<br />


## Project instructions
Build environment specification

<br/>

#### Run NATS and rqlite

```
git clone https://angadsharma1016/mormon
cd mormon
chmod +x bin/*
./bin/build
./bin/run-env
```

#### Stop both 

```
./bin/run-env stop
```

<br/>

#### Run server

```
./bin/start -p 3000
```

<br/>


#### Publisher side

```
./bin/publisher "[log ID]" "[Timestamp]" "[Log]" "[Host]"
```

<br/>

#### Subscriber side
On the occurence of the event, subscriber will run `./bin/callback`

```
./bin/subscriber "[event name]" "[callback executible]"
```

<br/>

<iframe width="560" height="315" src="https://www.youtube.com/embed/u2-FNYZKUsE" frameborder="0" allow="accelerometer; autoplay; encrypted-media; gyroscope; picture-in-picture" allowfullscreen></iframe>


