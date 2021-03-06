package main

import (
	"context"
	"ec2-grpc-ip-test/server/IPTestService"
	"log"
	"strings"

	"google.golang.org/grpc/peer"
)

type IPTestServiceHandler struct {
	IPTestService.UnimplementedIP_TestServer
}

func (server IPTestServiceHandler) GetIP(ctx context.Context, req *IPTestService.Req) (res *IPTestService.Res, err error) {
	reqPeer, _ := peer.FromContext(ctx)
	addrString := reqPeer.Addr.String()
	addrSplit := strings.Split(addrString, ":")

	log.Println("Req Type:", reqPeer.Addr.Network())
	log.Println("Req Address:", reqPeer.Addr.String())

	res = &IPTestService.Res{}

	res.Address = addrString
	res.Ip = addrSplit[0]
	res.Port = addrSplit[1]

	return
}
