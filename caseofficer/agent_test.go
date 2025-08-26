package caseofficer

import (
	"fmt"
)

func ExampleNewAgent() {
	a := newAgent("core:common:agent/caseofficer/request/http/test")

	fmt.Printf("test: NewAgent() -> [%v]\n", a.Name())

	//Output:
	//test: NewAgent() -> [core:common:agent/caseofficer/request/http/test]

}
