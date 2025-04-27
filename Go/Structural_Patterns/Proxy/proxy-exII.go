package proxy

import "fmt"

type Service interface {
	Request(user string)
}

type RealService struct{}

func (r *RealService) Request(user string) {
	fmt.Printf("ProxyService: Receive %s request\n", user)
	fmt.Printf("RealService: Handle %s request\n", user)
}

type ProxyService struct {
	realservice *RealService
	allowUsers  map[string]bool
}

func (p *ProxyService) Request(user string) {
	if !p.allowUsers[user] {
		fmt.Printf("ProxyService: Reject %s request\n", user)
		return
	}
	p.realservice.Request(user)
}

func NewProxyService() *ProxyService {
	return &ProxyService{
		realservice: &RealService{},
		allowUsers: map[string]bool{
			"admin": true,
			"user1": true,
		},
	}
}
