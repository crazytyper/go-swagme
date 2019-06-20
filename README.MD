# go-swagme

Generates go code from swagger spec.

It generate types and enums only.

It neither generates a HTTP client nor a HTTP server.

This project is an unfinished prototype. It succeeds with my swagger specs but will very likely fail with yours.

## Getting Started

```bash
go get -u github.com/crazytyper/go-swagme/cmd/swagme
```

## Usage

```bash
swagme -p mypackagename myswagger.json | gofmt > mypackagename.go
```

