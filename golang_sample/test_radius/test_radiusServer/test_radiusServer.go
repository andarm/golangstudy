package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bronze1man/radius"
)

type radiusService struct{}

func (p radiusService) RadiusHandle(request *radius.Packet) *radius.Packet {
	// a pretty print of the request.
	fmt.Printf("[Authenticate] %s\n", request.String())
	npac := request.Reply()
	switch request.Code {
	case radius.AccessRequest:
		fmt.Println("get int here")
		// check username and password
		if request.GetUsername() == "bob" && request.GetPassword() == "hello" {
			npac.Code = radius.AccessAccept
			// add Vendor-specific attribute - Vendor Cisco (code 9) Attribute h323-remote-address (code 23)
			npac.AddVSA(radius.VSA{Vendor: 9, Type: 23, Value: []byte("10.20.30.40")})
		} else {
			npac.Code = radius.AccessReject
			npac.AddAVP(radius.AVP{Type: radius.ReplyMessage, Value: []byte("you dick!")})
		}
	case radius.AccountingRequest:
		// accounting start or end
		npac.Code = radius.AccountingResponse
	default:
		npac.Code = radius.AccessAccept
	}
	return npac
}

func main() {
	s := radius.NewServer(":1812", "testing123", radiusService{})

	// or you can convert it to a server that accept request
	// from some host with different secret
	// cls := radius.NewClientList([]radius.Client{
	// 		radius.NewClient("127.0.0.1", "secret1"),
	// 		radius.NewClient("10.10.10.10", "secret2"),
	// })
	// s.WithClientList(cls)

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	errChan := make(chan error)
	go func() {
		fmt.Println("waiting for packets...")
		err := s.ListenAndServe()
		if err != nil {
			errChan <- err
		}
	}()
	select {
	case <-signalChan:
		log.Println("stopping server...")
		s.Stop()
	case err := <-errChan:
		log.Println("[ERR] %v", err.Error())
	}
}
