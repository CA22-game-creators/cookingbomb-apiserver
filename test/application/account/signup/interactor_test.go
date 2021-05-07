package account_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	application "github.com/CA22-game-creators/cookingbomb-apiserver/application/account/signup"

	mockDomain "github.com/CA22-game-creators/cookingbomb-apiserver/mock/domain/model/user"
	tdDomain "github.com/CA22-game-creators/cookingbomb-apiserver/test/testdata/domain/user"
	tdCommonString "github.com/CA22-game-creators/cookingbomb-apiserver/test/testdata/string/common"
	tdUserString "github.com/CA22-game-creators/cookingbomb-apiserver/test/testdata/string/user"
)

type testHandler struct {
	interactor application.InputPort

	factory    *mockDomain.MockFactory
	repository *mockDomain.MockRepository
}

func TestCreate(t *testing.T) {
	t.Parallel()

	tests := []struct {
		title    string
		before   func(testHandler)
		input    application.InputData
		expected application.OutputData
	}{
		{
			title: "【正常系】ユーザーを登録する",
			before: func(h testHandler) {
				h.factory.EXPECT().Create(tdDomain.Name).Return(
					tdDomain.Entity,
					uuid.MustParse(tdCommonString.UUID.Valid),
					nil,
				)
				h.repository.EXPECT().Save(tdDomain.Entity).Return(nil)
			},
			input: application.InputData{Name: tdUserString.Name.Valid},
			expected: application.OutputData{
				Account:   tdDomain.Entity,
				AuthToken: uuid.MustParse(tdCommonString.UUID.Valid),
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
	h.factory = mockDomain.NewMockFactory(ctrl)
	h.repository = mockDomain.NewMockRepository(ctrl)

	h.interactor = application.New(h.factory, h.repository)
}
