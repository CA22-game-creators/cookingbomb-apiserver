package user_test

import (
	"errors"
	"testing"

	"github.com/oklog/ulid/v2"
	"github.com/stretchr/testify/assert"

	"github.com/CA22-game-creators/cookingbomb-apiserver/domain/model/user"

	tdDomain "github.com/CA22-game-creators/cookingbomb-apiserver/test/testdata/domain/user"
	tdString "github.com/CA22-game-creators/cookingbomb-apiserver/test/testdata/string/common"
)

func TestNewID(t *testing.T) {
	t.Parallel()

	tests := []struct {
		title     string
		input     ulid.ULID
		expected1 user.ID
		expected2 error
	}{
		{
			title:     "【正常系】ユーザーIDの値オブジェクトが作成できる",
			input:     ulid.MustParse(tdString.ULID.Valid),
			expected1: tdDomain.ID,
			expected2: nil,
		},
		{
			title:     "【異常系】ULIDがzero値",
			input:     ulid.ULID{},
			expected1: user.ID{},
			expected2: errors.New("user id is nil"),
		},
	}

	for _, td := range tests {
		td := td

		t.Run("NewID:"+td.title, func(t *testing.T) {
			t.Parallel()

			output1, output2 := user.NewID(td.input)

			assert.Equal(t, td.expected1, output1)
			assert.Equal(t, td.expected2, output2)
		})
	}
}
