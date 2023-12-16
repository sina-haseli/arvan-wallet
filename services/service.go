package services

import (
	"arvanwallet/repositories"
	"arvanwallet/services/consumer"
	"arvanwallet/services/consumer/redis"
	"arvanwallet/services/wallet"
	"arvanwallet/services/wallet/arvanWallet"
)

type Services struct {
	Consumer consumer.Consumer
	Wallet   wallet.Wallet
}

func NewServices(repository *repositories.Repository) *Services {
	return &Services{
		Consumer: redis.NewRedisConsumer(repository),
		Wallet:   arvanWallet.NewR1Wallet(repository),
	}
}
