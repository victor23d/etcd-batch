package utils

import (
	"github.com/sirupsen/logrus"
)

func FlatMap(m map[string]interface{}, fp map[string]interface{}, sep string, prefix string, log *logrus.Logger) {
	for k, v := range m {
		switch vv := v.(type) {
		case bool:
			fp[prefix+k] = v
		case string:
			fp[prefix+k] = v
			log.Println(fp)
		// case int:
		// 	fp[prefix+k] = v
		// 	log.Println(fp)
		// 	log.Println("int")
		case float64:
			fp[prefix+k] = v
			log.Println(fp)
		case map[string]interface{}:
			FlatMap(vv, fp, sep, prefix+k+sep, log)
		default:
			log.Fatalf(k, "is default", vv)
		}
	}
}
