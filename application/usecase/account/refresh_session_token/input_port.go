//go:generate mockgen -source=$GOFILE -destination=../../../../mock/application/usecase/account/refresh_session_token/$GOFILE
package account

type InputPort interface {
	Handle(InputData) OutputData
}
