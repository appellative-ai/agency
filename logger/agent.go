package logger

import (
	"github.com/appellative-ai/agency/logx"
	"github.com/appellative-ai/collective/exchange"
	"github.com/appellative-ai/core/messaging"
	"github.com/appellative-ai/core/rest"
	"github.com/appellative-ai/core/std"
	"net/http"
	"time"
)

const (
	AgentName    = "common:resiliency:agent/logging/request"
	defaultRoute = "host"
)

// AgentT - agent
type AgentT interface {
	messaging.Agent
	LogEgress(start time.Time, duration time.Duration, route string, req any, resp any, timeout time.Duration)
	LogStatus(name string, status *std.Status)
}

var (
	Agent AgentT
)

// init - register an agent constructor
func init() {
	exchange.RegisterConstructor(AgentName, func() messaging.Agent {
		return newAgent()
	})
	Agent = newAgent()
	exchange.Register(Agent)
}

type agentT struct {
	operators []logx.Operator
}

func newAgent() *agentT {
	return new(agentT)
}

func (a *agentT) Name() string { return AgentName }
func (a *agentT) Message(m *messaging.Message) {
	if m == nil {
		return
	}
	switch m.Name {
	case messaging.ConfigEvent:
		a.configure(m)
		return
	}
}

// Link - chainable exchange
func (a *agentT) Link(next rest.Exchange) rest.Exchange {
	return func(r *http.Request) (resp *http.Response, err error) {
		start := time.Now().UTC()
		resp, err = next(r)
		logx.LogAccess(a.operators, logx.IngressTraffic, start, time.Since(start), defaultRoute, r, resp)
		return
	}
}

func (a *agentT) LogEgress(start time.Time, duration time.Duration, route string, req any, resp any, timeout time.Duration) {
	logx.LogEgress(a.operators, start, duration, route, req, resp, timeout)
}

func (a *agentT) LogStatus(name string, status *std.Status) {
	logx.LogStatus(nil)
}
