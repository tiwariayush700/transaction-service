package server

import (
	"context"
	"transaction-service/pkg/config"
)

func Start(ctx context.Context) {

	_ = config.GetAppConfiguration()
	
}
