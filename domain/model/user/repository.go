//go:generate mockgen -source=$GOFILE -destination=../../../mock/domain/model/user/$GOFILE
package user

type Repository interface {
	Save(User) error
	Find(ID) (User, error)
}
