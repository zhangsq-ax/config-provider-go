package config_provider

import (
	"fmt"
	"github.com/spf13/viper"
	_ "github.com/spf13/viper/remote"
	"github.com/zhangsq-ax/config-provider-go/constants"
	"github.com/zhangsq-ax/config-provider-go/options"
	"github.com/zhangsq-ax/config-provider-go/sources"
)

// ConfigProvider Config provider
type ConfigProvider struct {
	options       *options.ConfigProviderOptions // config provider options
	viperInstance *viper.Viper                   // config provider instance, used to get config
}

// init Initialize the config provider
func (provider *ConfigProvider) init() error {
	// different config source have different initialization methods
	switch provider.options.ConfigSource {
	case constants.FILE:
		instance, err := sources.InitConfigProviderInstanceByFile(provider.options)
		if err != nil {
			return err
		}
		provider.viperInstance = instance
	case constants.NACOS:
		instance, err := sources.InitConfigProviderInstanceByNacos(provider.options)
		if err != nil {
			return err
		}
		provider.viperInstance = instance
	default:
		return fmt.Errorf("unknown config source")
	}
	return nil
}

// OnConfigChanged Set config changed callback
func (provider *ConfigProvider) OnConfigChanged(configChangedCallback func()) {
	provider.options.OnConfigChange = configChangedCallback
}

// Instance Get instance of the config provider
func (provider *ConfigProvider) Instance() *viper.Viper {
	return provider.viperInstance
}

// NewConfigProvider Instantiate the config provider
func NewConfigProvider(opts *options.ConfigProviderOptions) (*ConfigProvider, error) {
	err := opts.Validate()
	if err != nil {
		return nil, err
	}

	provider := &ConfigProvider{
		options: opts,
	}
	err = provider.init()
	if err != nil {
		return nil, err
	}

	return provider, nil
}
