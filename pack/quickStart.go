package main

import "go.uber.org/zap"

func main() {

	GatewayLogger.Debug("Hi Gateway Im Debug", zap.String("cs", "string"))
	GatewayLogger.Info("Hi Gateway  Im Info   ccccc")
}
