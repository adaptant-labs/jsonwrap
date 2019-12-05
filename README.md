# JSONWrap

`jsonwrap` provides a simple API for JSON object nesting in Go.

A convenience CLI app implementing the API and capable of forwarding the wrapped JSON to a remote HTTP endpoint is
also provided.

## CLI Quick Start

Using the default Consul service catalog as an example:

```shell script
$ curl -X GET localhost:8500/v1/catalog/services
{
    "consul": []
}
```

This can be nested under as many levels as needed, simply repeat the `--wrap` flag for each level:

```shell script
$ curl -s -X GET localhost:8500/v1/catalog/services | jsonwrap --wrap services --wrap consul
{
    "consul": {
        "services": {
            "consul": []
        }
    }
}
```

The wrapped output can, in turn, be forwarded to a remote HTTP endpoint:

```shell script
$ curl -s -X GET localhost:8500/v1/catalog/services | \
  jsonwrap --wrap services --wrap consul --method PUT --target http://localhost:8181/v1/data
```

## Usage

```shell script
$ jsonwrap --help
NAME:
   jsonwrap - CLI for JSON object nesting and HTTP forwarding

USAGE:
    [global options] command [command options] <JSON file>

VERSION:
   0.0.1

DESCRIPTION:
   A convenience CLI tool for nesting JSON objects and forwarding the result to a remote HTTP endpoint

AUTHOR:
   Adaptant Labs <labs@adaptant.io>

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --target value  Target URL to forward wrapped JSON to
   --method value  Forwarding method to use (default: "POST")
   --wrap value    Name of object to nest under
   --help, -h      show help
   --version, -v   print the version

COPYRIGHT:
   (c) 2019 Adaptant Solutions AG
```

## API Documentation

Online API documentation is provided through `godoc`, and can be accessed directly on the
[package entry](https://godoc.org/github.com/adaptant-labs/jsonwrap) in the godoc package repository.

## Features and bugs

Please file feature requests and bugs at the [issue tracker][tracker].

[tracker]: https://github.com/adaptant-labs/jsonwrap/issues

## License

Licensed under the terms of the Apache 2.0 license, the full version of which can be found in the
[LICENSE](https://raw.githubusercontent.com/adaptant-labs/jsonwrap/master/LICENSE) file included in the
distribution.
