package main

import (
	"context"
	ccapi "radi/rpc/kitex_gen/ccAPI"
)

// CCOperationImpl implements the last service interface defined in the IDL.
type CCOperationImpl struct{}

// CCInstall implements the CCOperationImpl interface.
func (s *CCOperationImpl) CCInstall(ctx context.Context, req *ccapi.CCInstallReq) (resp *ccapi.CCInstallResp, err error) {
	// TODO: Your code here...
	return
}

// CCInvoke implements the CCOperationImpl interface.
func (s *CCOperationImpl) CCInvoke(ctx context.Context, req *ccapi.CCInvokeReq) (resp *ccapi.CCInvokeResp, err error) {
	// TODO: Your code here...
	return
}

// CCQuery implements the CCOperationImpl interface.
func (s *CCOperationImpl) CCQuery(ctx context.Context, resp *ccapi.CCQueryResp) (resp *ccapi.CCQueryResp, err error) {
	// TODO: Your code here...
	return
}
