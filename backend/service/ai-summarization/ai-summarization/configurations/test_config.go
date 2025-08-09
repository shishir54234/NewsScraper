package configurations

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestInitConfig_LoadsDevelopmentConfig(t *testing.T) {
	// Arrange: point CONFIG_PATH to test fixtures folder
	wd, err := os.Getwd()
	require.NoError(t, err)

	// Assuming config.development.json is in the same folder as this test
	configDir := filepath.Join(wd)
	os.Setenv("APP_ENV", "development")
	os.Setenv("CONFIG_PATH", configDir)

	// Act
	cfg, llmCfg, grpcCfg,_, loggerCfg, jaegerCfg, err := InitConfig()

	// Assert
	require.NoError(t, err)
	require.NotNil(t, cfg)
	require.NotNil(t, llmCfg)
	require.NotNil(t, grpcCfg)
	require.NotNil(t, loggerCfg)
	require.NotNil(t, jaegerCfg)

	// Check main config values
	require.Equal(t, "product_service", cfg.ServiceName)
	require.Equal(t, "product_service", jaegerCfg.ServiceName)
	require.Equal(t, "product_tracer", jaegerCfg.TracerName)

	// Check logger
	require.Equal(t, "debug", loggerCfg.LogLevel)

	// Check RabbitMQ
	require.Equal(t, "guest", cfg.Rabbitmq.User)
	require.Equal(t, "localhost", cfg.Rabbitmq.Host)
	require.Equal(t, 5672, cfg.Rabbitmq.Port)

	// Check gRPC
	require.Equal(t, ":6600", grpcCfg.Port)
	require.True(t, grpcCfg.Development)

	// Check LLM
	require.Equal(t, "AIzaSyAd75EAAm70NXIhrUfp8nJNIZKpSZOBX_U", llmCfg.ApiKey)
	require.Equal(t, "https://generativelanguage.googleapis.com/v1beta/models/gemini-2.0-flash:generateContent", llmCfg.BaseUrl)
}

func TestFilenameAndDirname(t *testing.T) {
	// filename()
	f, err := filename()
	require.NoError(t, err)
	require.Contains(t, f, "config_test.go") // should point to this test file

	// dirname()
	d, err := dirname()
	require.NoError(t, err)
	require.NotEmpty(t, d)
	require.DirExists(t, d)
}

func TestGetMicroserviceName(t *testing.T) {
	got := GetMicroserviceName("my_service")
	require.Equal(t, "MY_SERVICE", got)
}
