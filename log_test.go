package hm_log

import (
	"testing"
)

func TestLog(t *testing.T) {
	config := make(map[string]string, 8)
	config["log_path"] = "."
	config["log_name"] = "server"
	config["log_chan_size"] = "50000" //chan size
	config["log_split_type"] = "size"
	config["log_split_size"] = "104857600" // 100MB
	err := InitLog("file", config)
	if err != nil {
		return
	}

	for {
		log.Warn("%s", "123")
	}
}
