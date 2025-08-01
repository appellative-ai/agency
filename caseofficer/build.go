package caseofficer

import (
	"errors"
	"fmt"
	"github.com/appellative-ai/collective/exchange"
	"github.com/appellative-ai/collective/namespace"
	"github.com/appellative-ai/core/messaging"
	"github.com/appellative-ai/core/std"
)

func buildNetwork(a messaging.Agent, netCfg map[string]map[string]string, roles []string) (chain []any, errs []error) {
	if a == nil {
		return nil, []error{errors.New("agent is nil")}
	}
	if len(netCfg) == 0 {
		return nil, []error{errors.New("network configuration is nil or empty")}
	}

	for _, role := range roles {
		agentCfg, ok := netCfg[role]
		if !ok {
			continue
		}
		link, err := buildLink(role, agentCfg, a)
		if err != nil {
			errs = append(errs, err)
			continue
		}
		chain = append(chain, link)
	}
	return
}

func buildLink(role string, cfg map[string]string, officer messaging.Agent) (any, error) {
	name, ok := cfg[NameKey]
	if !ok || name == "" {
		return nil, errors.New(fmt.Sprintf("agent or exchange name not found or is empty for role: %v", role))
	}
	switch std.Kind(name) {
	case namespace.HandlerKind:
		// Since this is only code and no state, the same link can be used in all networks
		link := exchange.ExchangeHandler(name)
		if link == nil {
			return nil, errors.New(fmt.Sprintf("exchange handler is nil for name: %v and role: %v", name, role))
		}
		return link, nil
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
			return nil, errors.New(fmt.Sprintf("agent is nil for name: %v and role: %v", name, role))
		}

		// Add agent to case officer exchange if not global
		if !global {
			m := messaging.NewAgentMessage(agent).AddTo(officer.Name())
			officer.Message(m)

			// TODO: wait for reply?
			agent.Message(messaging.NewMapMessage(cfg))
			agent.Message(messaging.NewAgentMessage(officer))
		}
		return agent, nil
	default:
	}
	return nil, errors.New(fmt.Sprintf("invalid Namespace kind: %v and role: %v", std.Kind(name), role))
}
