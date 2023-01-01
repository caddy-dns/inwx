INWX module for Caddy
=====================

This package contains a DNS provider module for [Caddy](https://github.com/caddyserver/caddy). It can be used to manage DNS records with INWX.

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
		endpoint_url <endpoint_url>"
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
		endpoint_url <endpoint_url>"
	}
}
```
