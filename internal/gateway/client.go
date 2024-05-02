package gateway

import "github.com/jdaibello/walletcore-fc/internal/entity"

type ClientGateway interface {
	Get(id string) (*entity.Client, error)
	Save(client *entity.Client) error
}
