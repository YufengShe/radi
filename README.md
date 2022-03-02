## Note: Kitex rpc框架不兼容Win环境，在Win下无法正常使用

## Intro
radi将一个一组织两节点的Hyperledger Fabric网络部署在服务器82.156.74.62上
同时在该服务器上启动了一个rpc server，可以接收并处理rpc client发送的请求，包括
链码安装、链码函数调用（invoke）和链码函数查询（query）。

## API
### 链码安装
```
func CCInstall(req CCInstallReq) (CCInstallResp, error)

type CCInstallReq struct{
    Name string
    Path string
}

type CCInstallResp struct{
    Txid string
}
```
- 安装链码的传入参数为CCInstallReq，包含两个字段，Name是你为待安装的链码起的名字，Path为链码存放路径
- 注意：这里Path是一个相对路径，例如我将链码文件放在raditrace文件夹下，首先需要将radiTrace文件夹放置在$GOPATH/src/chaincode路径下
这里的Path字段需要赋值为"chaincode/raditrace
- 传出参数CCInstallResp有一个字段Txid，为该链码初始化的交易id

### 链码调用（invoke）
```
func CCInvoke(req CCInvokeReq) (CCInvokeResp, error)

type CCInvokeReq struct{
    ChaincodeId string//chaincode name to invoke
    FuncId string     //funcId to invoke
    Args []string //arguments to invoke chaincode func
}

type CCInvokeResp struct{
    Txid string //txid of chaincode invoke
    Payload string //response payload of chaincode invoke
}
```
- 链码调用传入参数CCInvokeReq有三个参数，ChaincodeId为链码安装时所起的链码名； FuncId为调用的函数名称；
Args是字符串数组，里面放置调用链码所需的参数；
- 传出参数CCInvokeResp包含两个字段，Txid为调用链码的交易Id，payload为链码返回的响应结果数据；

### 链码查询 （query）
- query和invoke传入和传出参数名称不同，但参数字段是一致的，在此不做更多介绍，请参考invoke

## 使用rpc client进行链码安装、调用
将radi/src/rpc目录放置在自己的项目文件夹下
而后参考radi/src/rpc/client中的代码，创建相应的rpc客户端，即可调用rpc创建、调用自己的链码

##环境清理和重启
如果安装链码有问题，需要重新安装，则需要清理环境，重新启动fabric网络
进入82.156.74.62
进入~/project/radi
./scripts/teardown.sh //关掉网络、清理环境
./scripts/radi.sh     //重新启动网络、创建通道
./scripts/radiInit.sh //启动rpc server

而后继续通过rpc client进行相应的操作即可
