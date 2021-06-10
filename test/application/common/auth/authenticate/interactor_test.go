package auth_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	application "github.com/CA22-game-creators/cookingbomb-apiserver/application/common/auth/authenticate"
	domain "github.com/CA22-game-creators/cookingbomb-apiserver/domain/model/user"
	"github.com/CA22-game-creators/cookingbomb-apiserver/errors"

	mockAuth "github.com/CA22-game-creators/cookingbomb-apiserver/mock/application/common/auth/authenticate"
	mockDomain "github.com/CA22-game-creators/cookingbomb-apiserver/mock/domain/model/user"
	tdDomain "github.com/CA22-game-creators/cookingbomb-apiserver/test/testdata/domain/user"
	tdString "github.com/CA22-game-creators/cookingbomb-apiserver/test/testdata/string/common"
)

type testHandler struct {
	interactor application.InputPort

	repository       *mockDomain.MockRepository
	accountIDFetcher *mockAuth.MockAccountIDFetcher
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
			title: "【正常系】sessionTokenからアカウントを取得する",
			before: func(h testHandler) {
				h.accountIDFetcher.EXPECT().Handle(tdString.UUID.Valid).Return(tdDomain.ID, nil)
				h.repository.EXPECT().Find(tdDomain.ID).Return(tdDomain.Entity, nil)
			},
			input: application.InputData{
				SessionToken: tdString.UUID.Valid,
			},
			expected: application.OutputData{
				Account: tdDomain.Entity,
			},
		},
		{
			title: "【異常系】sessionTokenに紐づくアカウントがなかった",
			before: func(h testHandler) {
				h.accountIDFetcher.EXPECT().Handle(gomock.Any()).Return(domain.ID{}, nil)
			},
			input: application.InputData{
				SessionToken: uuid.Must(uuid.NewRandom()).String(),
			},
			expected: application.OutputData{
				Err: errors.Unauthenticated("session not found"),
			},
		},
		{
			title: "【異常系】sessionはあったけどそれに紐づくUserがいなかった（起こり得ない）",
			before: func(h testHandler) {
				h.accountIDFetcher.EXPECT().Handle(tdString.UUID.Valid).Return(tdDomain.ID, nil)
				h.repository.EXPECT().Find(tdDomain.ID).Return(domain.User{}, nil)
			},
			input: application.InputData{
				SessionToken: tdString.UUID.Valid,
			},
			expected: application.OutputData{
				Err: errors.Internal("user not found"),
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
	h.repository = mockDomain.NewMockRepository(ctrl)
	h.accountIDFetcher = mockAuth.NewMockAccountIDFetcher(ctrl)

	h.interactor = application.New(h.repository, h.accountIDFetcher)
}
