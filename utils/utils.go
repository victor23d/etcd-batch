package utils

import (
	"strconv"
	"strings"
)

// FlatMap Usage
// fp := make(map[string]interface{})
// utils.FlatMap(m, fp, "/", "" )
func FlatMap(m map[string]interface{}, fp map[string]interface{}, sep string, prefix string) {
	for k, v := range m {
		switch vv := v.(type) {
		case bool:
			fp[prefix+k] = v
		case string:
			fp[prefix+k] = v
		// case int:
		//	fp[prefix+k] = v
		case float64:
			fp[prefix+k] = v
			// log.Println(fp)
		case map[string]interface{}:
			FlatMap(vv, fp, sep, prefix+k+sep)
		default:
			log.Fatalf(k, "type known", vv)
		}
	}
}

// StringFlatedMap stringify FlatedMap
func StringFlatedMap(fp map[string]interface{}) map[string]string {
	sfp := make(map[string]string)
	for k, v := range fp {
		switch vv := v.(type) {
		case string:
			val := v.(string)
			sfp[k] = val
		case bool:
			val := strconv.FormatBool(v.(bool))
			sfp[k] = val
		case float64:
			val := strconv.FormatFloat(v.(float64), 'f', -1, 64)
			sfp[k] = val
		default:
			log.Fatalf(k, "type known", vv)
		}
		// TODO
	}
	return sfp
}

func TextSFP(fp map[string]string) strings.Builder {
	var sb strings.Builder
	sb.WriteString("\n")
	for k, v := range fp {
		sb.WriteString(k + "=" + v + "\n")

	}
	return sb
}
