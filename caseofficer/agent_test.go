package caseofficer

import (
	"fmt"
	"github.com/appellative-ai/collective/operations/operationstest"
)

func ExampleNewAgent() {
	a := newAgent("core:common:agent/caseofficer/request/http/test", operationstest.NewNotifier())

	fmt.Printf("test: NewAgent() -> [%v]\n", a.Name())

	//Output:
	//test: NewAgent() -> [core:common:agent/caseofficer/request/http/test]

}
