package services

import "github.com/MovieStoreGuy/lucid/types"

func init() {
	GetFactoryInstance().register("dns", NewDNSService)
}

type DNSService struct {
}

func NewDNSService() (types.Service, error) {
	return &DNSService{}, nil
}

func (d *DNSService) Output() ([]byte, error) {
	return []byte{}, nil
}

func (d *DNSService) Run() error {
	return nil
}
