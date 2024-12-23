INWX module for Caddy
=====================

[![Go Reference](https://pkg.go.dev/badge/github.com/caddy-dns/inwx.svg)](https://pkg.go.dev/github.com/caddy-dns/inwx)

This package contains a DNS provider module for [Caddy](https://github.com/caddyserver/caddy). It can be used to manage DNS records with [INWX](https://www.inwx.de/en).

## Caddy module name

```
dns.providers.inwx
```

## Config examples

To use this module for the ACME DNS challenge, [configure the ACME issuer in your Caddy JSON](https://caddyserver.com/docs/json/apps/tls/automation/policies/issuer/acme/) like so:

```json
{
	"module": "acme",
	"challenges": {
		"dns": {
			"provider": {
				"name": "inwx",
				"username": "INWX_USERNAME",
				"password": "INWX_PASSWORD",
				"shared_secret": "INWX_SHARED_SECRET",
				"endpoint_url": "INWX_ENDPOINT_URL"
			}
		}
	}
}
```

or with the Caddyfile:

```
# globally
{
	acme_dns inwx {
		username <username>
		password <password>
		shared_secret <shared_secret>
		endpoint_url <endpoint_url>
	}
}
```

```
# one site
tls {
	dns inwx {
		username <username>
		password <password>
		shared_secret <shared_secret>
		endpoint_url <endpoint_url>
	}
}
```

If you don’t provide an `endpoint_url` the URL of the production environment (https://ote.domrobot.com/jsonrpc/) is used by default. If you want to use the test environment, set `endpoint_url` to https://api.ote.domrobot.com/jsonrpc/.

For more information on the configuration, see [libdns/inwx](https://github.com/libdns/inwx).
