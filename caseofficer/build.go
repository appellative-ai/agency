package caseofficer

import (
	"errors"
	"fmt"
	"github.com/appellative-ai/collective/exchange"
	"github.com/appellative-ai/collective/namespace"
	"github.com/appellative-ai/core/messaging"
	"github.com/appellative-ai/core/std"
)

func buildNetwork(a messaging.Agent, netCfg []map[string]string) (operatives []any, errs []error) {
	if a == nil {
		return nil, []error{errors.New("agent is nil")}
	}
	if len(netCfg) == 0 {
		return nil, []error{errors.New("network configuration is nil or empty")}
	}

	for _, m := range netCfg {
		name, ok := m[NameKey]
		if !ok || name == "" {
			errs = append(errs, errors.New(fmt.Sprintf("operative name not found or is empty")))
			continue
		}
		op, err := buildOperative(a, name, m)
		if err != nil {
			errs = append(errs, err)
			continue
		}
		operatives = append(operatives, op)
	}
	return
}

func buildOperative(officer messaging.Agent, name string, cfg map[string]string) (any, error) {
	switch std.Kind(name) {
	case namespace.HandlerKind:
		// Since this is only code and no state, the same operative can be used in all networks
		operative := exchange.ExchangeHandler(name)
		if operative == nil {
			return nil, errors.New(fmt.Sprintf("exchange handler is nil for name: %v", name))
		}
		return operative, nil
	case namespace.AgentKind:
		var agent messaging.Agent
		var global bool

		// Determine if a global assignment is requested
		if cfg[AssignmentKey] != AssignmentLocal {
			global = true
			agent = exchange.Agent(name)
		} else {
			// Construct a new agent as each agent has state, and a new instance is required for each network
			agent = exchange.NewAgent(name)
		}
		if agent == nil {
			return nil, errors.New(fmt.Sprintf("agent is nil for name: %v", name))
		}

		// Add agent to case officer exchange if not global
		if !global {
			m := messaging.NewConfigMessage(agent).AddTo(officer.Name())
			officer.Message(m)

			// TODO: wait for reply?
			agent.Message(messaging.NewConfigMessage(cfg))
			agent.Message(messaging.NewConfigMessage(officer))
		}
		return agent, nil
	default:
	}
	return nil, errors.New(fmt.Sprintf("invalid namespace kind: %v", std.Kind(name)))
}
