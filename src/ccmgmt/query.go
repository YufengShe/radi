package ccmgmt

import (
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"log"
	"radi/SdkInit"
)

func QueryCC(ccId string, funId string, args []string) (channel.Response, error) {
	//构造传入参数
	var ccArgs [][]byte
	for _, arg := range args {
		ccArgs = append(ccArgs, []byte(arg))
	}

	channelReq := channel.Request{
		ChaincodeID: ccId,
		Fcn:         funId,
		Args:        ccArgs,
	}

	//Query
	response, err := SdkInit.ResCli.ClientChannel.Query(channelReq)
	if err != nil {
		log.Println("Err in query CC: ", err)
		return channel.Response{}, err
	} else {
		log.Println("Success to query func " + funId + " chaincode " + ccId)
		return response, nil
	}
}
