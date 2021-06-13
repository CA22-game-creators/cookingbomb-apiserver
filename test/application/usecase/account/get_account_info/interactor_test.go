package account_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	auth "github.com/CA22-game-creators/cookingbomb-apiserver/application/common/auth/authenticate"
	application "github.com/CA22-game-creators/cookingbomb-apiserver/application/usecase/account/get_account_info"

	mockAuth "github.com/CA22-game-creators/cookingbomb-apiserver/mock/application/common/auth/authenticate"
	tdDomain "github.com/CA22-game-creators/cookingbomb-apiserver/test/testdata/domain/user"
	tdString "github.com/CA22-game-creators/cookingbomb-apiserver/test/testdata/string/common"
)

type testHandler struct {
	interactor application.InputPort

	auth *mockAuth.MockInputPort
}

func TestHandle(t *testing.T) {
	t.Parallel()

	tests := []struct {
		title    string
		before   func(testHandler)
		input    application.InputData
		expected application.OutputData
	}{
		{
			title: "【正常系】sessionTokenからユーザーをアカウントを取得する",
			before: func(h testHandler) {
				h.auth.EXPECT().Handle(
					auth.InputData{SessionToken: tdString.UUID.Valid},
				).Return(
					auth.OutputData{Account: tdDomain.Entity},
				)
			},
			input: application.InputData{
				SessionToken: tdString.UUID.Valid,
			},
			expected: application.OutputData{
				Account: tdDomain.Entity,
			},
		},
	}

	for _, td := range tests {
		td := td

		t.Run("Handle:"+td.title, func(t *testing.T) {
			t.Parallel()

			var tester testHandler
			tester.setupTest(t)

			td.before(tester)

			output := tester.interactor.Handle(td.input)

			assert.Equal(t, td.expected, output)
		})
	}
}

func (h *testHandler) setupTest(t *testing.T) {
	ctrl := gomock.NewController(t)
	h.auth = mockAuth.NewMockInputPort(ctrl)

	h.interactor = application.New(h.auth)
}
