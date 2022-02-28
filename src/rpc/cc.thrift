namespace go ccAPI

struct CCInstallReq {
    1: string name  //name you defined for the chaincode
    2: string path  //file path of the chaincode directory
}

struct CCInstallResp {
    1: string txid //txid of install operation
    2: string err //error message
}

service CCOperation {
    CCInstallResp CCInstall(1: CCInstallReq req)
}