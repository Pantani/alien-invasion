# Alien Invasion

Mad​ ​aliens​ ​are​ ​about​ ​to​ ​invade​ ​the​ ​earth​ ​and​ ​you​ ​are​ ​tasked​ ​with​ ​simulating​ ​the invasion.

## Setup

### Requirements

 * [Go Toolchain](https://golang.org/doc/install) versions 1.14+

### From Source

```shell
make start
```

### Docker

Build and run from local Dockerfile:

```shell
docker-compose up
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

## Useful Commands

### Unit tests
```shell
make test
```

### Help

```shell
$ make help

 Choose a command run in alien-invasion:

  install            Install missing dependencies. Runs `go get` internally. e.g; make install get=github.com/foo/bar
  start              Start all applications in development mode.
  start-simulation   Start alian simulation.
  stop               Stop development mode.
  compile            Compile the project.
  exec               Run given command. e.g; make exec run="go test ./..."
  clean              Clean build files. Runs `go clean` internally.
  test               Run all tests.
  unit               Run all unit tests.
  fmt                Run `go fmt` for all go files.
  govet              Run go vet.
  golint             Run golint.
```

