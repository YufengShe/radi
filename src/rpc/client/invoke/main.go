package main

import (
	"context"
	client2 "github.com/cloudwego/kitex/client"
	"log"
	ccapi "radi/rpc/kitex_gen/ccAPI"
	ccoperate "radi/rpc/kitex_gen/ccAPI/ccoperation"
)

func main() {
	client, err := ccoperate.NewClient("ccOperate", client2.WithHostPorts("0.0.0.0:8888"))
	if err != nil {
		log.Fatalln("rpc client setup error" + err.Error())
	}

	args := []string{
		"1",
		"syf-data",
		"syf-bit-cs",
		"syf",
		"hash123",
		"data-addr",
	}
	req := &ccapi.CCInvokeReq{
		ChaincodeId: "testCC",
		FuncId:      "MetaRegister",
		Args_:       args,
	}
	resp, err := client.CCInvoke(context.Background(), req)
	if err != nil {
		log.Println("rpc invoke error: " + err.Error())
		return
	} else {
		log.Println("rpc invoke success txid: " + resp.Txid + " payload: " + resp.Payload)
	}
}
