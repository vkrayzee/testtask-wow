## Words of Wisdom TCP server

This is a simple TCP server that returns a random quote from a list of quotes.

Server is protected by a Proof of Work system (hashcash), so you need to solve a puzzle before you can get a quote.

Quotes are stored in a SQLite database.

### Usage (with Docker)

#### Server
```bash
$ docker build -t wow-server -f Dockerfile.server .
$ docker run -p 8080:8080 wow-server
```

#### Client
```bash
$ docker build -t wow-client -f Dockerfile.client .
$ docker run -it wow-client
```

or

```bash
$ docker run -it wow-client --host=host.docker.internal --port=8080
```

### Usage (without Docker)

#### Server
```bash
$ go run cmd/server/main.go
```

#### Client
```bash
$ go run cmd/client/main.go
```

### Protocol

#### Server -> Client

##### Challenge
```
<challenge-string>
```

#### Client -> Server

##### Solution
```
<solution>
```

#### Server -> Client

##### Quote
```
<quote>
```

