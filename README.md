# teler (WAF) Proxy

An adapter for Envoy, Istio, Nginx, and other platforms, enabling seamless integration with teler WAF to protect against a variety of web-based attacks, such as cross-site scripting (XSS), SQL injection, known vulnerabilities, exploits, malicious actors, botnets, unwanted crawlers or scrapers, and directory bruteforce attacks.

**See also:**

* [kitabisa/teler](https://github.com/kitabisa/teler): Real-time HTTP Intrusion Detection.
* [kitabisa/teler-waf](https://github.com/kitabisa/teler-waf): Go HTTP middleware that provides teler IDS functionality.

**Table of Contents**

* [Architecture](#architecture)
* [Install](#installation)
  * [from Source](#source)
  * [with Docker](#docker)
* [Usage](#usage)
  * [Options](#options)

## Architecture

<img width="40%" src="https://github.com/kitabisa/teler-proxy/assets/25837540/5474b8e3-b8f7-4443-8775-f0a250eb3eb0">

## Installation

### Source

Using [Go](https://golang.org/doc/install) (v1.19+) compiler:

```bash
CGO_ENABLED=1 go install github.com/kitabisa/teler-proxy/cmd/teler-proxy@latest
```

### — or

Manual building executable from source code:

```bash
git clone https://github.com/kitabisa/teler-proxy.git
cd teler-proxy/
make build
```

### Docker

Pull the [Docker](https://docs.docker.com/get-docker/) image by running:

```bash
docker pull ghcr.io/kitabisa/teler-proxy:latest
```

## Usage

Simply, `teler-proxy` can be run with:

```bash
teler-proxy -d <ADDR>:<PORT> [OPTIONS...]
```

### Options

Here are all the options it supports.

```bash
teler-proxy -h
```

|          **Flag**           |                            **Description**                              | **Example**                                                                     |
|:--------------------------: |:---------------------------------------------------------------------:  |-------------------------------------------------------------------------------  |
| -p, --port `<PORT>`         | Set the local port to listen on **(default: 1337)**                     | `teler-proxy -p 8000 -d localhost:80`                                           |
| -d, --dest `<ADDR>:<PORT>`  | Set the destination address for forwarding requests                     | `teler-proxy -d localhost:80`                                                   |
| -c, --conf `<FILE>`         | Specify the path to the teler WAF configuration file                    | `teler-proxy -d localhost:80 -c /path/to/teler-waf.conf.yaml`                   |
| -f, --format `<FORMAT>`     | Specify the configuration file format (json/yaml) **(default: yaml)**   | `teler-proxy -d localhost:80 -c /path/to/teler-waf.conf.json -f json`           |
| --cert `<FILE>`             | Specify the path to the SSL certificate file                            | `teler-proxy -d localhost:80 --cert /path/to/cert.key --key /path/to/key.pem`   |
| --key `<FILE>`              | Specify the path to the SSL private key file                            | `teler-proxy -d localhost:80 --cert /path/to/cert.key --key /path/to/key.pem`   |
| -V, --version               | Display the current teler-proxy version                                 | `teler-proxy -V`                                                                |
| -h, --help                  | Display this helps text                                                 | `teler-proxy -h`                                                                |

## License

This program is developed and maintained by members of Kitabisa Security Team, and this is not an officially supported Kitabisa product. This program is free software: you can redistribute it and/or modify it under the terms of the [Apache-2.0 license](/LICENSE). Kitabisa teler-proxy and any contributions are copyright © by Dwi Siswanto 2023.