package main

import (
	"fmt"
	"os"
	"radi/SdkInit"
	"radi/ccmgmt"
	"radi/envSet"
)

const (
	configpath = "../sdkconfig.yaml"
)

func main() {

	err := envSet.EnvSet()
	/*Set system environment*/
	if err != nil {
		return
	}

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
	err = SdkInit.Sdkinit(SdkInit.Info)
	if err != nil {
		return
	}

	/*create channel by sdk instance*/
	txid := SdkInit.ChannelCreate(SdkInit.Info)
	fmt.Println("CHANNEL "+SdkInit.Info.ChannelId+" 's creating txid is : ", txid)

	/*join channel by org peers */
	err = SdkInit.JoinChannel(SdkInit.Info)
	if err != nil {
		return
	}

	/*package & install & instantiate the cc*/
	txid = ccmgmt.InstallAndInitCC(SdkInit.Info, SdkInit.ResCli)
	fmt.Println("CC"+" radiTrace "+"INSTANTIATE's txid is : ", txid)

}
