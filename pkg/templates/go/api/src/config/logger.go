package config

import (
	"github.com/sirupsen/logrus"
)

func InitLogrus(env string) {
	if env == "local" {
		logrus.SetFormatter(
			&logrus.JSONFormatter{
				PrettyPrint: true,
				FieldMap: logrus.FieldMap{
					"time":  "timestamp",
					"level": "level",
					"msg":   "message",
				},
			},
		)
		logrus.SetReportCaller(true)
		logrus.SetLevel(logrus.DebugLevel)
		return
	}
	logrus.SetFormatter(
		&logrus.JSONFormatter{
			FieldMap: logrus.FieldMap{
				"time":  "timestamp",
				"level": "level",
				"msg":   "message",
			},
		},
	)
	logrus.SetLevel(logrus.InfoLevel)

}
