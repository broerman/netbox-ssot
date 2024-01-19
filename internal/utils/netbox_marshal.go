package utils

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

// This function takes an object pointer, and returns a json body,
// that can be used to create that object in netbox API.
// This is essential because default marshal of the object
// isn't compatible with netbox API when attributes have nested
// objects.
func NetboxJsonMarshal(obj interface{}) ([]byte, error) {
	objMap, err := StructToNetboxJsonMap(obj)
	if err != nil {
		return nil, fmt.Errorf("error converting object to json map: %s", err)
	}
	return json.Marshal(objMap)
}

// Function that converts an object to a map[string]interface{}
// which can be used to create a json body for netbox API, especially
// for POST requests.
func StructToNetboxJsonMap(obj interface{}) (map[string]interface{}, error) {
	v := reflect.ValueOf(obj)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	netboxJsonMap := make(map[string]interface{})
	for i := 0; i < v.NumField(); i++ {
		fieldValue := v.Field(i)
		fieldType := v.Type().Field(i)
		jsonTag := fieldType.Tag.Get("json")
		jsonTag = strings.Split(jsonTag, ",")[0]

		if fieldType.Name == "Id" {
			continue
		}

		// Special case when object inherits from NetboxObject
		if fieldType.Name == "NetboxObject" {
			diffMap, err := StructToNetboxJsonMap(fieldValue.Interface())
			if err != nil {
				return nil, fmt.Errorf("error processing ObjToJsonMap when processing NetboxObject %s", err)
			}
			for k, v := range diffMap {
				netboxJsonMap[k] = v
			}
			continue
		}

		// If field is a pointer, we need to get the element it points to
		if fieldValue.Kind() == reflect.Ptr {
			fieldValue = fieldValue.Elem()
		}

		// If a field is empty we skip it
		if !fieldValue.IsValid() || fieldValue.IsZero() {
			continue
		}

		switch fieldValue.Kind() {
		case reflect.Slice:
			if fieldValue.Len() == 0 {
				continue
			}
			sliceItems := make([]interface{}, 0)
			for j := 0; j < fieldValue.Len(); j++ {
				attribute := fieldValue.Index(j)
				if attribute.Kind() == reflect.Ptr {
					attribute = attribute.Elem()
				}
				if attribute.Kind() == reflect.Struct {
					id := attribute.FieldByName("Id")
					if id.IsValid() && id.Int() != 0 {
						sliceItems = append(sliceItems, id.Int())
					} else {
						sliceItems = append(sliceItems, attribute.Interface())
					}
				} else {
					sliceItems = append(sliceItems, attribute.Interface())
				}
			}
			netboxJsonMap[jsonTag] = sliceItems
		case reflect.Struct:
			if isChoiceEmbedded(fieldValue) {
				choiceValue := fieldValue.FieldByName("Value")
				if choiceValue.IsValid() {
					netboxJsonMap[jsonTag] = choiceValue.Interface()
				}
			} else {
				id := fieldValue.FieldByName("Id")
				if id.IsValid() {
					netboxJsonMap[jsonTag] = id.Int()
				} else {
					netboxJsonMap[jsonTag] = fieldValue.Interface()
				}
			}
		default:
			netboxJsonMap[jsonTag] = fieldValue.Interface()
		}
	}
	return netboxJsonMap, nil
}