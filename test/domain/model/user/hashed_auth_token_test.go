package user_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/CA22-game-creators/cookingbomb-apiserver/domain/model/user"
	"github.com/CA22-game-creators/cookingbomb-apiserver/errors"

	tdDomain "github.com/CA22-game-creators/cookingbomb-apiserver/test/testdata/domain/user"
	tdString "github.com/CA22-game-creators/cookingbomb-apiserver/test/testdata/string/common"
)

func TestNewHashedAuthToken(t *testing.T) {
	t.Parallel()

	tests := []struct {
		title     string
		input     []byte
		expected1 user.HashedAuthToken
		expected2 error
	}{
		{
			title:     "【正常系】ユーザーの認証トークンハッシュ値の値オブジェクトが作成できる",
			input:     []byte(tdString.UUID.Encrypted),
			expected1: tdDomain.HashedAuthToken,
			expected2: nil,
		},
		{
			title:     "【異常系】認証トークンハッシュ値がzero値",
			input:     nil,
			expected1: nil,
			expected2: errors.Internal("user hashed_auth_token is nil"),
		},
	}

	for _, td := range tests {
		td := td

		t.Run("NewHashedAuthToken:"+td.title, func(t *testing.T) {
			t.Parallel()

			output1, output2 := user.NewHashedAuthToken(td.input)

			assert.Equal(t, td.expected1, output1)
			assert.Equal(t, td.expected2, output2)
		})
	}
}
