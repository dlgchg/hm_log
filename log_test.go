package hm_log

import (
	"testing"
	"time"
)

func TestLog(t *testing.T) {
	config := make(map[string]string, 8)
	config["log_path"] = "."
	config["log_name"] = "server"
	config["log_chan_size"] = "50000"
	err := InitLog("file", config)
	if err != nil {
		return
	}

	for {
		log.Warn("%s", "123")
		time.Sleep(time.Second)
	}
}
