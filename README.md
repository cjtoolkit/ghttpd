[![Build Status](https://travis-ci.com/cjtoolkit/ghttpd.svg?branch=master)](https://travis-ci.com/cjtoolkit/ghttpd)

# GHTTPD

Lightweight HTTP Server for Static Files.

It's useful for experimenting with JavaScript and even WebAssembly, without having to run a full on http
server such as Apache (Oh gosh) and Nginx (Better than former, but still overkill for simple experiments).

# Installation

Must be using at least go 1.11, has not been tested with earlier versions of go.

```sh
go get github.com/cjtoolkit/ghttpd
```

# Usage

Just execute `ghttpd` in your wd (working directory), it will host the server for wd with
default configuration see below.

# Configuration

Create a new file called `ghttpd.toml` and place that in wd (so it's easy to save to repo as well), if you
are happy with the default setting you don't need to create the file, it will use the default setting =)

```toml
# All of below are optional.

[http]
debug = false # If true logs every request to terminal, default is false.
address = ":8080" # Bind server to address, default is ":8080"
cacheTime = 3600 # Cache Control Max Age, default is 3600

# Runs server in tls mode, don't need tls leave this out, Default is nil
[tls]
cert = "/path/to/cert.pem"
key = "/path/to/private.pem"

# Adds mime type to server, default is nil.
[[mime]]
extension = ".wasm" # Must begin with .
type = "application/wasm"

# You can add as many mime type as you want.
[[mime]]
extension = ".jpg"
type = "image/jpeg"
```