# Example: Connect fo Go and Scenarigo

## Prerequisites

- [Buf CLI](https://buf.build/docs/installation)
- [gRPCurl](https://github.com/fullstorydev/grpcurl)
- [Evans](https://github.com/ktr0731/evans)

You can install these prerequisites through [asdf](https://asdf-vm.com):

```shell
asdf install
```

## Clients

### grpcurl

Without reflection:

```shell
grpcurl -protoset <(buf build -o -) -plaintext -d '{"name": "Jane"}' localhost:8080 greet.v1.GreetService/Greet
```

With reflection:

```shell
grpcurl -plaintext -d '{"name": "Jane"}' localhost:8080 greet.v1.GreetService/Greet
```

### Go

```shell
go run ./cmd/client/main.go
```

### Evans

```shell
$ evans --reflection --host 127.0.0.1 --port 8080 repl
greet.v1.GreetService@127.0.0.1:8080> show package
greet.v1.GreetService@127.0.0.1:8080> show service
greet.v1.GreetService@127.0.0.1:8080> show message
greet.v1.GreetService@127.0.0.1:8080> call Greet
name (TYPE_STRING) => Jane
{
  "greeting": "Hello, Jane!"
}
greet.v1.GreetService@127.0.0.1:8080> exit
```
