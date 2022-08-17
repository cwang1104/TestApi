package rpcLogin

import (
	"TestApi/common/db"
	pbLogin "TestApi/proto/login"
	"context"
)

func (rpc *rpcLogin) AccountRegister(_ context.Context, req *pbLogin.AccountRegisterReq) (*pbLogin.AccountRegisterResp, error) {
	if exist := db.IsExistUserByPhone(req.Phone); !exist {
		if err := db.UserRegister(req); err != nil {
			return &pbLogin.AccountRegisterResp{
				ErrCode: 2002,
				ErrMsg:  err.Error(),
			}, nil
		}
	}

	user, _ := db.FindUserByPhone(req)

	return &pbLogin.AccountRegisterResp{
		ErrCode: 0,
		ErrMsg:  "",
		Data: &pbLogin.UserInfo{
			UID:       user.UID,
			NickName:  user.NickName,
			HeadUrl:   user.HeadURL,
			Gender:    user.Gender,
			Posizton:  user.Position,
			Birthday:  user.Birthday,
			Email:     user.Email,
			Introduce: user.Introduce,
			Love:      user.Love,

			//todo: 加入score
			//Score10:   float32(score),
		},
	}, nil
}
