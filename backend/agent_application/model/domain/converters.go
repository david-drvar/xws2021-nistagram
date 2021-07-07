package domain

import protopb "github.com/david-drvar/xws2021-nistagram/common/proto"

func (l LoginRequest) ConvertFromGrpc(request *protopb.LoginRequestAgentApp) LoginRequest {
	return LoginRequest{
		Email:    request.Email,
		Password: request.Password,
	}
}
