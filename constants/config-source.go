package constants

// ConfigSource config source
type ConfigSource string

const (
	FILE ConfigSource = "file" // the content of the configuration comes from the local file on disk
	//ETCD      ConfigSource = "etcd"
	//ETCD3     ConfigSource = "etcd3"
	//CONSUL    ConfigSource = "consul"
	//FIRESTORE ConfigSource = "firestore"
	NACOS ConfigSource = "nacos" // the content of the configuration comes from the Nacos service
)
