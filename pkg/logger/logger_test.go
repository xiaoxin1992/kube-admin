package logger

import "testing"

func TestLogger(t *testing.T) {
	err := InitLogger(DefaultOption())
	if err != nil {
		t.Error(err)
		return
	}
	log := GetLogger()
	log.S("test").Info("xxx")
}
