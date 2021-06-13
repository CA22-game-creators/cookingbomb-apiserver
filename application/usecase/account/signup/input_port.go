//go:generate mockgen -source=$GOFILE -destination=../../../../mock/application/usecase/account/signup/$GOFILE
package account

type InputPort interface {
	Handle(InputData) OutputData
}
