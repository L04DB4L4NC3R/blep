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
- [X] Abstraction between persistent and application layer
- [X] Seamless integration across different codebases

<br />

### Why NATS?
---
NATS is an event sourcing tool which we will be using to publish logs and distribute related data between different services. The reason for NATS is:

* RabbitMQ is limited to HTTP and HTTPS natively

* NATS supports gRPC and works fast due to being type safe as well as removing extra overhead of marshalling and unmarshalling

* Publishing events on a different thread and subscribing from a different thread allows non-blocking log sourcing.

* NATS is very high throughput when it comes to requests per second

<br />

![NATS](./images/nats.png)

<br />
<br />

### Why orbitDB?
---
While integrating with the ELK stack, we want different elastic containers to feed from a single DB. But this will lead to network bottlenecking. orbitDB gives us the following advantages:

* It has a native support for immutable logs store

* It leverages peer to peer networking.

* It uses IPFS, and peer to peer networking reduces bottlenecking


<br />

![orbitDB layers](./images/orbitDB1.png)


<br />


