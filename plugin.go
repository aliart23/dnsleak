package dnsleak

import (
	"context"
	"fmt"
	"github.com/coredns/coredns/plugin"
	"github.com/coredns/coredns/request"
	"github.com/go-redis/redis/v9"
	"github.com/miekg/dns"
	"time"
)

type Plugin struct {
	client *redis.Client
	next   plugin.Handler
}

func (p Plugin) ServeDNS(ctx context.Context, writer dns.ResponseWriter, msg *dns.Msg) (int, error) {
	state := request.Request{W: writer, Req: msg}

	p.client.SetEx(
		ctx,
		fmt.Sprintf("dnsleak:ip:%s", state.QName()),
		state.IP(),
		60*time.Second,
	)

	return plugin.NextOrFailure(p.Name(), p.next, ctx, writer, msg)
}

func (p Plugin) Name() string {
	return "dnsleak"
}
