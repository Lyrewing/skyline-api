package middleware

import (
	"fmt"
	"net"

	"github.com/hashicorp/consul/api"
)

type ConsulRegister struct {
	Address                        string
	Service                        string
	Tags                           []string
	Port                           int
	DeregisterCriticalServiceAfter string
	Interval                       string
}

func NewConsulRegister() *ConsulRegister {
	return &ConsulRegister{
		Address: "192.168.31.207",
		Service: "skyline-api",
		Tags:    []string{""},
		Port:    18500,
		DeregisterCriticalServiceAfter: "60s",
		Interval:                       "10s",
	}
}
func RegisterServer(register *ConsulRegister) error {
	config := api.DefaultConfig()
	var client *api.Client
	var err error
	if client, err = api.NewClient(config); err != nil {
		return err
	}
	IP := localIP()
	agent := client.Agent()
	reg := &api.AgentServiceRegistration{
		ID:      fmt.Sprintf("%v-%v-%v", register.Service, IP, register.Port),
		Name:    register.Service,
		Tags:    []string{"", ""},
		Address: IP,
		Check: &api.AgentServiceCheck{
			HTTP:                           fmt.Sprintf("http://%s:%d/health", register.Address, register.Port),
			Timeout:                        "3s",
			Interval:                       register.Interval,
			DeregisterCriticalServiceAfter: register.DeregisterCriticalServiceAfter,
		},
	}
	if err := agent.ServiceRegister(reg); err != nil {
		return err
	}
	return nil

}

func localIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, address := range addrs {
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}
