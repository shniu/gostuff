package serializer

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func jsonDemo() {
	var obj interface{}
	err := json.Unmarshal([]byte("{\"abc\":123}"), &obj)

	if err != nil {
		return
	}

	objMap, ok := obj.(map[string]interface{})
	if ok {
		for k, v := range objMap {
			switch value := v.(type) {
			case string:
				fmt.Printf("type of %s is string, value is %v\n", k, value)
			case float64:
				fmt.Printf("type of %s is float64, value is %v\n", k, value)
			case interface{}:
				fmt.Printf("type of %s is interface{}, value is %v, typeof: %v\n",
					k, value, reflect.TypeOf(value))
			default:
				fmt.Printf("type of %s is wrong, value is %v\n", k, value)
			}
		}
	}
}
