package user_test

import (
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	domain "github.com/CA22-game-creators/cookingbomb-apiserver/domain/model/user"
	repoImpl "github.com/CA22-game-creators/cookingbomb-apiserver/infrastructure/repository/user"

	tdDomain "github.com/CA22-game-creators/cookingbomb-apiserver/test/testdata/domain/user"
	tdCommonString "github.com/CA22-game-creators/cookingbomb-apiserver/test/testdata/string/common"
	tdUserString "github.com/CA22-game-creators/cookingbomb-apiserver/test/testdata/string/user"
)

type testHandler struct {
	repository domain.Repository

	db      *gorm.DB
	sqlmock sqlmock.Sqlmock
}

func TestSave(t *testing.T) {
	t.Parallel()

	tests := []struct {
		title    string
		before   func(testHandler)
		input    domain.User
		expected error
	}{
		{
			title: "【正常系】ユーザーをDBに保存できる",
			before: func(h testHandler) {
				h.sqlmock.ExpectBegin()
				h.sqlmock.ExpectExec(
					regexp.QuoteMeta("INSERT INTO `users` (`id`,`name`,`hashed_auth_token`) VALUES (?,?,?)"),
				).WithArgs(
					tdCommonString.ULID.Valid,
					tdUserString.Name.Valid,
					tdCommonString.UUID.Encrypted,
				).WillReturnResult(sqlmock.NewResult(1, 1))
				h.sqlmock.ExpectCommit()
			},
			input:    tdDomain.Entity,
			expected: nil,
		},
	}

	for _, td := range tests {
		td := td

		t.Run("Save:"+td.title, func(t *testing.T) {
			t.Parallel()

			var tester testHandler
			tester.setupTest(t)
			td.before(tester)

			output := tester.repository.Save(td.input)

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
	h.repository = repoImpl.NewRepository(h.db)
}
