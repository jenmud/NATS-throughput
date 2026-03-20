# NATS-throughput

This is a very simple demo for using NATS and will show basic throughput.


## Build

```bash
$ make build
```


## Examples

### Dev build with `go run`

```bash
# run with the defaults
# 1 Consumer
# 1 Publisher
# 1M message size
# for 5 seconds duration
$ make run
time=2026-03-21T00:03:55.907+11:00 level=INFO msg="starting up NATS server"
time=2026-03-21T00:03:55.908+11:00 level=INFO msg="NATS server running and accepting client connections" nats.server.address=nats://0.0.0.0:36295
time=2026-03-21T00:04:00.911+11:00 level=INFO msg="-> publish complete" nats.server.address=nats://0.0.0.0:36295 msgSize="1.00 MiB" duration=5s runtime=5.001931427s published=6011 msgPerSec=1202 totalSize="5.87 GiB" throughput="1.17 GiB/s"
time=2026-03-21T00:04:01.907+11:00 level=INFO msg="<- consumer complete" nats.server.address=nats://0.0.0.0:36295 clientID=0 msgSize="1.00 MiB" nats.server.address=nats://0.0.0.0:36295 runtime=5.99835194s received=6011 msgPerSec=1202 totalSize="5.87 GiB" throughput="1002.11 MiB/s"
```

```bash
# run with the defaults
# 5 Consumer
# 1 Publisher
# 1M message size
# for 5 seconds duration
$ make CONSUMERS=5 run
time=2026-03-21T00:07:01.195+11:00 level=INFO msg="starting up NATS server"
time=2026-03-21T00:07:01.196+11:00 level=INFO msg="NATS server running and accepting client connections" nats.server.address=nats://0.0.0.0:37531
time=2026-03-21T00:07:06.205+11:00 level=INFO msg="-> publish complete" nats.server.address=nats://0.0.0.0:37531 msgSize="1.00 MiB" duration=5s runtime=5.008388458s published=2247 msgPerSec=449 totalSize="2.19 GiB" throughput="448.65 MiB/s"
time=2026-03-21T00:07:07.194+11:00 level=INFO msg="<- consumer complete" nats.server.address=nats://0.0.0.0:37531 clientID=4 msgSize="1.00 MiB" nats.server.address=nats://0.0.0.0:37531 runtime=5.998433497s received=2247 msgPerSec=449 totalSize="2.19 GiB" throughput="374.60 MiB/s"
time=2026-03-21T00:07:07.194+11:00 level=INFO msg="<- consumer complete" nats.server.address=nats://0.0.0.0:37531 clientID=1 msgSize="1.00 MiB" nats.server.address=nats://0.0.0.0:37531 runtime=5.998443966s received=2247 msgPerSec=449 totalSize="2.19 GiB" throughput="374.60 MiB/s"
time=2026-03-21T00:07:07.195+11:00 level=INFO msg="<- consumer complete" nats.server.address=nats://0.0.0.0:37531 clientID=2 msgSize="1.00 MiB" nats.server.address=nats://0.0.0.0:37531 runtime=5.998475317s received=2247 msgPerSec=449 totalSize="2.19 GiB" throughput="374.60 MiB/s"
time=2026-03-21T00:07:07.195+11:00 level=INFO msg="<- consumer complete" nats.server.address=nats://0.0.0.0:37531 clientID=3 msgSize="1.00 MiB" nats.server.address=nats://0.0.0.0:37531 runtime=5.998377765s received=2247 msgPerSec=449 totalSize="2.19 GiB" throughput="374.60 MiB/s"
time=2026-03-21T00:07:07.195+11:00 level=INFO msg="<- consumer complete" nats.server.address=nats://0.0.0.0:37531 clientID=0 msgSize="1.00 MiB" nats.server.address=nats://0.0.0.0:37531 runtime=5.99860505s received=2247 msgPerSec=449 totalSize="2.19 GiB" throughput="374.59 MiB/s"
```

### Prod build with `make build`

```bash
# run with the defaults
# 1 Consumer
# 1 Publisher
# 1M message size
# for 5 seconds duration
$ ./builds/NATS-throughput
time=2026-03-21T00:08:46.146+11:00 level=INFO msg="starting up NATS server"
time=2026-03-21T00:08:46.147+11:00 level=INFO msg="NATS server running and accepting client connections" nats.server.address=nats://0.0.0.0:39409
time=2026-03-21T00:08:51.151+11:00 level=INFO msg="-> publish complete" nats.server.address=nats://0.0.0.0:39409 msgSize="1.00 MiB" duration=5s runtime=5.00311448s published=6049 msgPerSec=1209 totalSize="5.91 GiB" throughput="1.18 GiB/s"
time=2026-03-21T00:08:52.145+11:00 level=INFO msg="<- consumer complete" nats.server.address=nats://0.0.0.0:39409 clientID=0 msgSize="1.00 MiB" nats.server.address=nats://0.0.0.0:39409 runtime=5.998034332s received=6049 msgPerSec=1209 totalSize="5.91 GiB" throughput="1008.50 MiB/s"
```