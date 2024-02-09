package grpc

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"hexarch/internal/adapters/framework/left/grpc/pb"
)

func (grpca Adapter) GetAddition(ctx context.Context, req *pb.OperationParameters) (*pb.Answer, error) {
	if req.GetA() == 0 || req.GetB() == 0 {
		return &pb.Answer{}, status.Error(codes.InvalidArgument, "missing required")
	}

	answer, err := grpca.api.GetAddition(req.GetA(), req.GetB())
	if err != nil {
		return &pb.Answer{}, status.Error(codes.Internal, "unexpected error")
	}

	return &pb.Answer{Value: answer}, nil
}

func (grpca Adapter) GetSubtraction(ctx context.Context, req *pb.OperationParameters) (*pb.Answer, error) {
	if req.GetA() == 0 || req.GetB() == 0 {
		return &pb.Answer{}, status.Error(codes.InvalidArgument, "missing required")
	}

	answer, err := grpca.api.GetSubtraction(req.GetA(), req.GetB())
	if err != nil {
		return &pb.Answer{}, status.Error(codes.Internal, "unexpected error")
	}

	return &pb.Answer{Value: answer}, nil
}

func (grpca Adapter) GetMultiplication(ctx context.Context, req *pb.OperationParameters) (*pb.Answer, error) {
	if req.GetA() == 0 || req.GetB() == 0 {
		return &pb.Answer{}, status.Error(codes.InvalidArgument, "missing required")
	}

	answer, err := grpca.api.GetMultiplication(req.GetA(), req.GetB())
	if err != nil {
		return &pb.Answer{}, status.Error(codes.Internal, "unexpected error")
	}

	return &pb.Answer{Value: answer}, nil
}

func (grpca Adapter) GetDivision(ctx context.Context, req *pb.OperationParameters) (*pb.Answer, error) {
	if req.GetA() == 0 || req.GetB() == 0 {
		return &pb.Answer{}, status.Error(codes.InvalidArgument, "missing required")
	}

	answer, err := grpca.api.GetDivision(req.GetA(), req.GetB())
	if err != nil {
		return &pb.Answer{}, status.Error(codes.Internal, "unexpected error")
	}

	return &pb.Answer{Value: answer}, nil
}
