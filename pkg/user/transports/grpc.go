package transports

import (
	"context"
	"errors"
	"mysite/pkg/user/pb"

	"mysite/pkg/user/endpoints"

	"github.com/go-kit/kit/auth/jwt"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/go-kit/kit/transport/grpc"
	"github.com/golang/protobuf/ptypes/timestamp"
)

type GRPCServer struct {
	login  grpc.Handler
	signUp grpc.Handler
	logger log.Logger
}

func (s *GRPCServer) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	ctx, res, err := s.login.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	r, ok := res.(*pb.LoginResponse)
	if !ok {
		return nil, errors.New("wrong type")
	}
	return r, nil
}

func (s *GRPCServer) SignUp(ctx context.Context, req *pb.SignUpRequest) (*pb.SingUpRespond, error) {
	ctx, res, err := s.signUp.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	r, ok := res.(*pb.SingUpRespond)
	if !ok {
		return nil, errors.New("wrong type")
	}
	return r, nil
}

func NewGrpcServer(_ context.Context, endpoint endpoints.EndPoints, logger log.Logger) pb.UserServiceServer {
	logger = level.Error(logger)
	return &GRPCServer{
		logger: logger,
		login:  grpc.NewServer(endpoint.LoginEndPoint, grpcDecodeLoginRequest, grpcEncodeLoginResponse, grpc.ServerErrorHandler(&grpcErrorHandler{logger: logger}), grpc.ServerErrorLogger(logger), grpc.ServerBefore(jwt.GRPCToContext())),
		signUp: grpc.NewServer(endpoint.SignUpEndPoint, grpcDecodeSinUpRequest, grpcEncodeSignUpResponse, grpc.ServerErrorHandler(&grpcErrorHandler{logger: logger}), grpc.ServerErrorLogger(logger)),
	}
}

func grpcEncodeLoginResponse(_ context.Context, r interface{}) (interface{}, error) {
	errRes, ok := r.(endpoints.Response)
	if ok {
		if errRes.Code() != "" {
			return &pb.LoginResponse{
				Code:  errRes.Code(),
				Error: errRes.Error(),
			}, nil
		}
	}

	res, ok := r.(endpoints.LoginResponse)
	if !ok {
		return nil, errors.New("empty")
	}

	return &pb.LoginResponse{
		Code: "100",
		User: &pb.User{
			Id:       res.ID,
			UserName: res.Email,
			CreatedAt: &timestamp.Timestamp{
				Seconds: res.CreatedAt.Unix(),
			},
			UpdatedAt: &timestamp.Timestamp{
				Seconds: res.UpdatedAt.Unix(),
			},
		},
	}, nil
}

func grpcEncodeSignUpResponse(_ context.Context, r interface{}) (interface{}, error) {
	errRes := r.(endpoints.Response)

	if errRes.Code() != "" {
		return &pb.LoginResponse{
			Code:  errRes.Code(),
			Error: errRes.Error(),
		}, nil
	}

	return &pb.SingUpRespond{}, nil
}

func grpcDecodeLoginRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(*pb.LoginRequest)
	return endpoints.LoginRequest{
		Email: req.UserName,
		Pswd:  req.Pswd,
	}, nil
}

func grpcDecodeSinUpRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(*pb.SignUpRequest)
	return endpoints.SignUpRequest{
		Email: req.UserName,
		Pswd:  req.Pswd,
	}, nil
}

type grpcErrorHandler struct {
	logger log.Logger
}

func (h *grpcErrorHandler) Handle(ctx context.Context, err error) {
	level.Error(h.logger).Log("error", err)
}
