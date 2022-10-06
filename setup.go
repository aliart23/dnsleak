package dnsleak

import (
	"github.com/coredns/caddy"
	"github.com/coredns/coredns/core/dnsserver"
	"github.com/coredns/coredns/plugin"
	"github.com/go-redis/redis/v9"
)

func init() {
	plugin.Register("dnsleak", setup)
}

func setup(c *caddy.Controller) error {
	c.Next()

	if !c.NextArg() {
		return plugin.Error("dnsleak", c.ArgErr())
	}

	dnsserver.GetConfig(c).AddPlugin(func(next plugin.Handler) plugin.Handler {
		return Plugin{
			client: redis.NewClient(&redis.Options{
				Addr: c.Val(),
			}),
			next: next,
		}
	})

	return nil
}
