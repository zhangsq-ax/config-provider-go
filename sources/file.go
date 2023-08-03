package sources

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"github.com/zhangsq-ax/config-provider-go/options"
	"os"
)

// InitConfigProviderInstanceByFile Initialize an instance of config provider as file source
func InitConfigProviderInstanceByFile(opts *options.ConfigProviderOptions) (*viper.Viper, error) {
	instance := viper.New()
	configFile := opts.ConfigFile
	// If the config file path is not set then try to get it from the environment variable
	if configFile == "" {
		configFile = os.Getenv("CONFIG_FILE")
	}
	if configFile == "" {
		return nil, fmt.Errorf("no set CONFIG_FILE option")
	}
	instance.SetConfigFile(configFile)
	instance.SetConfigType(string(opts.ConfigFormat))
	err := instance.ReadInConfig()
	if err != nil {
		return nil, err
	}
	if opts.WatchConfig {
		instance.OnConfigChange(func(e fsnotify.Event) {
			if opts.OnConfigChange != nil {
				opts.OnConfigChange()
			}
		})
		instance.WatchConfig()
	}
	return instance, nil
}
