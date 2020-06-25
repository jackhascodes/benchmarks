package logging

import (
	"bytes"
	"encoding/json"
	"github.com/jackhascodes/kit/log"
	"github.com/sirupsen/logrus"
	stdLog "log"
	"testing"
	"time"
)

func BenchmarkLog_InfoJson(b *testing.B) {
	var buf bytes.Buffer
	logger := log.InitLog(log.WithOutput(&buf), log.WithFormatter(&log.JsonFormatter{}))
	for i := 0; i < b.N; i++ {
		logger.Info("foo", "bar")
		buf.Truncate(0)
	}
}

func BenchmarkLogrus_InfoJson(b *testing.B) {
	var buf bytes.Buffer
	logger := logrus.New()
	logger.SetOutput(&buf)
	logger.SetFormatter(&logrus.JSONFormatter{})
	for i := 0; i < b.N; i++ {
		logger.Info("foo", "bar")
		buf.Truncate(0)
	}
}

func BenchmarkStandardLog_PrintJson(b *testing.B) {
	// We should bench standard logging with some kind of equivalent functionality.
	prepLog := func(level, msg string) string {
		out := map[string]string{
			"time":  time.Now().Format(time.RFC3339),
			"level": level,
			"msg":   msg,
		}
		by, _ := json.Marshal(out)
		return string(by)
	}
	var buf bytes.Buffer
	stdLog.SetOutput(&buf)
	for i := 0; i < b.N; i++ {
		stdLog.Print(prepLog("Info", "foo"))
		buf.Truncate(0)
	}
}



func BenchmarkLog_InfoJson_withfields(b *testing.B) {
	var buf bytes.Buffer
	logger := log.InitLog(
		log.WithOutput(&buf),
		log.WithFormatter(&log.JsonFormatter{}),
		log.WithFields(log.KV{"foo","bar"}, log.KV{"baz","qux"}))
	for i := 0; i < b.N; i++ {
		logger.Info("foo", "bar")
		buf.Truncate(0)
	}
}

func BenchmarkLogrus_InfoJson_withfields(b *testing.B) {
	var buf bytes.Buffer
	logger := logrus.New()
	logger.SetOutput(&buf)
	logger.SetFormatter(&logrus.JSONFormatter{})
	for i := 0; i < b.N; i++ {
		// unless you wrap logrus with something that takes an initial field set and applies it each time you call
		// a log method, you need to define the fields each time.
		logger.WithFields(logrus.Fields{"foo":"bar","baz":"qux"}).Info("quux")
		buf.Truncate(0)
	}
}