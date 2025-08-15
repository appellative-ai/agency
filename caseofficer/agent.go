package caseofficer

import (
	"fmt"
	"github.com/appellative-ai/collective/notification"
	"github.com/appellative-ai/core/messaging"
	"sync/atomic"
)

type Agent interface {
	messaging.Agent
	BuildNetwork(m []map[string]string) ([]any, []error)
	Operative(mame string) messaging.Agent
	Trace()
}

// TODO : need host name
type agentT struct {
	running atomic.Bool
	name    string

	ex       *messaging.Exchange
	notifier *notification.Interface
}

// NewAgent - create a new agent
func NewAgent(name string) Agent {
	return newAgent(name, notification.Notifier)
}

func newAgent(name string, notifier *notification.Interface) *agentT {
	a := new(agentT)
	a.running.Store(false)
	a.name = name
	a.notifier = notifier

	a.ex = messaging.NewExchange()
	return a
}

// Name - agent identifier
func (a *agentT) Name() string { return a.name }

func (a *agentT) Trace() {
	list := a.ex.List()
	for _, v := range list {
		fmt.Printf("trace: %v -> %v\n", a.Name(), v)
	}
}

// Message - message the agent
func (a *agentT) Message(m *messaging.Message) {
	if m == nil {
		return
	}
	switch m.Name {
	case messaging.ConfigEvent:
		a.config(m)
		return
	case messaging.StartupEvent:
		if a.running.Load() {
			return
		}
		a.running.Store(true)
		//a.run()
		//a.emissary.C <- m
		a.ex.Broadcast(m)
		return
	case messaging.ShutdownEvent:
		if !a.running.Load() {
			return
		}
		a.running.Store(false)
		//a.emissary.C <- m
		a.ex.Broadcast(m)
		return
	case messaging.PauseEvent, messaging.ResumeEvent:
		//a.emissary.C <- m
		a.ex.Broadcast(m)
		return
	}

	list := m.To()
	// No recipient, or only the case officer recipient
	if len(list) == 0 || (len(list) == 1 && list[0] == a.name) {
		/*
			switch m.Channel() {
			case messaging.ChannelEmissary:
				a.emissary.C <- m
			case messaging.ChannelControl:
				a.emissary.C <- m
			default:
				fmt.Printf("limiter - invalid channel %v\n", m)
			}

		*/
		return
	}
	// Send to appropriate agent
	a.ex.Message(m)
}

func (a *agentT) BuildNetwork(net []map[string]string) (operatives []any, errs []error) {
	return buildNetwork(a, net)
}

func (a *agentT) Operative(name string) messaging.Agent {
	return a.ex.Get(name)
}

/*
func (a *agentT) configure(m *messaging.Message) {
	switch m.ContentType() {
	case messaging.ContentTypeAgent:
		agent, status := messaging.AgentContent(m)
		if !status.OK() {
			messaging.Reply(m, status, a.Name())
			return
		}
		err := a.ex.Register(agent)
		if err != nil {
			messaging.Reply(m, messaging.NewStatus(messaging.StatusInvalidContent, err.Error()), a.Name())
		}
	}
	messaging.Reply(m, messaging.StatusOK(), a.Name())
}


*/
