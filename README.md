# httpshare

[![Go Reference](https://pkg.go.dev/badge/github.com/9illes/httpshare.svg)](https://pkg.go.dev/github.com/9illes/httpshare)

Share a single file or a directory over http

## Usage

### Binary

```sh
go install github.com/9illes/httpshare

# share files in current directory on :3000
httpshare -d . -p 3000

# share a single file
httpshare -d ./file.ext
```

### Dev

Build

```sh
make build
```

Share a file on the default port (:8080)

```sh
go run . -d ../path/to/a/file.ext
```

Share all file of a directory

```sh
go run . -d ../path/to/a/directory
```

## Troubleshooting

Ensure that the port is open.

```sh
firewall-cmd --get-active-zones
firewall-cmd --zone=<zone_name> --permanent --add-port=<port>
```
