# Mormon
A high availability log sourcing tool for modern DevOps

<br />

### Our implementation
---

- [X] A production ready package
- [X] Easily deployable and testable image build
- [X] Easy integration with an existing ELK stack
- [X] Decentralizing calls to log-store
- [X] Fast logs sourcing
- [X] Low bottlenecking
- [X] Distribute logstash flows from multiple elastic containers
- [X] Abstraction between persistent (subscriber) and application (publisher) layer
- [X] Seamless integration across different codebases

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

![orbitDB layers](./images/rqlite.png)


<br />


