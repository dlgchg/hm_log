package hm_log

import (
	"testing"
)

func TestLog(t *testing.T) {
	config := make(map[string]string, 8)
	config["log_path"] = "."
	config["log_name"] = "log"
	config["log_level"] = "debug"
	err := InitLog("console", config)
	if err != nil {
		return
	}

	log.Debug("%s", "123")
}
