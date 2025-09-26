package configuration

import (
	"github.com/spf13/viper"
)

type Conf struct {
	DatabaseURL         string `mapstructure:"MONGODB_URL"`
	DatabaseName        string `mapstructure:"MONGODB_DB"`
	BatchInsertInterval string `mapstructure:"BATCH_INSERT_INTERVAL"`
	AuctionInterval     string `mapstructure:"AUCTION_INTERVAL"`
	MaxBatchSize        int    `mapstructure:"MAX_BATCH_SIZE"`
}

func LoadConfig() (*Conf, error) {
	viper.SetDefault("MONGODB_URL", "mongodb://admin:admin@localhost:27017/auctions?authSource=admin")
	viper.SetDefault("MONGODB_DB", "auctions")
	viper.SetDefault("BATCH_INSERT_INTERVAL", "20s")
	viper.SetDefault("AUCTION_INTERVAL", "20s")
	viper.SetDefault("MAX_BATCH_SIZE", 4)

	viper.AutomaticEnv()

	var cfg Conf
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
