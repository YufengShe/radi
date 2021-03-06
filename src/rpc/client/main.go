package main

import (
	"context"
	client2 "github.com/cloudwego/kitex/client"
	"log"
	ccapi "radi/rpc/kitex_gen/ccAPI"
	ccoperate "radi/rpc/kitex_gen/ccAPI/ccoperation"
)

func main() {
	client, err := ccoperate.NewClient("ccOperate", client2.WithHostPorts("82.156.74.62:8888"))
	if err != nil {
		log.Fatalln("rpc client setup error" + err.Error())
	}

	req := &ccapi.CCInstallReq{
		Name: "testCC",
		Path: "chaincode/testcc",
	}
	resp, err := client.CCInstall(context.Background(), req)
	if err != nil {
		log.Println("rpc invoke error: " + err.Error())
		return
	}
	log.Println("rpc invoke success: " + resp.Txid)
}
