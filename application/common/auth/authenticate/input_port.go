//go:generate mockgen -source=$GOFILE -destination=../../../../mock/application/common/auth/authenticate/$GOFILE
package auth

type InputPort interface {
	Handle(InputData) OutputData
}
