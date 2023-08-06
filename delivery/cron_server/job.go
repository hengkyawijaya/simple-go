package cron_server

import "fmt"

// Hello is a function to say hello
func (c *cronServer) Hello() {
	fmt.Println("CRON Hello(): ", c.uc.HelloUsecase.Hello())
}

// GetIP is a function to get IP info
func (c *cronServer) GetIP() {
	ip, _ := c.repo.IPInfoRepository.GetIPInfo()
	fmt.Println("CRON GetIP(): ", ip)
}

