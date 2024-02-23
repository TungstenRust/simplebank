package gapi

import (
	"context"
	"github.com/TungstenRust/simplebank/pb"
	"github.com/TungstenRust/simplebank/val"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
)

func (server *Server) VerifyEmail(ctx context.Context, req *pb.VerifyEmailRequest) (*pb.VerifyEmailResponse, error) {
	violations := validateVerifyEmailRequest(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}

	rsp := &pb.VerifyEmailResponse{}
	return rsp, nil
}
func validateVerifyEmailRequest(req *pb.VerifyEmailRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := val.ValidateUsername(req.GetUsername()); err != nil {
		violations = append(violations, fieldValidation("username", err))
	}
	if err := val.ValidatePassword(req.GetPassword()); err != nil {
		violations = append(violations, fieldValidation("password", err))
	}
	if err := val.ValidateFullname(req.GetFullName()); err != nil {
		violations = append(violations, fieldValidation("full_name", err))
	}
	if err := val.ValidateEmail(req.GetEmail()); err != nil {
		violations = append(violations, fieldValidation("email", err))
	}
	return violations
}
