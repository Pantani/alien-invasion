# Alien Invasion

Mad​ ​aliens​ ​are​ ​about​ ​to​ ​invade​ ​the​ ​earth​ ​and​ ​you​ ​are​ ​tasked​ ​with​ ​simulating​ ​the invasion.

## Setup

### Requirements

 * [Go Toolchain](https://golang.org/doc/install) versions 1.14+

### From Source

```shell
$ make
// OR
$ go run main.go
```

### Docker

Build and run from local Dockerfile:

```shell
$ docker-compose up
```

## Arguments
```shell
-w string
    world file name (default "test/world_1.txt").
-a int
    number of aliens (default 10).
-i int
    number of iterations (default 10000).
```

E.g.:
```shell
$ go run main.go -w test/world_2.txt

$ go run main.go -w test/world_2.txt -a 20 -i 1000

$ go run main.go -w test/world_2.txt -a 2 -i 10

$ go run main.go -a 20 -i 1000
```

All test files are located inside the `test` folder.

## Useful Commands

### Unit tests
```shell
$ make test
```

### Help

```shell
$ make help
```

