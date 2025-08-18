package caseofficer

import (
	"fmt"
	"github.com/appellative-ai/collective/notification/notificationtest"
	"github.com/appellative-ai/core/messaging"
)

func ExampleConfig() {
	name := "core:common:agent/caseofficer/request/http/test"
	a := newAgent(name, notificationtest.NewNotifier())
	fmt.Printf("test: newAgent() -> [count:%v]\n", a.agents.Count())

	a2 := newAgent("core:common:agent/test", notificationtest.NewNotifier())

	m := messaging.NewConfigMessage(a2).AddTo(name)
	a.Message(m)
	fmt.Printf("test: Message() -> [count:%v]\n", a.agents.Count())

	//Output:
	//test: newAgent() -> [count:0]
	//test: Message() -> [count:1]

}
