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
	// TODO: Your code here...
	var (
		txid string
	)
	resp = new(ccapi.CCInstallResp)
	txid, err = ccmgmt.InstallCC(req.GetName(), req.GetPath())
	resp.Txid = txid
	resp.Err = err.Error()
	return
}
