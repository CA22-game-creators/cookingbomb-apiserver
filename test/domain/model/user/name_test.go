package user_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/CA22-game-creators/cookingbomb-apiserver/domain/model/user"
	"github.com/CA22-game-creators/cookingbomb-apiserver/errors"

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
			title:     "【異常系】ユーザー名が予定されていない文字列",
			input:     tdString.Name.Invalid,
			expected1: "",
			expected2: errors.InvalidArgument("ユーザー名は半角英数字か日本語の1-10文字である必要があります"),
		},
		{
			title:     "【異常系】ユーザー名が不正文字コード",
			input:     tdString.Name.InvalidChar,
			expected1: "",
			expected2: errors.InvalidArgument("ユーザー名は半角英数字か日本語の1-10文字である必要があります"),
		},
		{
			title:     "【異常系】ユーザー名が長すぎる",
			input:     tdString.Name.TooLong,
			expected1: "",
			expected2: errors.InvalidArgument("ユーザー名は半角英数字か日本語の1-10文字である必要があります"),
		},
		{
			title:     "【異常系】ユーザー名が短すぎる",
			input:     tdString.Name.TooShort,
			expected1: "",
			expected2: errors.InvalidArgument("ユーザー名は半角英数字か日本語の1-10文字である必要があります"),
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
