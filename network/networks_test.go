package network

import (
	"fmt"
	"os"
)

const (
	networkFileName  = "network-config-primary.json"
	endpointFileName = "endpoint-config.json"
	subDir           = "/networktest/resource/"
)

func readFile(fileName string) ([]byte, error) {
	dir, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	return os.ReadFile(dir + subDir + fileName)
}

func ExampleBuildNetworkConfig() {
	cfg, err := BuildConfig(networkFileName, readFile)
	fmt.Printf("test: buildNetworkConfig() -> [%v] [err:%v]\n", cfg, err)

	//Output:
	//test: buildNetworkConfig() -> [map[common:core:role/authorization/http:map[name:common:resiliency:handler/authorization/http] common:core:role/cache/request/http:map[cache-control:no-store, no-cache, max-age=0 fri:22-23 host:localhost:8081 mon:8-16 name:common:resiliency:agent/cache/request/http sat:3-8 sun:13-15 thu:0-23 timeout-duration:750ms tue:6-10 wed:12-12] common:core:role/logging/request:map[name:common:resiliency:agent/logging/request] common:core:role/rate-limiting/request/http:map[assignment:local name:common:resiliency:agent/rate-limiting/request/http off-peak-duration:5m peak-duration:750ms rate-limit:1234 review-duration:567] common:core:role/routing/request/http:map[app-host:localhost:8080 assignment:global cache-host:localhost:8081 name:common:resiliency:agent/routing/request/http review-duration:4m timeout-duration:2m]]] [err:<nil>]

}

/*
func ExampleValidateOfficer() {
	agent, err := validateOfficerType("")
	fmt.Printf("test: validateOfficerType() -> [agent:%v] [err:%v]\n", agent, err)

	agent, err = validateOfficerType("test-case-officer")
	fmt.Printf("test: validateOfficerType() -> [agent:%v] [err:%v]\n", agent, err)

	name := "test-agent"
	a := messagingtest.NewAgent(name)
	exchange.Register(a)
	agent, err = validateOfficerType(name)
	fmt.Printf("test: validateOfficerType() -> [agent:%v] [err:%v]\n", agent, err)

	agent, err = validateOfficerType(caseofficer.NamespaceNamePrimary)
	fmt.Printf("test: validateOfficerType() -> [agent:%v] [err:%v]\n", agent != nil, err)

	//Output:
	//test: validateOfficerType() -> [agent:<nil>] [err:case officer name is empty]
	//test: validateOfficerType() -> [agent:<nil>] [err:agent lookup is nil for case officer: test-case-officer]
	//test: validateOfficerType() -> [agent:<nil>] [err:agent is not of type caseofficer.Agent for case officer: test-agent]
	//test: validateOfficerType() -> [agent:true] [err:<nil>]

}


*/

func ExampleReadEndpointConfig() {
	cfg, err := ReadEndpointConfig(func() ([]byte, error) {
		return readFile(endpointFileName)
	})

	fmt.Printf("test: ReadEndpointConfig() -> %v [err:%v]\n", cfg, err)

	//Output:
	//test: ReadEndpointConfig() -> [map[endpoint:primary network:network-config-primary.json pattern:/primary test:true] map[endpoint:secondary network:network-config-secondary.json pattern:/secondary test:false]] [err:<nil>]
	
}
