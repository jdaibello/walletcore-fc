package gateway

import "github.com/jdaibello/walletcore-fc/internal/entity"

type TransactionGateway interface {
	Create(transaction *entity.Transaction) error
}
