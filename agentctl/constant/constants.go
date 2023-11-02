package constant

var (
	DATE_FORMAT string = "2006-01-02 15:04:05"

	DEFAULT_LOG_PATH  string = "/var/log/agentctl"
	DEFAULT_LOG_COUNT int    = 1
	DEFAULT_LOG_TIME  int    = 24

	DEFAULT_CERT_PATH string = "/opt/agentctl/cert"

	// common params
	CMD_PARAM_IP   string = "ip"
	CMD_PARAM_PORT string = "port"

	// fluentbit params
	FLUENTBIT_PARAM_CONFIG string = "config"

	// suricata params
	SURICATA_PARAM_UPLOAD string = "upload"
	SURICATA_PARAM_FILE   string = "file"
)
