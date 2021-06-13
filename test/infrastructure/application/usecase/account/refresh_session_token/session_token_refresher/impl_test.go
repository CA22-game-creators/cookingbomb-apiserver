package auth_test

import (
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/oklog/ulid/v2"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	applicaiton "github.com/CA22-game-creators/cookingbomb-apiserver/application/usecase/account/refresh_session_token"
	domain "github.com/CA22-game-creators/cookingbomb-apiserver/domain/model/user"
	infra "github.com/CA22-game-creators/cookingbomb-apiserver/infrastructure/application/usecase/account/refresh_session_token/session_token_refresher"

	tdDomain "github.com/CA22-game-creators/cookingbomb-apiserver/test/testdata/domain/user"
	tdString "github.com/CA22-game-creators/cookingbomb-apiserver/test/testdata/string/common"
)

type testHandler struct {
	sessionTokenRefresher applicaiton.SessionTokenRefresher

	db      *gorm.DB
	sqlmock sqlmock.Sqlmock
}

func TestHandle(t *testing.T) {
	t.Parallel()

	expiredAt := time.Now().Add(20 * time.Hour)

	tests := []struct {
		title    string
		before   func(testHandler)
		input1   domain.User
		input2   uuid.UUID
		input3   time.Time
		expected error
	}{
		{
			title: "【正常系】sessionTokenをUpsertできる",
			before: func(h testHandler) {
				h.sqlmock.ExpectBegin()
				h.sqlmock.ExpectExec(
					regexp.QuoteMeta("UPDATE `sessions` SET `session_token`=?,`expired_at`=? WHERE `user_id` = ?"),
				).WithArgs(
					tdString.UUID.Valid,
					expiredAt,
					ulid.ULID(tdDomain.Entity.ID).String(),
				).WillReturnResult(sqlmock.NewResult(1, 1))
				h.sqlmock.ExpectCommit()
			},
			input1:   tdDomain.Entity,
			input2:   uuid.MustParse(tdString.UUID.Valid),
			input3:   expiredAt,
			expected: nil,
		},
	}

	for _, td := range tests {
		td := td

		t.Run("Handle:"+td.title, func(t *testing.T) {
			t.Parallel()

			var tester testHandler
			tester.setupTest(t)
			td.before(tester)

			output := tester.sessionTokenRefresher.Handle(td.input1, td.input2, td.input3)

			assert.Equal(t, td.expected, output)
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
	h.sessionTokenRefresher = infra.New(h.db)
}
