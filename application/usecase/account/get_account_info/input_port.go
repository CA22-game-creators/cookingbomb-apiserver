//go:generate mockgen -source=$GOFILE -destination=../../../../mock/application/usecase/account/get_account_info/$GOFILE
package account

type InputPort interface {
	Handle(InputData) OutputData
}
