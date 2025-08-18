package logger

import (
	"github.com/appellative-ai/agency/logx"
	"github.com/appellative-ai/core/messaging"
	"github.com/appellative-ai/core/std"
)

func (a *agentT) configure(m *messaging.Message) {
	if m == nil || m.Name != messaging.ConfigEvent {
		return
	}
	if ops, ok := messaging.ConfigContent[[]logx.Operator](m); ok {
		if len(ops) > 0 {
			var err error
			a.operators, err = logx.InitOperators(ops)
			if err != nil {
				messaging.Reply(m, std.NewStatus(std.StatusInvalidArgument, "", err), a.Name())
			}
		}
	}
}
