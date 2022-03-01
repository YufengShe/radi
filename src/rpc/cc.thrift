namespace go ccAPI

struct CCInstallReq {
    1: string name  //name you defined for the chaincode
    2: string path  //file path of the chaincode directory
}

struct CCInstallResp {
    1: string txid //txid of install operation
}

struct CCInvokeReq {
    1: string chaincodeId //chaincode name to invoke
    2: string funcId      //funcId to invoke
    3: list<string> args  //arguments to invoke chaincode func
}

struct CCInvokeResp {
    1: string txid  //txid of chaincode invoke
    2: string payload  //response payload of chaincode invoke
}

struct CCQueryReq {
    1: string chaincodeId //chaincode name to query
    2: string funcId      //funcId to query
    3: list<string> args  //arguments to query chaincode func
}

struct CCQueryResp {
    1: string txid  //txid of chaincode query
    2: string payload  //response payload of chaincode query
}

service CCOperation {
    CCInstallResp CCInstall(1: CCInstallReq req)
    CCInvokeResp CCInvoke(1: CCInvokeReq req)
    CCQueryResp CCQuery(1: CCQueryResp resp)
}