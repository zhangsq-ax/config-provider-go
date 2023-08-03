package options

import (
	"fmt"
	"os"
	"strconv"
)

// NacosOptions The nacos service options
type NacosOptions struct {
	Host        string // nacos service host
	Port        int    // nacos service port
	Scheme      string // nacos service scheme
	ContextPath string // nacos service context path
	Username    string // nacos service username
	Password    string // nacos service password
	NamespaceId string // nacos service namespace id
	DataId      string // nacos service data id
	Group       string // nacos service group name
}

func (opts *NacosOptions) Validate() error {
	if opts.Host == "" {
		return fmt.Errorf("invalid NACOS_HOST option: %s", opts.Host)
	}
	if opts.Port == 0 {
		return fmt.Errorf("invalid NACOS_PORT option: %d", opts.Port)
	}
	if opts.NamespaceId == "" {
		return fmt.Errorf("invalid NACOS_NAMESPACE_ID option: %s", opts.NamespaceId)
	}
	if opts.DataId == "" {
		return fmt.Errorf("invalid NACOS_DATA_ID option: %s", opts.DataId)
	}
	if opts.Group == "" {
		return fmt.Errorf("invalid NACOS_GROUP option: %s", opts.Group)
	}
	return nil
}

// GetNacosOptionsFromEnv Get nacos service options from environment variables
func GetNacosOptionsFromEnv() (*NacosOptions, error) {
	host := os.Getenv("NACOS_HOST")
	port := os.Getenv("NACOS_PORT")
	scheme := os.Getenv("NACOS_SCHEME")
	contextPath := os.Getenv("NACOS_CONTEXT_PATH")
	username := os.Getenv("NACOS_USERNAME")
	password := os.Getenv("NACOS_PASSWORD")
	namespaceId := os.Getenv("NACOS_NAMESPACE_ID")
	dataId := os.Getenv("NACOS_DATA_ID")
	group := os.Getenv("NACOS_GROUP")

	portInt, err := strconv.Atoi(port)
	if err != nil {
		return nil, fmt.Errorf("no set NACOS_PORT option")
	}

	opts := &NacosOptions{
		Host:        host,
		Port:        portInt,
		Scheme:      scheme,
		ContextPath: contextPath,
		Username:    username,
		Password:    password,
		NamespaceId: namespaceId,
		DataId:      dataId,
		Group:       group,
	}

	err = opts.Validate()
	if err != nil {
		return nil, err
	}

	return opts, nil
}
