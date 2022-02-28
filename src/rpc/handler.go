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
