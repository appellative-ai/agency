package caseofficer

import (
	"fmt"
	"github.com/appellative-ai/core/messaging"
)

func ExampleConfig() {
	name := "core:common:agent/caseofficer/request/http/test"
	a := newAgent(name)
	fmt.Printf("test: newAgent() -> [count:%v]\n", a.agents.Count())

	a2 := newAgent("core:common:agent/test")

	m := messaging.NewConfigMessage(a2).AddTo(name)
	a.Message(m)
	fmt.Printf("test: Message() -> [count:%v]\n", a.agents.Count())

	//Output:
	//test: newAgent() -> [count:0]
	//test: Message() -> [count:1]

}
