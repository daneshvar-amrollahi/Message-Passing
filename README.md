# Message-Broker

This is a simple implementation of a message-broker in Go for the following cases:

1. one client - one server - sync

2. one client - one server - async

3. one client - one server - with overflow handling

4. one client - one server - two way message passing

5. three clients - one server

To run the code for each case, you have to enter the corresponding directory in the ```src/``` directory.
Here is an overview of the project structure:
```
├── Documentation
│   ├── Description.pdf
│   └── Report-810197685.pdf
├── go.mod
├── go.sum
├── logger.log
├── README.md
└── src
    ├── one-one
    │   ├── async
    │   │   └── main.go
    │   ├── overflow
    │   │   └── main.go
    │   ├── sync
    │   │   └── main.go
    │   └── two-way
    │       └── main.go
    └── one-three
        └── main.go
```

and then all you need to do is enter ```go run main.go```