package types

type EnvVar struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type Service struct {
	ImageName     string   `json:"image_name"`
	ListenAddress string   `json:"listen_address"`
	ListenPort    int      `json:"listen_port"`
	ListenProto   string   `json:"listen_proto"`
	RequestSave   bool     `json:"request_save"`
	ResponseSave  bool     `json:"response_save"`
	SocketTimeout int32    `json:"socket_timeout"`
	Tls           bool     `json:"tls"`
	EnvVars       []EnvVar `json:"envvars,omitempty"`
}

type DBConfig struct {
	DBType   string `json:"db_type"`
	User     string `json:"user"`
	Password string `json:"password"`
}
