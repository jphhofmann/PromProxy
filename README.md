# PromProxy

## Motivation

Most prometheus exporters does not provide any kind of request source verification. Means for example, if you are running prometheus node-exporter in a non-firewalled environment, metrics will be exposed to the public.

Usually, no one wants to provide potential security related insights to everyone. 

## How PromProxy helps

With PromProxy, you can run prometheus exporters on localhost and allow connections only from authorized ip-addresses, e.g. your very nice prometheus cluster.

## Stability

This is currently a PoC I have quickly written on a cold April evening, while being too lazy to work on other projects.

## Documentation

### Requirements

- Go >=1.16.3
- Git ;)

### Dependencies

Please see go.mod. However, notable external dependencies are:

- [logrus](https://github.com/sirupsen/logrus)
- [fasthttp](https://github.com/valyala/fasthttp)
- [fasthttp-router](https://github.com/fasthttp/router)
- [fasthttp-reverse-proxy](https://github.com/yeqown/fasthttp-reverse-proxy)

#### Why fasthttp and not net/http?

I'm using fasthttp for several projects at [combahton](https://www.combahton.net. It's not about performance, rather about standardization.

### Building

```git clone https://github.com/jphhofmann/promproxy```

Afterwards build the package:

```cd promproxy && go build```

A new binary named "promproxy" appears. Just move the binary to /usr/bin/promproxy. Afterwards, you can register a systemd service, see the promproxy.service unit file.

### Configuration

Please see the promproxy.yaml configuration file. Yaml parsing happens using the yaml.v2 package. Configuration is by default read from /etc/promproxy.yaml.

#### Listen (string)

Provide the listen address for fasthttp, e.g. 0.0.0.0:10000.

#### Debug (bool)

Adds debug output to stdout, e.g. unauthorized connections and fasthttp-reverse-proxy debug output.

#### Routes (map)

Registers routes with promproxy. Please note, the sequence name needs to match the location name currently. This is a not so good design, I didnt address yet.

#### Whitelist (map)

The whitelisted ip-address provided as yaml sequence.