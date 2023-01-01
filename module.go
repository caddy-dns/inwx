package inwx

import (
	"github.com/caddyserver/caddy/v2"
	"github.com/caddyserver/caddy/v2/caddyconfig/caddyfile"
	libdnsinwx "github.com/libdns/inwx"
)

// Provider lets Caddy read and manipulate DNS records hosted by this DNS provider.
type Provider struct{ *libdnsinwx.Provider }

func init() {
	caddy.RegisterModule(Provider{})
}

// CaddyModule returns the Caddy module information.
func (Provider) CaddyModule() caddy.ModuleInfo {
	return caddy.ModuleInfo{
		ID:  "dns.providers.inwx",
		New: func() caddy.Module { return &Provider{new(libdnsinwx.Provider)} },
	}
}

// Before using the provider config, resolve placeholders.
// Implements caddy.Provisioner.
func (p *Provider) Provision(ctx caddy.Context) error {
	p.Provider.Username = caddy.NewReplacer().ReplaceAll(p.Provider.Username, "")
	p.Provider.Password = caddy.NewReplacer().ReplaceAll(p.Provider.Password, "")
	p.Provider.SharedSecret = caddy.NewReplacer().ReplaceAll(p.Provider.SharedSecret, "")
	p.Provider.EndpointURL = caddy.NewReplacer().ReplaceAll(p.Provider.EndpointURL, "")

	return nil
}

// UnmarshalCaddyfile sets up the DNS provider from Caddyfile tokens. Syntax:
//
// inwx {
//     username <username>
//     password <password>
//     shared_secret <shared_secret>
//     endpoint_url <endpoint_url>
// }
func (p *Provider) UnmarshalCaddyfile(d *caddyfile.Dispenser) error {
	for d.Next() {
		if d.NextArg() {
			return d.ArgErr()
		}
		for nesting := d.Nesting(); d.NextBlock(nesting); {
			switch d.Val() {
			case "username":
				if p.Provider.Username != "" {
					return d.Err("Username already set")
				}
				if d.NextArg() {
					p.Provider.Username = d.Val()
				}
				if d.NextArg() {
					return d.ArgErr()
				}
			case "password":
				if p.Provider.Password != "" {
					return d.Err("Password already set")
				}
				if d.NextArg() {
					p.Provider.Password = d.Val()
				}
				if d.NextArg() {
					return d.ArgErr()
				}
			case "shared_secret":
				if p.Provider.SharedSecret != "" {
					return d.Err("Shared secret already set")
				}
				if d.NextArg() {
					p.Provider.SharedSecret = d.Val()
				}
				if d.NextArg() {
					return d.ArgErr()
				}
			case "endpoint_url":
				if p.Provider.EndpointURL != "" {
					return d.Err("Endpoint URL already set")
				}
				if d.NextArg() {
					p.Provider.EndpointURL = d.Val()
				}
				if d.NextArg() {
					return d.ArgErr()
				}
			default:
				return d.Errf("unrecognized subdirective '%s'", d.Val())
			}
		}
	}

	if p.Provider.Username == "" {
		return d.Err("missing username")
	}

	if p.Provider.Password == "" {
		return d.Err("missing password")
	}

	return nil
}

// Interface guards
var (
	_ caddyfile.Unmarshaler = (*Provider)(nil)
	_ caddy.Provisioner     = (*Provider)(nil)
)
