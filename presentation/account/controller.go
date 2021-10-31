package account

import (
	"context"

	pb "github.com/CA22-game-creators/cookingbomb-proto/server/pb/api"
	validator "github.com/CA22-game-creators/cookingbomb-proto/server/validation"
	"github.com/oklog/ulid/v2"
	"google.golang.org/grpc/metadata"

	getAccountInfo "github.com/CA22-game-creators/cookingbomb-apiserver/application/usecase/account/get_account_info"
	refreshSessionToken "github.com/CA22-game-creators/cookingbomb-apiserver/application/usecase/account/refresh_session_token"
	signup "github.com/CA22-game-creators/cookingbomb-apiserver/application/usecase/account/signup"
	"github.com/CA22-game-creators/cookingbomb-apiserver/errors"
)

type controller struct {
	signup              signup.InputPort
	refreshSessionToken refreshSessionToken.InputPort
	getAccountInfo      getAccountInfo.InputPort
}

func New(su signup.InputPort, gs refreshSessionToken.InputPort, ga getAccountInfo.InputPort) pb.AccountServicesServer {
	return controller{
		signup:              su,
		refreshSessionToken: gs,
		getAccountInfo:      ga,
	}
}

func (c controller) Signup(ctx context.Context, request *pb.SignupRequest) (*pb.SignupResponse, error) {
	if err := validator.Validate(request); err != nil {
		return nil, errors.InvalidArgument(err.Error())
	}

	input := signup.InputData{Name: request.Name}
	output := c.signup.Handle(input)
	if output.Err != nil {
		return nil, output.Err
	}

	return &pb.SignupResponse{
		Account: &pb.Account{
			Id:   ulid.ULID(output.Account.ID).String(),
			Name: string(output.Account.Name),
		},
		AuthToken: output.AuthToken.String(),
	}, nil
}

func (c controller) GetSessionToken(ctx context.Context, request *pb.GetSessionTokenRequest) (*pb.GetSessionTokenResponse, error) {
	if err := validator.Validate(request); err != nil {
		return nil, errors.InvalidArgument(err.Error())
	}

	input := refreshSessionToken.InputData{UserID: request.UserId, AuthToken: request.AuthToken}
	output := c.refreshSessionToken.Handle(input)
	if output.Err != nil {
		return nil, output.Err
	}

	return &pb.GetSessionTokenResponse{
		SessionToken: output.SessionToken.String(),
	}, nil
}

func (c controller) GetAccountInfo(ctx context.Context, _ *pb.GetAccountInfoRequest) (*pb.GetAccountInfoResponse, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok || len(md.Get("session-token")) == 0 {
		return nil, errors.Unauthenticated("session-token is required in metadata")
	}

	input := getAccountInfo.InputData{SessionToken: md.Get("session-token")[0]}
	output := c.getAccountInfo.Handle(input)
	if output.Err != nil {
		return nil, output.Err
	}

	return &pb.GetAccountInfoResponse{
		Account: &pb.Account{
			Id:   ulid.ULID(output.Account.ID).String(),
			Name: string(output.Account.Name),
		},
	}, nil
}
