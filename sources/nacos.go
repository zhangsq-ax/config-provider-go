package sources

import (
	"github.com/spf13/viper"
	_ "github.com/spf13/viper/remote"
	nacos "github.com/yoyofxteam/nacos-viper-remote"
	"github.com/zhangsq-ax/config-provider-go/options"
)

func InitConfigProviderInstanceByNacos(opts *options.ConfigProviderOptions) (*viper.Viper, error) {
	nacosOpts := opts.NacosOptions
	var err error
	if nacosOpts == nil {
		nacosOpts, err = options.GetNacosOptionsFromEnv()
		if err != nil {
			return nil, err
		}
	}
	opt := &nacos.Option{
		Url:         nacosOpts.Host,
		Port:        uint64(nacosOpts.Port),
		NamespaceId: nacosOpts.NamespaceId,
		GroupName:   nacosOpts.Group,
		Config: nacos.Config{
			DataId: nacosOpts.DataId,
		},
	}
	if nacosOpts.Username != "" && nacosOpts.Password != "" {
		opt.Auth = &nacos.Auth{
			Enable:   true,
			User:     nacosOpts.Username,
			Password: nacosOpts.Password,
		}
	}
	nacos.SetOptions(opt)
	instance := viper.New()
	err = instance.AddRemoteProvider(string(opts.ConfigSource), nacosOpts.Host, "")
	if err != nil {
		return nil, err
	}
	instance.SetConfigType(string(opts.ConfigFormat))
	err = instance.ReadRemoteConfig()
	if err != nil {
		return nil, err
	}
	if opts.WatchConfig {
		remoteProvider := nacos.NewRemoteProvider(string(opts.ConfigFormat))
		resChan := remoteProvider.WatchRemoteConfigOnChannel(instance)
		go func() {
			for {
				<-resChan
				if opts.OnConfigChange != nil {
					opts.OnConfigChange()
				}
			}
		}()
	}
	return instance, nil
}
