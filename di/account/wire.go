// +build wireinject

package account

import (
	pb "github.com/CA22-game-creators/cookingbomb-proto/server/pb/api"
	"github.com/google/wire"

	auth "github.com/CA22-game-creators/cookingbomb-apiserver/application/common/auth/authenticate"
	getAccountInfo "github.com/CA22-game-creators/cookingbomb-apiserver/application/usecase/account/get_account_info"
	refreshSessionToken "github.com/CA22-game-creators/cookingbomb-apiserver/application/usecase/account/refresh_session_token"
	signup "github.com/CA22-game-creators/cookingbomb-apiserver/application/usecase/account/signup"
	domain "github.com/CA22-game-creators/cookingbomb-apiserver/domain/model/user"
	accountIDFetcher "github.com/CA22-game-creators/cookingbomb-apiserver/infrastructure/application/common/auth/authenticate/account_id_fetcher"
	sessionTokenRefresher "github.com/CA22-game-creators/cookingbomb-apiserver/infrastructure/application/usecase/account/refresh_session_token/session_token_refresher"
	"github.com/CA22-game-creators/cookingbomb-apiserver/infrastructure/mysql"
	repoImpl "github.com/CA22-game-creators/cookingbomb-apiserver/infrastructure/repository/user"
	controller "github.com/CA22-game-creators/cookingbomb-apiserver/presentation/account"
	"github.com/CA22-game-creators/cookingbomb-apiserver/util"
)

func DI() pb.AccountServicesServer {
	wire.Build(
		controller.New,
		signup.New,
		refreshSessionToken.New,
		getAccountInfo.New,
		domain.NewFactory,
		repoImpl.NewRepository,
		accountIDFetcher.New,
		auth.New,
		sessionTokenRefresher.New,
		util.NewCryptManager,
		util.NewIDGenerator,
		util.NewTokenGenerator,
		mysql.NewDB,
	)

	return nil
}
