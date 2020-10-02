package endpoints

import (
	"context"
	"mysite/pkg/errors"
	"mysite/pkg/user/service"

	"github.com/go-kit/kit/endpoint"
)

func MakeGetProfileEndPoint(svc service.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var err error

		resp := LoginResponse{response: &response{}}
		user, err := svc.GetProfile(ctx)
		if err != nil {
			svcErr, ok := err.(errors.Error)
			if ok {
				resp.code = svcErr.Code()
			} else {
				resp.code = DishandlableErrorCode
			}
		} else {
			resp.ID, resp.Email, resp.CreatedAt, resp.UpdatedAt = user.ID, user.Email, user.CreatedAt, user.UpdatedAt
		}

		resp.error = err
		return resp, nil
	}
}
