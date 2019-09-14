package common

import (
	"bufio"
	json "github.com/json-iterator/go"
	"io/ioutil"
	"os"
	"strings"
	"log"
)


func ReadStringFromCommand() (string, error) {
	Prefix := os.Args[1]
	log.Printf("Prefix %s",Prefix)

	scanner := bufio.NewScanner(os.Stdin)
	var s strings.Builder
	for scanner.Scan() {
		s.WriteString(scanner.Text() + "\n")
	}
	if scanner.Err() != nil {
		return "", scanner.Err()
	}
	return s.String(), nil
}

func ReadJSONFromFile(filename string ) (map[string]interface{}, error) {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	log.Printf("Read JSONfile %s",string(b))
	var mf interface{}

	err = json.Unmarshal(b, &mf)
	if err != nil {
		return nil, err
	}

	m := mf.(map[string]interface{})

	return m, nil
}

//func SetLog(log *logrus.Logger) *logrus.Logger {
//	log.SetReportCaller(true)
//	log.Formatter = &logrus.TextFormatter{
//		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
//			filename := path.Base(f.File)
//			return fmt.Sprintf("%s()", f.Function), fmt.Sprintf("%s:%d", filename, f.Line)
//		},
//	}
//	return log
//}
