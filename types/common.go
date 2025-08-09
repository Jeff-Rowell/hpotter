package types

type Service struct {
	ImageName     string `json:"image_name"`
	ListenPort    int16  `json:"listen_port"`
	ListenProto   string `json:"listen_proto"`
	RequestSave   bool   `json:"request_save"`
	ResponseSave  bool   `json:"response_save"`
	SocketTimeout int32  `json:"socket_timeout"`
	Tls           bool   `json:"tls"`
}
