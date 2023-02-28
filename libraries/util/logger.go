package util

import "go.uber.org/zap"

func Logger() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	logger.Info("failed to fetch URL",
		// Structured context as strongly typed Field values.
		zap.String("url", "aaa"),
		zap.Int("attempt", 3),
	)
}
