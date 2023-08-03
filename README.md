# config-provider-go
Provide multiple configuration sources and format support for applications. Based on [https://github.com/spf13/viper](https://github.com/spf13/viper)

## Features

### Support Source

* Local config file
* Nacos

### Support Config Format

* JSON
* YAML
* Properties
* TOML
* INI

### Watch Config Changed

Support watch configuration changes

### Dependents

* [github.com/spf13/viper](https://github.com/spf13/viper)
* [github.com/yoyofxteam/nacos-viper-remote](https://github.com/yoyofxteam/nacos-viper-remote")

## Example

### From Local File
```go
package main

import (
	"fmt"
	config_provider "github.com/zhangsq-ax/config-provider-go"
	options "github.com/zhangsq-ax/config-provider-go/options"
	constants "github.com/zhangsq-ax/config-provider-go/constants"
)

func main() {
	provider, err := config_provider.NewConfigProvider(&options.ConfigProviderOptions{
		ConfigFormat: constants.YAML,
		ConfigSource: constants.FILE,
		ConfigFile:   "./config.yml",
		WatchConfig: true,
		OnConfigChange: func() {
			fmt.Println("config changed")
			// to do something
        },
    })
	if err != nil {
        panic(err)
	}
	
	fmt.Println(provider.Instance().AllKeys())
}
```

### From Nacos

```go
package main

import (
	"fmt"
	config_provider "github.com/zhangsq-ax/config-provider-go"
	options "github.com/zhangsq-ax/config-provider-go/options"
	constants "github.com/zhangsq-ax/config-provider-go/constants"
)

func main() {
	provider, err := config_provider.NewConfigProvider(&options.ConfigProviderOptions{
		ConfigFormat: constants.YAML,
		ConfigSource: constants.NACOS,
		WatchConfig: true,
		OnConfigChange: func() {
			fmt.Println("config changed")
			// to do something
        },
		NacosOptions: &options.NacosOptions{
			Host:        "localhost",
			Port:        8848,
			NamespaceId: "public",
			Group:       "DEFAULT_GROUP",
			DataId:      "...",
			Username:    "...",
			Password:    "...",
        },
    })
	if err != nil {
        panic(err)
	}
	
	fmt.Println(provider.Instance().AllKeys())
}
```

## Methods

### `NewConfigProvider(opts *options.ConfigProviderOptions)`

Instantiate a config provider

### `provider.Instance()`

Get instance of the config provider. The instance is a `Viper` object, configuration content reading methods reference [https://github.com/spf13/viper](https://github.com/spf13/viper)
