package main

import (
	"fmt"
	"log"
	"os"
	"radi/SdkInit"
	"radi/rpc"
	ccapi "radi/rpc/kitex_gen/ccAPI/ccoperation"
)

const (
	configpath = "../sdkconfig.yaml"
)

func main() {

	/* Init the SdkInfo struct */
	SdkInit.Info = SdkInit.SdkInfo{
		ConfigPath:        configpath,
		ChannelId:         "radichannel",
		OrgName:           "Radi",
		OrgAdmin:          "Admin",
		OrdererOrgName:    "OrdererOrg",
		OrdererOrgAdmin:   "Admin",
		ChannelConfigPath: os.Getenv("ProjectDir") + "/channel-artifitial/radichannel.tx",
		OrdererEndPoint:   "orderer.radi.trace.com",
	}

	/*Define the resClient struct to manage resoruces*/
	/*initialize the sdk instance*/
	err := SdkInit.Sdkinit(SdkInit.Info)
	if err != nil {
		log.Fatalln("SdkInit error" + err.Error())
	}

	/*create channel by sdk instance*/
	txid := SdkInit.ChannelCreate(SdkInit.Info)
	fmt.Println("CHANNEL "+SdkInit.Info.ChannelId+" 's creating txid is : ", txid)

	/*join channel by org peers */
	err = SdkInit.JoinChannel(SdkInit.Info)
	if err != nil {
		log.Fatalln("Join Channel error" + err.Error())
	}

	/*create channel client*/
	err = SdkInit.ChannelClientCreate(SdkInit.Info)
	if err != nil {
		log.Fatalln("ChannelClient create error: " + err.Error())
	}
	///*package & install & instantiate the cc*/
	//txid = ccmgmt.InstallAndInitCC(SdkInit.Info, SdkInit.ResCli)
	//log.Println("CC"+" radiTrace "+"INSTANTIATE's txid is : ", txid)

	svr := ccapi.NewServer(new(rpc.CCOperationImpl))
	err = svr.Run()
	if err != nil {
		log.Fatalln("rpc server setup error" + err.Error())
	}

}
