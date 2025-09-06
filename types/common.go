package types

type EnvVar struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type Service struct {
	ImageName             string   `json:"image_name"`
	ListenAddress         string   `json:"listen_address"`
	ListenPort            int      `json:"listen_port"`
	ListenProto           string   `json:"listen_proto"`
	RequestSave           bool     `json:"request_save"`
	EnvVars               []EnvVar `json:"envvars,omitempty"`
	ServiceName           string   `json:"service_name"`
	CollectCredentials    bool     `json:"collect_credentials"`
	CredentialLogPattern  string   `json:"credential_log_pattern,omitempty"`
	SessionDataLogPattern string   `json:"session_data_log_pattern,omitempty"`
}

type DBConfig struct {
	DBType      string `json:"db_type"`
	User        string `json:"user"`
	Password    string `json:"password"`
	DBImageName string `json:"image_name"`
}

type IPInfo struct {
	Latitude  float32 `json:"lat"`
	Longitude float32 `json:"lon"`
}
