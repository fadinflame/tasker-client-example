# Tasker Example Client

> Implementation of gRPC & HTTP clients for [Tasker Go](https://github.com/fadinflame/tasker-go)

## Build

```
go build .
```

## Run

```
go run .
```

## Usage

First, compile or download [Tasker Go](https://github.com/fadinflame/tasker-go)

Then, run this package. That's all :)

If you did everything well, you will see this console log:

```
[gRPC] Task Create Success
[gRPC] Task Get Success
[gRPC] Task Update Success
[gRPC] Task Delete Success
gRPC client tests success: true

[HTTP] Task Create Success
[HTTP] Task Get Success
[HTTP] Task Update Success
[HTTP] Task Delete Success
HTTP client tests success: true
```