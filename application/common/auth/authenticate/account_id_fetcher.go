//go:generate mockgen -source=$GOFILE -destination=../../../../mock/application/common/auth/authenticate/$GOFILE
package auth

import (
	domain "github.com/CA22-game-creators/cookingbomb-apiserver/domain/model/user"
)

type AccountIDFetcher interface {
	Handle(sessionToken string) (domain.ID, error)
}
