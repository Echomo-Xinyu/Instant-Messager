package main

import (
	"context"
	rpc "github.com/TikTokTechImmersion/assignment_demo_2023/http-server/kitex_gen/rpc"
)

// IMServiceImpl implements the last service interface defined in the IDL.
type IMServiceImpl struct{}

// Send implements the IMServiceImpl interface.
func (s *IMServiceImpl) Send(ctx context.Context, req *rpc.SendRequest) (resp *rpc.SendResponse, err error) {
	// TODO: Your code here...
	return
}

// Pull implements the IMServiceImpl interface.
func (s *IMServiceImpl) Pull(ctx context.Context, req *rpc.PullRequest) (resp *rpc.PullResponse, err error) {
	// TODO: Your code here...
	return
}
