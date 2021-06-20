package account_test

import (
	"context"
	"testing"

	pb "github.com/CA22-game-creators/cookingbomb-proto/server/pb/api"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	getAccountInfo "github.com/CA22-game-creators/cookingbomb-apiserver/application/usecase/account/get_account_info"
	refreshSessionToken "github.com/CA22-game-creators/cookingbomb-apiserver/application/usecase/account/refresh_session_token"
	signup "github.com/CA22-game-creators/cookingbomb-apiserver/application/usecase/account/signup"
	controller "github.com/CA22-game-creators/cookingbomb-apiserver/presentation/account"

	mockGetAccountInfo "github.com/CA22-game-creators/cookingbomb-apiserver/mock/application/usecase/account/get_account_info"
	mockRefreshSessionToken "github.com/CA22-game-creators/cookingbomb-apiserver/mock/application/usecase/account/refresh_session_token"
	mockSignup "github.com/CA22-game-creators/cookingbomb-apiserver/mock/application/usecase/account/signup"
	tdDomain "github.com/CA22-game-creators/cookingbomb-apiserver/test/testdata/domain/user"
	tdCommonString "github.com/CA22-game-creators/cookingbomb-apiserver/test/testdata/string/common"
	tdUserString "github.com/CA22-game-creators/cookingbomb-apiserver/test/testdata/string/user"
)

type testHandler struct {
	controller pb.AccountServicesServer

	context             context.Context
	signup              *mockSignup.MockInputPort
	refreshSessionToken *mockRefreshSessionToken.MockInputPort
	getAccountInfo      *mockGetAccountInfo.MockInputPort
}

func TestSignup(t *testing.T) {
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
				input := signup.InputData{Name: tdUserString.Name.Valid}
				output := signup.OutputData{
					Account:   tdDomain.Entity,
					AuthToken: uuid.MustParse(tdCommonString.UUID.Valid),
				}
				h.signup.EXPECT().Handle(input).Return(output)
			},
			input: &pb.SignupRequest{Name: tdUserString.Name.Valid},
			expected1: &pb.SignupResponse{
				AccountInfo: &pb.AccountInfo{
					Id:   tdCommonString.ULID.Valid,
					Name: tdUserString.Name.Valid,
				},
				AuthToken: tdCommonString.UUID.Valid,
			},
			expected2: nil,
		},
		{
			title:     "【異常系】ユーザー名が短すぎる",
			input:     &pb.SignupRequest{Name: tdUserString.Name.TooShort},
			expected1: nil,
			expected2: status.Error(codes.InvalidArgument, "ユーザー名は半角英数字か日本語の1-10文字である必要があります"),
		},
		{
			title:     "【異常系】ユーザー名が長すぎる",
			input:     &pb.SignupRequest{Name: tdUserString.Name.TooLong},
			expected1: nil,
			expected2: status.Error(codes.InvalidArgument, "ユーザー名は半角英数字か日本語の1-10文字である必要があります"),
		},
		{
			title:     "【異常系】ユーザー名が不正な文字コード",
			input:     &pb.SignupRequest{Name: tdUserString.Name.InvalidChar},
			expected1: nil,
			expected2: status.Error(codes.InvalidArgument, "ユーザー名は半角英数字か日本語の1-10文字である必要があります"),
		},
		{
			title:     "【異常系】ユーザー名に記号とか入ってる",
			input:     &pb.SignupRequest{Name: tdUserString.Name.Invalid},
			expected1: nil,
			expected2: status.Error(codes.InvalidArgument, "ユーザー名は半角英数字か日本語の1-10文字である必要があります"),
		},
	}

	for _, td := range tests {
		td := td

		t.Run("Signup:"+td.title, func(t *testing.T) {
			t.Parallel()

			var tester testHandler
			tester.setupTest(t)
			if td.before != nil {
				td.before(tester)
			}

			output1, output2 := tester.controller.Signup(tester.context, td.input)

			assert.Equal(t, td.expected1, output1)
			assert.Equal(t, td.expected2, output2)
		})
	}
}

