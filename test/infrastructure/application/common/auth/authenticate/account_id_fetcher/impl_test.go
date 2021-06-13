package auth_test

import (
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/oklog/ulid/v2"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	applicaiton "github.com/CA22-game-creators/cookingbomb-apiserver/application/common/auth/authenticate"
	domain "github.com/CA22-game-creators/cookingbomb-apiserver/domain/model/user"
	"github.com/CA22-game-creators/cookingbomb-apiserver/errors"
	infra "github.com/CA22-game-creators/cookingbomb-apiserver/infrastructure/application/common/auth/authenticate/account_id_fetcher"

	tdDomain "github.com/CA22-game-creators/cookingbomb-apiserver/test/testdata/domain/user"
	tdString "github.com/CA22-game-creators/cookingbomb-apiserver/test/testdata/string/common"
)

type testHandler struct {
	accountIDFetcher applicaiton.AccountIDFetcher

	db      *gorm.DB
	sqlmock sqlmock.Sqlmock
}

func TestHandle(t *testing.T) {
	t.Parallel()

	tests := []struct {
		title     string
		before    func(testHandler)
		input     string
		expected1 domain.ID
		expected2 error
	}{
		{
			title: "【正常系】sessionTokenに紐づくユーザーIDを取得",
			before: func(h testHandler) {
				h.sqlmock.ExpectQuery(
					regexp.QuoteMeta("SELECT * FROM `sessions` WHERE session_token = ? ORDER BY `sessions`.`user_id` LIMIT 1"),
				).WithArgs(
					tdString.UUID.Valid,
				).WillReturnRows(
					sqlmock.NewRows(
						[]string{"user_id", "session_token", "expired_at"},
					).AddRow(ulid.ULID(tdDomain.ID).String(), tdString.UUID.Valid, time.Now().Add(1*time.Second)),
				)
			},
			input:     tdString.UUID.Valid,
			expected1: tdDomain.ID,
			expected2: nil,
		},
		{
			title: "【正常系】sessionTokenに紐づくユーザーIDが存在しない",
			before: func(h testHandler) {
				h.sqlmock.ExpectQuery(
					regexp.QuoteMeta("SELECT * FROM `sessions` WHERE session_token = ? ORDER BY `sessions`.`user_id` LIMIT 1"),
				).WillReturnError(gorm.ErrRecordNotFound)
			},
			input:     tdString.UUID.Valid,
			expected1: domain.ID{},
			expected2: nil,
		},
		{
			title: "【異常系】sessionがタイムアウトしていた",
			before: func(h testHandler) {
				h.sqlmock.ExpectQuery(
					regexp.QuoteMeta("SELECT * FROM `sessions` WHERE session_token = ? ORDER BY `sessions`.`user_id` LIMIT 1"),
				).WithArgs(
					tdString.UUID.Valid,
				).WillReturnRows(
					sqlmock.NewRows(
						[]string{"user_id", "session_token", "expired_at"},
					).AddRow(ulid.ULID(tdDomain.ID).String(), tdString.UUID.Valid, time.Now().Add(-1*time.Second)),
				)
			},
			input:     tdString.UUID.Valid,
			expected1: domain.ID{},
			expected2: errors.Unauthenticated("session timeout"),
		},
	}

	for _, td := range tests {
		td := td

		t.Run("Handle:"+td.title, func(t *testing.T) {
			t.Parallel()

			var tester testHandler
			tester.setupTest(t)
			td.before(tester)

			output1, output2 := tester.accountIDFetcher.Handle(td.input)

			assert.Equal(t, td.expected1, output1)
			assert.Equal(t, td.expected2, output2)
		})
	}
}

func (h *testHandler) setupTest(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	gormDB, err := gorm.Open(
		mysql.Dialector{Config: &mysql.Config{
			DriverName:                "mysql",
			Conn:                      db,
			SkipInitializeWithVersion: true,
		}}, &gorm.Config{})
	assert.NoError(t, err)

	h.db = gormDB
	h.sqlmock = mock
	h.accountIDFetcher = infra.New(h.db)
}
