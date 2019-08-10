package endpoints

import (
	"context"
	"fmt"
	"mysite/pkg/errors"
	"mysite/pkg/user/service"

	"github.com/go-kit/kit/endpoint"
)

func MakeSinUpEndPoint(svc service.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var err error
		req := request.(SignUpRequest)
		fmt.Println(req)
		err = req.Validate()
		if err != nil {
			return nil, err
		}

		err = svc.SignUp(ctx, req.Email, req.Pswd)
		res := SignUpResponse{&response{}}
		if err != nil {
			e, ok := err.(errors.Error)
			if ok {
				res.code = e.Code()
				res.error = e
			} else {
				res.code = DishandlableErrorCode
				res.error = err
			}
		}
		return res, nil
	}
}

type SignUpRequest struct {
	Email string
	Pswd  string
}

func (req *SignUpRequest) Validate() error {
	if req.Email == "" {
		return errors.AddDetail(FieldNotAllowEmptyError, "email", "can not empty")
	}
	if req.Pswd == "" {
		return errors.AddDetail(FieldNotAllowEmptyError, "pswd", "can not empty")
	}
	return nil
}

type SignUpResponse struct {
	*response
}