func TestGetAccountInfo(t *testing.T) {
	t.Parallel()

	tests := []struct {
		title     string
		before    func(testHandler)
		input     *pb.GetAccountInfoRequest
		expected1 *pb.GetAccountInfoResponse
		expected2 error
	}{
		{
			title: "【正常系】sessionTokenからアカウント情報を取得できる",
			before: func(h testHandler) {
				input := getAccountInfo.InputData{SessionToken: tdCommonString.UUID.Valid}
				output := getAccountInfo.OutputData{Account: tdDomain.Entity}
				h.getAccountInfo.EXPECT().Handle(input).Return(output)
			},
			input: &pb.GetAccountInfoRequest{
				SessionToken: tdCommonString.UUID.Valid,
			},
			expected1: &pb.GetAccountInfoResponse{
				AccountInfo: &pb.AccountInfo{
					Id:   tdCommonString.ULID.Valid,
					Name: tdUserString.Name.Valid,
				},
			},
			expected2: nil,
		},
		{
			title: "【異常系】sessionTokenが不正値",
			input: &pb.GetAccountInfoRequest{
				SessionToken: tdCommonString.UUID.Invalid,
			},
			expected1: nil,
			expected2: status.Error(codes.InvalidArgument, "sessionTokenが不正な形式です"),
		},
	}

	for _, td := range tests {
		td := td

		t.Run("GetAccountInfo:"+td.title, func(t *testing.T) {
			t.Parallel()

			var tester testHandler
			tester.setupTest(t)
			if td.before != nil {
				td.before(tester)
			}

			output1, output2 := tester.controller.GetAccountInfo(tester.context, td.input)

			assert.Equal(t, td.expected1, output1)
			assert.Equal(t, td.expected2, output2)
		})
	}
}
func TestGetSessionToken(t *testing.T) {
	t.Parallel()

	tests := []struct {
		title     string
		before    func(testHandler)
		input     *pb.GetSessionTokenRequest
		expected1 *pb.GetSessionTokenResponse
		expected2 error
	}{
		{
			title: "【正常系】UserIDとAuthTokenからSessionTokenを取得できる",
			before: func(h testHandler) {
				input := refreshSessionToken.InputData{
					UserID:    tdCommonString.ULID.Valid,
					AuthToken: tdCommonString.UUID.Valid,
				}
				output := refreshSessionToken.OutputData{SessionToken: uuid.MustParse(tdCommonString.UUID.Valid)}
				h.refreshSessionToken.EXPECT().Handle(input).Return(output)
			},
			input: &pb.GetSessionTokenRequest{
				UserId:    tdCommonString.ULID.Valid,
				AuthToken: tdCommonString.UUID.Valid,
			},
			expected1: &pb.GetSessionTokenResponse{
				SessionToken: tdCommonString.UUID.Valid,
			},
			expected2: nil,
		},
		{
			title: "【異常系】userIDが不正値",
			input: &pb.GetSessionTokenRequest{
				UserId:    tdCommonString.ULID.Invalid,
				AuthToken: tdCommonString.UUID.Valid,
			},
			expected1: nil,
			expected2: status.Error(codes.InvalidArgument, "ユーザーIDが不正な形式です"),
		},
		{
			title: "【異常系】authTokenが不正値",
			input: &pb.GetSessionTokenRequest{
				UserId:    tdCommonString.ULID.Valid,
				AuthToken: tdCommonString.UUID.Invalid,
			},
			expected1: nil,
			expected2: status.Error(codes.InvalidArgument, "authTokenが不正な形式です"),
		},
		{
			title: "【異常系】userIDもauthTokenも不正値",
			input: &pb.GetSessionTokenRequest{
				UserId:    tdCommonString.ULID.Invalid,
				AuthToken: tdCommonString.UUID.Invalid,
			},
			expected1: nil,
			expected2: status.Error(codes.InvalidArgument, "ユーザーIDが不正な形式です\nauthTokenが不正な形式です"),
		},
	}

	for _, td := range tests {
		td := td

		t.Run("GetSessionToken:"+td.title, func(t *testing.T) {
			t.Parallel()

			var tester testHandler
			tester.setupTest(t)
			if td.before != nil {
				td.before(tester)
			}

			output1, output2 := tester.controller.GetSessionToken(tester.context, td.input)

			assert.Equal(t, td.expected1, output1)
			assert.Equal(t, td.expected2, output2)
		})
	}
}

func (h *testHandler) setupTest(t *testing.T) {
	h.context = context.TODO()

	ctrl := gomock.NewController(t)
	h.signup = mockSignup.NewMockInputPort(ctrl)
	h.refreshSessionToken = mockRefreshSessionToken.NewMockInputPort(ctrl)
	h.getAccountInfo = mockGetAccountInfo.NewMockInputPort(ctrl)

	h.controller = controller.New(h.signup, h.refreshSessionToken, h.getAccountInfo)
}
