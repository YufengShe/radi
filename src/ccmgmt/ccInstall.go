package ccmgmt

import (
	"fmt"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/resmgmt"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/errors/retry"
	packager "github.com/hyperledger/fabric-sdk-go/pkg/fab/ccpackager/gopackager"
	"github.com/hyperledger/fabric-sdk-go/third_party/github.com/hyperledger/fabric/common/cauthdsl"
	"log"
	"os"
	"radi/SdkInit"
)

//the process of chaincode install: firstly package, then install and finally instantiation
func InstallAndInitCC(info SdkInit.SdkInfo, client *SdkInit.ResClient) string {

	//package cc
	pkg, err := packager.NewCCPackage("chaincode/raditrace", os.Getenv("GOPATH"))
	if err != nil {
		fmt.Println("error in package chaincode : ", err)
		return ""
	} else {
		fmt.Println("successfully package the chaincode raditrace")
	}

	//install cc
	installccReq := resmgmt.InstallCCRequest{ //create install cc Request
		Name:    "radiTraceCC",
		Path:    "chaincode/raditrace",
		Version: "1.0",
		Package: pkg,
	}
	responses, err := client.ResMgmtClient.InstallCC(installccReq, resmgmt.WithRetry(retry.DefaultResMgmtOpts))
	if err != nil {
		fmt.Println("err in installCC process : ", err)
		return ""
	} else {
		fmt.Println("successfully install cc : ", responses)
	}

	//instantiate cc
	ccPolicy := cauthdsl.SignedByAnyMember([]string{"RadiMSP"})
	instantiateReq := resmgmt.InstantiateCCRequest{
		Name:    "radiTraceCC",
		Path:    "chaincode/raditrace",
		Version: "1.0",
		Args:    [][]byte{[]byte("init")}, //Args shoud define the arguments for chaincode initialize function
		Policy:  ccPolicy,                 //chaincode policy described when instantiate the chaincode
	}

	instantiateResponse, err := client.ResMgmtClient.InstantiateCC(info.ChannelId, instantiateReq, resmgmt.WithRetry(retry.DefaultResMgmtOpts))
	if err != nil {
		fmt.Println("err in instantiate the cc : ", err)
		return ""
	} else {
		fmt.Println("successfully instantiate the cc")
		return string(instantiateResponse.TransactionID)
	}
}

func InstallCC(name, path string) (string, error) {
	//package cc
	pkg, err := packager.NewCCPackage(path, os.Getenv("GOPATH"))
	if err != nil {
		log.Println("error in package chaincode : ", err)
		return "", err
	} else {
		log.Println("successfully package the chaincode raditrace")
	}

	//install cc
	installccReq := resmgmt.InstallCCRequest{ //create install cc Request
		Name:    name,
		Path:    path,
		Version: "1.0",
		Package: pkg,
	}
	responses, err := SdkInit.ResCli.ResMgmtClient.InstallCC(installccReq, resmgmt.WithRetry(retry.DefaultResMgmtOpts))
	if err != nil {
		log.Println("err in installCC process : ", err)
		return "", err
	} else {
		log.Println("successfully install cc : ", responses)
	}
	//instantiate cc
	ccPolicy := cauthdsl.SignedByAnyMember([]string{"RadiMSP"})
	instantiateReq := resmgmt.InstantiateCCRequest{
		Name:    name,
		Path:    path,
		Version: "1.0",
		Args:    [][]byte{[]byte("init")}, //Args shoud define the arguments for chaincode initialize function
		Policy:  ccPolicy,                 //chaincode policy described when instantiate the chaincode
	}

	instantiateResponse, err := SdkInit.ResCli.ResMgmtClient.InstantiateCC(SdkInit.Info.ChannelId, instantiateReq, resmgmt.WithRetry(retry.DefaultResMgmtOpts))
	if err != nil {
		log.Println("err in instantiate the cc : ", err)
		return "", err
	} else {
		fmt.Println("successfully instantiate the cc")
		return string(instantiateResponse.TransactionID), nil
	}
}
