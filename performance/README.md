# Performance

The performance directory contains some of the things that need to be taken into account when working on performant code.

Every directory here contains a couple of benchmarks demonstrating the performance of similar solutions that solve the same problem.

Run the benchmarks by executing `go test -gcflags '-l' -bench=.`.

The [`demo`](demo) directory contains a demo application that demonstrates a RabbitMQ integration and processing using workers.