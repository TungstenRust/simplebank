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
	if err := val.ValidateEmailId(req.GetEmailId()); err != nil {
		violations = append(violations, fieldValidation("email_id", err))
	}
	if err := val.ValidateSecretCode(req.GetSecretCode()); err != nil {
		violations = append(violations, fieldValidation("secret_code", err))
	}
	return violations
}
