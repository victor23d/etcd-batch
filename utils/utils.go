package utils

import (
	"errors"
	"github.com/sirupsen/logrus"
)

func FlatMap(m map[string]interface{}, fp map[string]interface{}, sep string, prefix string, log *logrus.Logger) error {
	// func prefix2 (prefix string, key string) string{
	// 	return prefix + key
	// }
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
			return errors.New("Unknown type: " + k + vv.(string))
		}
	}
	return nil
}
