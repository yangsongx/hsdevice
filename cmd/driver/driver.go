package driver

import (
	"github.com/edgexfoundry/go-mod-core-contracts/v2/clients/logger"
)

type HansongDevice struct {
	lc                 logger.LoggingClient
	temperature        int
//	serviceConfig     *config.ServiceConfig
}