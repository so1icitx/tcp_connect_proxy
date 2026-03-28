
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
### Using with a Browser

#### Firefox
1. Open Settings
2. Search for "proxy"
3. Click **Network Settings**
4. Select **Manual proxy configuration**
5. Set HTTP Proxy to `127.0.0.1` and port to `8080`
6. Check **Also use this proxy for HTTPS**
7. Click OK

#### Chrome
Chrome uses the OS proxy settings, so set it there:

1. Open Settings
2. Search for "proxy"
3. Click **Set up a proxy server** under Manual proxy setup
4. Set address to `127.0.0.1` and port to `8080`

> this is for windows if your on linux im sure you dont use a chromium browser

## Requirements

- Go 1.21 or higher

## Limitations

- Only handles HTTPS traffic via HTTP CONNECT
- Domain parsing is minimal and assumes standard CONNECT format
