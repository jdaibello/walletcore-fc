package gateway

import "github.com/jdaibello/walletcore-fc/internal/entity"

type AccountGateway interface {
	Save(account *entity.Account) error
	FindByID(id string) (*entity.Account, error)
}
