package user_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	domain "github.com/CA22-game-creators/cookingbomb-apiserver/domain/model/user"
	dbModel "github.com/CA22-game-creators/cookingbomb-apiserver/infrastructure/mysql/model/user"

	tdDomain "github.com/CA22-game-creators/cookingbomb-apiserver/test/testdata/domain/user"
	tdCommonString "github.com/CA22-game-creators/cookingbomb-apiserver/test/testdata/string/common"
	tdUserString "github.com/CA22-game-creators/cookingbomb-apiserver/test/testdata/string/user"
)

func TestNew(t *testing.T) {
	t.Parallel()

	tests := []struct {
		title    string
		input    domain.User
		expected dbModel.User
	}{
		{
			title: "【正常系】ユーザーエンティティからDBモデルが作成できる",
			input: tdDomain.Entity,
			expected: dbModel.User{
				ID:              tdCommonString.ULID.Valid,
				Name:            tdUserString.Name.Valid,
				HashedAuthToken: tdCommonString.UUID.Encrypted,
			},
		},
	}

	for _, td := range tests {
		td := td

		t.Run("New:"+td.title, func(t *testing.T) {
			t.Parallel()

			output := dbModel.New(td.input)

			assert.Equal(t, td.expected, output)
		})
	}
}
