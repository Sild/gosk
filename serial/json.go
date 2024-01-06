package serial

import (
	"encoding/json"
	"os"

	"github.com/sild/gosk/log"
)

func JsonSToObj[T interface{}](data string) T {
	return JsonBToObj[T]([]byte(data))
}

func JsonFToObj[T interface{}](filePath string) T {
	data, err := os.ReadFile(filePath)
	if err != nil {
		log.Error("JsonFToObj: %v", err)
		data = make([]byte, 0)
	}
	return JsonBToObj[T](data)
}

func JsonBToObj[T interface{}](data []byte) T {
	var dst T
	if err := json.Unmarshal(data, &dst); err != nil {
		log.Error("JsonBToObj: %v", err)
	}
	return dst
}

func ObjToJsonS(data interface{}) string {
	b, err := json.Marshal(data)
	if err != nil {
		log.Error("ObjToJsonS: %v", err)
	}
	return string(b)
}

func ObjToJsonF(data interface{}, filePath string) {
	str := ObjToJsonS(data)
	if err := os.WriteFile(filePath, []byte(str), 0644); err != nil {
		log.Error("ObjToJsonF: %v", err)
	}
}
