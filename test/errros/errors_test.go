package errors_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/CA22-game-creators/cookingbomb-apiserver/errors"
)

func TestInvalidArgument(t *testing.T) {
	t.Parallel()

	tests := []struct {
		title    string
		input    string
		expected error
	}{
		{
			title:    "【正常系】バッドリクエスト(InvalidArgument)エラーを作成できる",
			input:    "error",
			expected: status.Error(codes.InvalidArgument, "error"),
		},
	}

	for _, td := range tests {
		td := td

		t.Run("InvalidArgument:"+td.title, func(t *testing.T) {
			t.Parallel()

			output := errors.InvalidArgument(td.input)

			assert.Equal(t, td.expected, output)
		})
	}
}

func TestInternal(t *testing.T) {
	t.Parallel()

	tests := []struct {
		title    string
		input    string
		expected error
	}{
		{
			title:    "【正常系】インターナルサーバー(Internal)エラーを作成できる",
			input:    "error",
			expected: status.Error(codes.Internal, "error"),
		},
	}

	for _, td := range tests {
		td := td

		t.Run("Internal:"+td.title, func(t *testing.T) {
			t.Parallel()

			output := errors.Internal(td.input)

			assert.Equal(t, td.expected, output)
		})
	}
}
