package ccmgmt

import (
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"log"
	"radi/SdkInit"
)

func CCInvoke(ccId string, funId string, args []string) (channel.Response, error) {

	////register a cc event for the execution
	//reg, eventChan, err := SdkInit.ResCli.ClientChannel.RegisterChaincodeEvent(ccId, eventFilter)
	//if err != nil {
	//	fmt.Println("Err in CC Event Registration : ", ccId, " + ", eventFilter, " ", err)
	//	return channel.Response{}, err
	//} else {
	//	fmt.Println("Successfully Register CC Event : ", ccId, " + ", eventFilter)
	//}
	//defer resclient.ClientChannel.UnregisterChaincodeEvent(reg)

	//construct in params
	var ccArgs [][]byte
	for _, arg := range args {
		ccArgs = append(ccArgs, []byte(arg))
	}

	//execute chaincode
	channelReq := channel.Request{
		ChaincodeID: ccId,
		Fcn:         funId,
		Args:        ccArgs,
	}

	//send request for execution
	response, err := SdkInit.ResCli.ClientChannel.Execute(channelReq)

	if err != nil {
		log.Println("Err in Chaincode invoke : " + err.Error())
		return channel.Response{}, err
	} else {
		log.Println("Success to invoke func " + funId + " chaincode " + ccId)
		return response, nil
	}

	////listen to the event and receive the response
	//select {
	//case ccEvent := <-eventChan:
	//	fmt.Println("The Tx listened by event ", ccEvent, " has been committed into Blockchain FileSystem and Updated the State Ledger!")
	//	return response, nil
	//case <-time.After(time.Second * 50):
	//	fmt.Println("Err in ccEvent Receive!")
	//	return channel.Response{}, errors.New("")
	//}
}
