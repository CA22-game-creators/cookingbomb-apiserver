package account_test

import (
	"context"
	"testing"

	pb "github.com/CA22-game-creators/cookingbomb-proto/server/pb"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	signupApplication "github.com/CA22-game-creators/cookingbomb-apiserver/application/account/signup"
	controller "github.com/CA22-game-creators/cookingbomb-apiserver/presentation/account"

	mockSignupApplication "github.com/CA22-game-creators/cookingbomb-apiserver/mock/application/account/signup"
	tdDomain "github.com/CA22-game-creators/cookingbomb-apiserver/test/testdata/domain/user"
	tdCommonString "github.com/CA22-game-creators/cookingbomb-apiserver/test/testdata/string/common"
	tdUserString "github.com/CA22-game-creators/cookingbomb-apiserver/test/testdata/string/user"
)

type testHandler struct {
	controller pb.AccountServicesServer

	context           context.Context
	signupApplication *mockSignupApplication.MockInputPort
}

func TestCreate(t *testing.T) {
	t.Parallel()

	tests := []struct {
		title     string
		before    func(testHandler)
		input     *pb.SignupRequest
		expected1 *pb.SignupResponse
		expected2 error
	}{
		{
			title: "【正常系】ユーザーを登録する",
			before: func(h testHandler) {
				input := signupApplication.InputData{Name: tdUserString.Name.Valid}
				output := signupApplication.OutputData{
					Account:   tdDomain.Entity,
					AuthToken: uuid.MustParse(tdCommonString.UUID.Valid),
				}
				h.signupApplication.EXPECT().Handle(input).Return(output)
			},
			input: &pb.SignupRequest{Name: tdUserString.Name.Valid},
			expected1: &pb.SignupResponse{
				Account: &pb.Account{
					Id:        tdCommonString.ULID.Valid,
					Name:      tdUserString.Name.Valid,
					AuthToken: tdCommonString.UUID.Valid,
				},
			},
			expected2: nil,
		},
	}

	for _, td := range tests {
		td := td

		t.Run("Handle:"+td.title, func(t *testing.T) {
			t.Parallel()

			var tester testHandler
			tester.setupTest(t)
			td.before(tester)

			output1, output2 := tester.controller.Signup(tester.context, td.input)

			assert.Equal(t, td.expected1, output1)
			assert.Equal(t, td.expected2, output2)
		})
	}
}

func (h *testHandler) setupTest(t *testing.T) {
	h.context = context.TODO()

	ctrl := gomock.NewController(t)
	h.signupApplication = mockSignupApplication.NewMockInputPort(ctrl)

	h.controller = controller.New(h.signupApplication)
}
