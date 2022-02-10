# spanchk
Tiny utility to check whether a validator is in a heimdall span

## Build
```
go build
mv spanchk /usr/bin
```
## Usage
```
spanchk serve -v "0x8800000c1459d224383d000000cedd1fcee20000"
```

Options:
```
Serve queries

Usage:
  spanchk serve [flags]

Flags:
  -h, --help                   help for serve
  -a, --hmdaddr string         Heimdall endpoint to query for span (default "127.0.0.1:1317")
  -l, --listenaddr string      Endpoint to serve queries (default "127.0.0.1:13180")
  -v, --validatoraddr string   Validator to inform about (default "0x0")
```