package endpoints

import (
	"context"
	"mysite/pkg/errors"
	"mysite/pkg/user/service"
	"time"

	"github.com/go-kit/kit/endpoint"
)

func MakeLoginEndPoint(svc service.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var err error
		req := request.(LoginRequest)
		err = req.Validate()
		if err != nil {
			return nil, err
		}
		resp := LoginResponse{response: &response{}}
		resp.ID, resp.Email, resp.CreatedAt, resp.UpdatedAt, err = svc.Login(ctx, req.Email, req.Pswd)
		if err != nil {
			svcErr, ok := err.(errors.Error)
			if ok {
				resp.code = svcErr.Code()
			} else {
				resp.code = DishandlableErrorCode
			}
		}
		resp.error = err
		return resp, nil
	}
}

type LoginRequest struct {
	Email string
	Pswd  string
}

func (req *LoginRequest) Validate() error {
	return nil
}

type LoginResponse struct {
	ID                   string
	Email                string
	CreatedAt, UpdatedAt time.Time
	//error
	*response
}
