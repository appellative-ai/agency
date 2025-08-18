package caseofficer

import (
	"github.com/appellative-ai/core/messaging"
	"github.com/appellative-ai/core/std"
)

func (a *agentT) config(m *messaging.Message) {
	if m == nil || m.Name != messaging.ConfigEvent {
		return
	}
	if m.IsRecipient(a.name) {
		if a1, ok := messaging.ConfigContent[messaging.Agent](m); ok {
			err := a.agents.Register(a1)
			if err != nil {
				messaging.Reply(m, std.NewStatus(std.StatusInvalidContent, a.Name(), err), a.Name())
			}
		}
		return
	}
	a.agents.Message(m)
	return
}
