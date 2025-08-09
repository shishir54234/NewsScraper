package configurations

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"github.com/shishir54234/NewsScraper/backend/pkg/config"
	"github.com/shishir54234/NewsScraper/backend/pkg/grpc"
	"github.com/shishir54234/NewsScraper/backend/pkg/logger"
	"github.com/shishir54234/NewsScraper/backend/pkg/otel"
	"github.com/shishir54234/NewsScraper/backend/pkg/rabbitmq"
	"github.com/spf13/viper"
)


var configPath string 
type Config struct {
	ServiceName  string                        `mapstructure:"serviceName"`
	Logger       *logger.LoggerConfig          `mapstructure:"logger"`
	Rabbitmq     *rabbitmq.RabbitMQConfig      `mapstructure:"rabbitmq"`
	// Echo         *echoserver.EchoConfig        `mapstructure:"echo"`
	Grpc         *grpc.GrpcConfig              `mapstructure:"grpc"`
	Llmconfig  	*config.LlmConfig                  `mapstructure:"llmConfig"`
	Jaeger       *otel.JaegerConfig            `mapstructure:"jaeger"`
}
func Init(){
	flag.StringVar(&configPath, "config", "", "generating description microservices")
}

func InitConfig()(*Config, *config.LlmConfig, *grpc.GrpcConfig,*rabbitmq.RabbitMQConfig, 
*logger.LoggerConfig, 
*otel.JaegerConfig,
error){
	env:=os.Getenv("APP_ENV")
	if env == ""{
		env="development"
	}
	if configPath==""{
		configPathFromEnv := os.Getenv("CONFIG_PATH")
		if configPathFromEnv!="" {
			configPath=configPathFromEnv
		}else {
			d,err:=dirname()
			if err!=nil{
				return nil ,nil, nil,nil, nil,nil,err
			}
			configPath=d
		}
	}
	cfg := &Config{}

	viper.SetConfigName(fmt.Sprintf("config.%s", env))
	viper.AddConfigPath(configPath)
	viper.SetConfigType("json")

	if err := viper.ReadInConfig(); err != nil {
		return nil ,nil, nil, nil,nil,nil,err
	}
	if err := viper.Unmarshal(cfg); err != nil {
		return nil ,nil, nil, nil,nil,nil,err
	}
	fmt.Println("API_KEY", cfg.Llmconfig.ApiKey)
	fmt.Println("BASE_URL", cfg.Llmconfig.BaseUrl)
	return cfg, cfg.Llmconfig, cfg.Grpc,cfg.Rabbitmq, cfg.Logger, cfg.Jaeger, nil 
}

func GetMicroserviceName(serviceName string) string {
	return fmt.Sprintf("%s", strings.ToUpper(serviceName))
}

func filename() (string, error) {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		return "", errors.New("unable to get the current filename")
	}
	return filename, nil
}

func dirname() (string, error) {
	filename, err := filename()
	if err != nil {
		return "", err
	}
	return filepath.Dir(filename), nil
}

