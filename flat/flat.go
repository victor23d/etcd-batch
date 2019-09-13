package flat

import (
	"log"
	"math"
	"strconv"
	"strings"
	// json "github.com/json-iterator/go"
)

// FlatMap Usage
// fp := make(map[string]interface{})
// flat.FlatMap(m, fp, "/", "" )
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
		case map[string]interface{}:
			FlatMap(vv, fp, sep, prefix+k+sep)
		default:
			log.Fatalf(k, "type known", vv)
		}
	}
}

// StringFlatedMap cast other type in json value to string FlatedMap
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

// TextSFP turn sfp to etcd get prefix format
func TextSFP(sfp map[string]string) strings.Builder {
	var tfp strings.Builder
	tfp.WriteString("\n")
	for k, v := range sfp {
		tfp.WriteString(k + "\n" + v + "\n")

	}
	return tfp
}

// TODO
func UnFlatMap(tfp string, sep string) map[string]interface{} {
	m := make(map[string]interface{})
	kvs := strings.Split(tfp, "\n")
	log.Println(kvs)
	for i := 0; i < len(kvs); i++ {
		// test if is a key
		if math.Mod(float64(i), 2) == 0 {
			kk := strings.Split(kvs[i], sep)

			// t := reflect.TypeOf(kk)
			if len(kk) == 1 {
			}
			// recursive

		}
	}

	return m
}
