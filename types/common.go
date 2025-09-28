package types

type EnvVar struct {
	Key   string `yaml:"key"`
	Value string `yaml:"value"`
}

type Service struct {
	ListenAddress      string   `yaml:"listen_address"`
	ListenPort         int      `yaml:"listen_port"`
	ListenProto        string   `yaml:"listen_proto"`
	RequestSave        bool     `yaml:"request_save"`
	EnvVars            []EnvVar `yaml:"envvars,omitempty"`
	CollectCredentials bool     `yaml:"collect_credentials"`
	UseTLS             bool     `yaml:"use_tls,omitempty"`
	CertificatePath    string   `yaml:"certificate_path,omitempty"`
	KeyPath            string   `yaml:"key_path,omitempty"`
	GenerateCerts      bool     `yaml:"generate_certs,omitempty"`
	CommandLimit       int      `yaml:"command_limit,omitempty"`
}

type DBConfig struct {
	DBType   string `yaml:"db_type"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

type IPInfo struct {
	Latitude  float32 `json:"lat"`
	Longitude float32 `json:"lon"`
}
