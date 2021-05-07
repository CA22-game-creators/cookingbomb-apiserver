// +build wireinject

package account

import (
	pb "github.com/CA22-game-creators/cookingbomb-proto/server/pb"
	"github.com/google/wire"

	signupApplication "github.com/CA22-game-creators/cookingbomb-apiserver/application/account/signup"
	domain "github.com/CA22-game-creators/cookingbomb-apiserver/domain/model/user"
	"github.com/CA22-game-creators/cookingbomb-apiserver/infrastructure/mysql"
	repoImpl "github.com/CA22-game-creators/cookingbomb-apiserver/infrastructure/repository/user"
	controller "github.com/CA22-game-creators/cookingbomb-apiserver/presentation/account"
	"github.com/CA22-game-creators/cookingbomb-apiserver/util"
)

func DI() pb.AccountServicesServer {
	wire.Build(
		controller.New,
		signupApplication.New,
		domain.NewFactory,
		repoImpl.NewRepository,
		util.NewCryptManager,
		util.NewIDGenerator,
		util.NewTokenGenerator,
		mysql.NewDB,
	)

	return nil
}
