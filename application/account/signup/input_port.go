//go:generate mockgen -source=$GOFILE -destination=../../../mock/application/account/signup/$GOFILE
package account

type InputPort interface {
	Handle(InputData) OutputData
}
