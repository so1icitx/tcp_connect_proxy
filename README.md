
# TCP CONNECT Proxy

A simple HTTP CONNECT proxy written in Go.

## How it works

Accepts CONNECT requests from HTTP clients, establishes a TCP tunnel 
to the destination server, and forwards raw bytes bidirectionally 

## Usage

Run the proxy:
```
go run main.go
```

Then point your client at it:
```
curl.exe -x http://127.0.0.1:8080 https://example.com
```

## Requirements

- Go 1.21 or higher

## Limitations

- Only handles HTTPS traffic via HTTP CONNECT
- Domain parsing is minimal and assumes standard CONNECT format
