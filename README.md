# PromProxy

## Motivation
---

Most prometheus exporters does not provide any kind of request source verification. Means for example, if you are running prometheus node-exporter in a non-firewalled environment, metrics will be exposed to the public.

Usually, no one wants to provide potential security related insights to everyone. 

## How PromProxy helps
---

With PromProxy, you can run prometheus exporters on localhost and allow connections only from authorized ip-addresses, e.g. your prometheus cluster.

## Stability
---

This can be now considered as stable and is actively used within the infrastructure of combahton GmbH on > 80 servers.

## Documentation
---

### Requirements
---

Requires Go 1.16, see go.mod for dependencies.

### Dependencies
---

Please see go.mod. However, notable external dependencies are:

- [logrus](https://github.com/sirupsen/logrus)
- [fasthttp](https://github.com/valyala/fasthttp)
- [fasthttp-router](https://github.com/fasthttp/router)
- [fasthttp-reverse-proxy](https://github.com/yeqown/fasthttp-reverse-proxy)

#### Why fasthttp and not net/http?
---

I'm using fasthttp for several projects at [combahton](https://www.combahton.net). We use fasthttp as standard for several software projects, e.g. customer facing APIs and synchronization of DDoS-Analyzation.

### Building
---

```git clone https://github.com/jphhofmann/PromProxy```

Afterwards build the package:

```cd promproxy && go build```

A new binary named "promproxy" appears. Just move the binary to /usr/bin/promproxy. Afterwards, you can register a systemd service, see the promproxy.service unit file.

### Configuration
---

Please see the promproxy.yaml configuration file. Yaml parsing happens using the yaml.v2 package. Configuration is by default read from /etc/promproxy.yaml.

#### Listen (string)
---

Provide the listen address for fasthttp, e.g. 0.0.0.0:10000.

#### Debug (bool)
---

Adds debug output to stdout, e.g. unauthorized connections and fasthttp-reverse-proxy debug output.

#### Routes (map)
---

Registers routes with promproxy. An entry represents a route in the format of /ENTRY.

#### Whitelist (map)
---

The whitelisted ip-address provided as yaml sequence.