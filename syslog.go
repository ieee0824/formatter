package formatter

import (
	"encoding/json"
	"fmt"

	"github.com/sirupsen/logrus"
)

type SysLogFormatter struct {
	TimestampFormat string
}

func prefixFieldClashes(data logrus.Fields) {
	_, ok := data["msg"]
	if ok {
		data["fields.msg"] = data["msg"]
	}
}

func (f *SysLogFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	data := make(logrus.Fields, len(entry.Data)+3)
	for k, v := range entry.Data {
		switch v := v.(type) {
		case error:
			data[k] = v.Error()
		default:
			data[k] = v
		}
	}
	prefixFieldClashes(data)
	time := entry.Time
	t := fmt.Sprintf("%s %d %02d:%02d:%02d", time.Month().String()[0:3], time.Day(), time.Hour(), time.Minute(), time.Second())
	data["time"] = t
	data["msg"] = entry.Message
	data["level"] = entry.Level.String()
	serialized, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("Failed to marshal fields to JSON, %v", err)
	}
	return []byte(t + " " + string(serialized) + "\n"), nil
}
