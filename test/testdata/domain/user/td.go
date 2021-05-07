package testdata

import (
	"github.com/oklog/ulid/v2"

	"github.com/CA22-game-creators/cookingbomb-apiserver/domain/model/user"

	tdCommonString "github.com/CA22-game-creators/cookingbomb-apiserver/test/testdata/string/common"
	tdUserString "github.com/CA22-game-creators/cookingbomb-apiserver/test/testdata/string/user"
)

var Entity = user.User{ID, Name, HashedAuthToken}
var ID = user.ID(ulid.MustParse(tdCommonString.ULID.Valid))
var Name = user.Name(tdUserString.Name.Valid)
var HashedAuthToken = user.HashedAuthToken([]byte(tdCommonString.UUID.Encrypted))
