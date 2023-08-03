package options

import (
	"fmt"
	"github.com/zhangsq-ax/config-provider-go/constants"
)

// ConfigProviderOptions config provider init options
type ConfigProviderOptions struct {
	ConfigFormat   constants.ConfigFormat // config content format
	ConfigSource   constants.ConfigSource // config source
	WatchConfig    bool                   // whether to automatically update when the configuration changed
	OnConfigChange func()                 // When WatchConfig is true, the callback will be called after the configuration changed
	ConfigFile     string                 // path to the config file if configSource is FILE
	NacosOptions   *NacosOptions          // nacos options if configSource is NACOS
}

func (opts *ConfigProviderOptions) Validate() error {
	if opts.WatchConfig && opts.OnConfigChange == nil {
		return fmt.Errorf("OnConfigChange callback is required when WatchConfig is true")
	}
	return nil
}
