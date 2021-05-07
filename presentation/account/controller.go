package account

import (
	"context"

	pb "github.com/CA22-game-creators/cookingbomb-proto/server/pb"
	"github.com/oklog/ulid/v2"

	signupApplication "github.com/CA22-game-creators/cookingbomb-apiserver/application/account/signup"
)

type controller struct {
	signupApplication signupApplication.InputPort
}

func New(s signupApplication.InputPort) pb.AccountServicesServer {
	return &controller{
		signupApplication: s,
	}
}

func (c controller) Signup(ctx context.Context, request *pb.SignupRequest) (*pb.SignupResponse, error) {
	input := signupApplication.InputData{Name: request.Name}

	output := c.signupApplication.Handle(input)
	if output.Err != nil {
		return nil, output.Err
	}

	return &pb.SignupResponse{
		Account: &pb.Account{
			Id:        ulid.ULID(output.Account.ID).String(),
			Name:      string(output.Account.Name),
			AuthToken: output.AuthToken.String(),
		},
	}, nil
}
