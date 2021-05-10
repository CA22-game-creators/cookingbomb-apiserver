package user_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/CA22-game-creators/cookingbomb-apiserver/domain/model/user"
	"github.com/CA22-game-creators/cookingbomb-apiserver/errors"

	tdDomain "github.com/CA22-game-creators/cookingbomb-apiserver/test/testdata/domain/user"
)

func TestNew(t *testing.T) {
	t.Parallel()

	tests := []struct {
		title     string
		input1    user.ID
		input2    user.Name
		input3    user.HashedAuthToken
		expected1 user.User
		expected2 error
	}{
		{
			title:     "【正常系】ユーザーアカウントのエンティティが作成できる",
			input1:    tdDomain.ID,
			input2:    tdDomain.Name,
			input3:    tdDomain.HashedAuthToken,
			expected1: tdDomain.Entity,
			expected2: nil,
		},
		{
			title:     "【異常系】ユーザーアカウントIDがzero値",
			input1:    user.ID{},
			input2:    tdDomain.Name,
			input3:    tdDomain.HashedAuthToken,
			expected1: user.User{},
			expected2: errors.Internal("user id is nil"),
		},
		{
			title:     "【異常系】ユーザーアカウント名がzero値",
			input1:    tdDomain.ID,
			input2:    "",
			input3:    tdDomain.HashedAuthToken,
			expected1: user.User{},
			expected2: errors.Internal("user name is nil"),
		},
		{
			title:     "【異常系】認証トークンハッシュ値がzero値",
			input1:    tdDomain.ID,
			input2:    tdDomain.Name,
			input3:    nil,
			expected1: user.User{},
			expected2: errors.Internal("user hashed_auth_token is nil"),
		},
	}

	for _, td := range tests {
		td := td

		t.Run("New:"+td.title, func(t *testing.T) {
			t.Parallel()

			output1, output2 := user.New(td.input1, td.input2, td.input3)

			assert.Equal(t, td.expected1, output1)
			assert.Equal(t, td.expected2, output2)
		})
	}
}
