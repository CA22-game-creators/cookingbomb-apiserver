package account_test

import (
	"crypto/rand"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/oklog/ulid/v2"
	"github.com/stretchr/testify/assert"

	application "github.com/CA22-game-creators/cookingbomb-apiserver/application/usecase/account/refresh_session_token"
	domain "github.com/CA22-game-creators/cookingbomb-apiserver/domain/model/user"
	"github.com/CA22-game-creators/cookingbomb-apiserver/errors"

	mockApplication "github.com/CA22-game-creators/cookingbomb-apiserver/mock/application/usecase/account/refresh_session_token"
	mockDomain "github.com/CA22-game-creators/cookingbomb-apiserver/mock/domain/model/user"
	mockUtil "github.com/CA22-game-creators/cookingbomb-apiserver/mock/util"
	tdDomain "github.com/CA22-game-creators/cookingbomb-apiserver/test/testdata/domain/user"
	tdString "github.com/CA22-game-creators/cookingbomb-apiserver/test/testdata/string/common"
)

type testHandler struct {
	interactor application.InputPort

	repository            *mockDomain.MockRepository
	sessionTokenRefresher *mockApplication.MockSessionTokenRefresher
	tokenGenerator        *mockUtil.MockTokenGenerator
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
			title: "【正常系】UserIDとAuthTokenから、SessionTokenを更新・取得",
			before: func(h testHandler) {
				h.repository.EXPECT().Find(tdDomain.ID).Return(tdDomain.Entity, nil)
				h.tokenGenerator.EXPECT().Generate().Return(uuid.MustParse(tdString.UUID.Valid), nil)
				h.sessionTokenRefresher.EXPECT().Handle(tdDomain.Entity, uuid.MustParse(tdString.UUID.Valid), gomock.Any()).Return(nil)
			},
			input: application.InputData{
				UserID:    tdString.ULID.Valid,
				AuthToken: tdString.UUID.Valid,
			},
			expected: application.OutputData{
				SessionToken: uuid.MustParse(tdString.UUID.Valid),
			},
		},
		{
			title: "【異常系】UserIDに対応するユーザーが存在しない",
			before: func(h testHandler) {
				h.repository.EXPECT().Find(gomock.Any()).Return(domain.User{}, nil)
			},
			input: application.InputData{
				UserID:    ulid.MustNew(ulid.Timestamp(time.Unix(1000000, 0)), rand.Reader).String(),
				AuthToken: tdString.UUID.Valid,
			},
			expected: application.OutputData{
				Err: errors.InvalidArgument("user not found"),
			},
		},
		{
			title: "【異常系】UserIDに対応したAuthTokenではない",
			before: func(h testHandler) {
				h.repository.EXPECT().Find(tdDomain.ID).Return(tdDomain.Entity, nil)
			},
			input: application.InputData{
				UserID:    tdString.ULID.Valid,
				AuthToken: uuid.Must(uuid.NewRandom()).String(),
			},
			expected: application.OutputData{
				Err: errors.Unauthenticated("crypto/bcrypt: hashedPassword is not the hash of the given password"),
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
	h.sessionTokenRefresher = mockApplication.NewMockSessionTokenRefresher(ctrl)
	h.tokenGenerator = mockUtil.NewMockTokenGenerator(ctrl)

	h.interactor = application.New(h.repository, h.sessionTokenRefresher, h.tokenGenerator)
}
