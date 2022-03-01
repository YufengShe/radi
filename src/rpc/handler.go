package rpc

import (
	"context"
	"radi/ccmgmt"
	ccapi "radi/rpc/kitex_gen/ccAPI"
)

// CCOperationImpl implements the last service interface defined in the IDL.
type CCOperationImpl struct{}

// CCInstall implements the CCOperationImpl interface.
func (s *CCOperationImpl) CCInstall(ctx context.Context, req *ccapi.CCInstallReq) (resp *ccapi.CCInstallResp, err error) {
	resp = new(ccapi.CCInstallResp)
	txid, err := ccmgmt.InstallCC(req.GetName(), req.GetPath())
	resp.Txid = txid
	return
}

// CCInvoke implements the CCOperationImpl interface.
func (s *CCOperationImpl) CCInvoke(ctx context.Context, req *ccapi.CCInvokeReq) (resp *ccapi.CCInvokeResp, err error) {
	cResp, err := ccmgmt.CCInvoke(req.GetChaincodeId(), req.GetFuncId(), req.GetArgs_())
	if err != nil {
		return &ccapi.CCInvokeResp{}, err
	} else {
		return &ccapi.CCInvokeResp{
			Txid:    string(cResp.TransactionID),
			Payload: string(cResp.Payload),
		}, nil
	}
}

// CCQuery implements the CCOperationImpl interface.
func (s *CCOperationImpl) CCQuery(ctx context.Context, req *ccapi.CCQueryReq) (resp *ccapi.CCQueryResp, err error) {
	cResp, err := ccmgmt.CCQuery(req.GetChaincodeId(), req.GetFuncId(), req.GetArgs_())
	if err != nil {
		return &ccapi.CCQueryResp{}, err
	} else {
		return &ccapi.CCQueryResp{
			Txid:    string(cResp.TransactionID),
			Payload: string(cResp.Payload),
		}, nil
	}
}
