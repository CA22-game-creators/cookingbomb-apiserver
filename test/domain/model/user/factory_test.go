package user_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/oklog/ulid/v2"
	"github.com/stretchr/testify/assert"

	"github.com/CA22-game-creators/cookingbomb-apiserver/domain/model/user"

	mockUtil "github.com/CA22-game-creators/cookingbomb-apiserver/mock/util"
	tdDomain "github.com/CA22-game-creators/cookingbomb-apiserver/test/testdata/domain/user"
	tdString "github.com/CA22-game-creators/cookingbomb-apiserver/test/testdata/string/common"
)

type testHandler struct {
	factory user.Factory

	idGenerator    *mockUtil.MockIDGenerator
	tokenGenerator *mockUtil.MockTokenGenerator
	cryptoManager  *mockUtil.MockCryptoManager
}

func TestCreate(t *testing.T) {
	t.Parallel()

	tests := []struct {
		title     string
		before    func(testHandler)
		input     user.Name
		expected1 user.User
		expected2 uuid.UUID
		expected3 error
	}{
		{
			title: "【正常系】ユーザーアカウントエンティティを生成できる",
			before: func(h testHandler) {
				h.idGenerator.EXPECT().Generate().Return(ulid.MustParse(tdString.ULID.Valid), nil)
				h.tokenGenerator.EXPECT().Generate().Return(uuid.MustParse(tdString.UUID.Valid), nil)
				h.cryptoManager.EXPECT().Encrypt(tdString.UUID.Valid).Return([]byte(tdString.UUID.Encrypted), nil)
			},
			input:     tdDomain.Name,
			expected1: tdDomain.Entity,
			expected2: uuid.MustParse(tdString.UUID.Valid),
			expected3: nil,
		},
	}

	for _, td := range tests {
		td := td

		t.Run("Create:"+td.title, func(t *testing.T) {
			t.Parallel()

			var tester testHandler
			tester.setupTest(t)
			td.before(tester)

			output1, output2, output3 := tester.factory.Create(td.input)

			assert.Equal(t, td.expected1, output1)
			assert.Equal(t, td.expected2, output2)
			assert.Equal(t, td.expected3, output3)
		})
	}
}

func (h *testHandler) setupTest(t *testing.T) {
	ctrl := gomock.NewController(t)
	h.idGenerator = mockUtil.NewMockIDGenerator(ctrl)
	h.tokenGenerator = mockUtil.NewMockTokenGenerator(ctrl)
	h.cryptoManager = mockUtil.NewMockCryptoManager(ctrl)

	h.factory = user.NewFactory(h.idGenerator, h.tokenGenerator, h.cryptoManager)
}
