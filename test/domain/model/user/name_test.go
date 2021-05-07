package user_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/CA22-game-creators/cookingbomb-apiserver/domain/model/user"

	tdDomain "github.com/CA22-game-creators/cookingbomb-apiserver/test/testdata/domain/user"
	tdString "github.com/CA22-game-creators/cookingbomb-apiserver/test/testdata/string/user"
)

func TestNewName(t *testing.T) {
	t.Parallel()

	tests := []struct {
		title     string
		input     string
		expected1 user.Name
		expected2 error
	}{
		{
			title:     "【正常系】ユーザー名の値オブジェクトが作成できる",
			input:     tdString.Name.Valid,
			expected1: tdDomain.Name,
			expected2: nil,
		},
		{
			title:     "【異常系】ユーザー名がUTF-8でエンコードされた文字列ではない",
			input:     tdString.Name.Invalid,
			expected1: "",
			expected2: errors.New("user name string is invalid"),
		},
		{
			title:     "【異常系】ユーザー名が長すぎる",
			input:     tdString.Name.TooLong,
			expected1: "",
			expected2: fmt.Errorf(
				"user name should be %d to %d characters", user.NameMinLen, user.NameMaxLen),
		},
		{
			title:     "【異常系】ユーザー名が短すぎる",
			input:     tdString.Name.TooShort,
			expected1: "",
			expected2: fmt.Errorf(
				"user name should be %d to %d characters", user.NameMinLen, user.NameMaxLen),
		},
	}

	for _, td := range tests {
		td := td

		t.Run("NewName:"+td.title, func(t *testing.T) {
			t.Parallel()

			output1, output2 := user.NewName(td.input)

			assert.Equal(t, td.expected1, output1)
			assert.Equal(t, td.expected2, output2)
		})
	}
}
