package caseofficer

import (
	"fmt"
	"github.com/appellative-ai/collective/exchange"
	"github.com/appellative-ai/core/messaging"
	"github.com/appellative-ai/core/messaging/messagingtest"
	"github.com/appellative-ai/core/rest"
	"net/http"
	"reflect"
)

const (
	loggingRole       = "logging"
	authorizationRole = "authorization"
	cacheRole         = "cache"
	rateLimitingRole  = "rate-limiting"
	routingRole       = "routing"
	authorizationName = "Authorization"
	namespaceNameAuth = "test:resiliency:handler/authorization/http"
)

var (
	roles = []string{loggingRole, authorizationRole, cacheRole, rateLimitingRole, routingRole}
)

func authorization(next rest.Exchange) rest.Exchange {
	return func(r *http.Request) (resp *http.Response, err error) {
		auth := r.Header.Get(authorizationName)
		if auth == "" {
			return &http.Response{StatusCode: http.StatusUnauthorized}, nil
		}
		return next(r)
	}
}

func ExampleBuildOperative_Error() {
	cfg := make(map[string]string)
	agent := messagingtest.NewAgent("agent\test")

	name := "any:any:aspect/test/one"
	t, err := buildOperative(agent, name, cfg)
	fmt.Printf("test: buildOperative(\"%v\") -> [%v] [err:%v]\n", name, t, err)

	name = "any:any:link/test/one"
	t, err = buildOperative(agent, name, cfg)
	fmt.Printf("test: buildOperative(\"%v\") -> [%v] [err:%v]\n", name, t, err)

	name = "any:any:agent/test/one"
	t, err = buildOperative(agent, name, cfg)
	fmt.Printf("test: buildOperative(\"%v\") -> [%v] [err:%v]\n", name, t, err)

	//Output:
	//test: buildOperative("any:any:aspect/test/one") -> [<nil>] [err:invalid namespace kind: aspect]
	//test: buildOperative("any:any:link/test/one") -> [<nil>] [err:invalid namespace kind: link]
	//test: buildOperative("any:any:agent/test/one") -> [<nil>] [err:agent is nil for name: any:any:agent/test/one]

}

func ExampleBuildOperative() {
	name := "any:any:handler/test/one"
	cfg := make(map[string]string)
	cfg[NameKey] = "any:any:handler/test/one"

	agent := messagingtest.NewAgent("agent\test")
	exchange.RegisterExchangeHandler(name, authorization)

	cfg[NameKey] = name
	t, err := buildOperative(agent, name, cfg)
	fmt.Printf("test: buildOperative() -> [%v] [err:%v]\n", reflect.TypeOf(t), err)

	name = "any:any:agent/test/one"
	cfg[NameKey] = name
	exchange.RegisterConstructor(name, func() messaging.Agent {
		return messagingtest.NewAgent(name)
	})
	t, err = buildOperative(agent, name, cfg)
	fmt.Printf("test: buildOperative() -> [%v] [err:%v]\n", reflect.TypeOf(t), err)

	//Output:
	//test: buildOperative() -> [func(rest.Exchange) rest.Exchange] [err:<nil>]
	//test: buildOperative() -> [*messagingtest.AgentT] [err:<nil>]

}

func ExampleBuildNetwork_Error() {
	officer := messagingtest.NewAgent("*:*:agent/test")
	var netCfg []map[string]string

	chain, errs := buildNetwork(nil, nil)
	fmt.Printf("test: buildNetwork() -> [chain:%v] %v\n", chain, errs)

	chain, errs = buildNetwork(officer, netCfg)
	fmt.Printf("test: buildNetwork() -> [chain:%v] %v\n", chain, errs)

	netCfg = append(netCfg, make(map[string]string))
	chain, errs = buildNetwork(officer, netCfg)
	fmt.Printf("test: buildNetwork() -> [chain:%v] %v\n", chain, errs)

	//Output:
	//test: buildNetwork() -> [chain:[]] [agent is nil]
	//test: buildNetwork() -> [chain:[]] [network configuration is nil or empty]
	//test: buildNetwork() -> [chain:[]] [operative name not found or is empty]

}

func ExampleBuildNetwork() {
	officer := messagingtest.NewAgent("*:*:agent/test")
	var netCfg []map[string]string
	exchange.RegisterExchangeHandler(namespaceNameAuth, authorization)

	netCfg = append(netCfg, map[string]string{NameKey: namespaceNameAuth})
	chain, errs := buildNetwork(officer, netCfg)
	fmt.Printf("test: buildNetwork() -> [chain:%v] %v\n", len(chain), errs)

	//Output:
	//test: buildNetwork() -> [chain:1] []

}
