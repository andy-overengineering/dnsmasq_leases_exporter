# dnsmasq_leases_exporter
HTTP Exporter for DHCP Leases of dnsmasq exposes all current leases of dnsmasq over http in json format.

## Installation:

1. Clone repo:

```shell
git clone git@github.com:andy-overengineering/dnsmasq_leases_exporter.git
```

2. Build

```shell
cd dnsmasq_leases_exporter && go build
```


## Usage:

```shell
./dnsmasq_leases_exporter [-listen 0.0.0.0:9154] [-leases_path /var/lib/misc/dnsmasq.leases]
```

`-listen <address:port>` specifies the address to listen on

`-leases_path /some/file/path` specifies the path to the leases file from dnsmasq
