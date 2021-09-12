# dd-remote

An efficient network service for direct writing to remote storage via HTTP.
A server application listens for incoming PUT requests and writes data to a predetermined path.
A client application reads from a local path and sends file content to a specified web URI.

dd-remote has been designed for minimal memory usage and will refuse to handle concurrent requests (first-come-first-served).

## Server

```
go get git.home.jm0.eu/josua/dd-remote/cmd/ddrd

ddrd --help
Usage:
  ddrd [OPTIONS]

Application Options:
  -l, --listen-host= address to listen on for incoming connections (default: localhost)
  -p, --listen-port= port to listen on for incoming connections (default: 80)
  -o, --output=      destination file for writing to (default: /dev/stdout)

Help Options:
  -h, --help         Show this help message
```

## Client

```
go get git.home.jm0.eu/josua/dd-remote/cmd/ddr

ddr --help
Usage:
  ddr [OPTIONS]

Application Options:
  -u, --uri=   dd-remote server address (default: http://localhost/)
  -i, --input= input file to send (default: /dev/stdin)

Help Options:
  -h, --help   Show this help message
```
