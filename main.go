package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func printJS(entity interface{}, space string) {
	val := reflect.ValueOf(entity)// getting the value of entity

	switch val.Kind() { // switch on kind of value that is present in the entity
	case reflect.Map: // case map
		fmt.Println(space + "Type: Map")
		for _, key := range val.MapKeys() {
			fmt.Println(space + "Key: " + key.String())
			printJS(val.MapIndex(key).Interface(), space+"  ") // recurssive call to each element of the map
		} // case slice
	case reflect.Slice:
		fmt.Println(space + "Type: Slice")
		for i := 0; i < val.Len(); i++ {
			printJS(val.Index(i).Interface(), space+"  ") // recurssive call to all the elemnts of the slice
		}
	default:
		fmt.Println(space + "Type: " + val.Type().String())
		fmt.Println(space + "Value: " + fmt.Sprintf("%v", val.Interface()))	// other than that we don't have to do recurssion and hence simply printed
	}
}

func main() {
	data := `{
		"name" : "Tolexo Online Pvt. Ltd",
		"age_in_years" : 8.5,
		"origin" : "Noida",
		"head_office" : "Noida, Uttar Pradesh",
		"address" : [
			{
				"street" : "91 Springboard",
				"landmark" : "Axis Bank",
				"city" : "Noida",
				"pincode" : 201301,
				"state" : "Uttar Pradesh"
			},
			{
				"street" : "91 Springboard",
				"landmark" : "Axis Bank",
				"city" : "Noida",
				"pincode" : 201301,
				"state" : "Uttar Pradesh"
			}
		],
		"sponsers" : {
			"name" : "One"
		},
		"revenue" : "19.8 million$",
		"no_of_employee" : 630,
		"str_text" : ["one","two"],
		"int_text" : [1,3,4]
	}`

	var entity interface{} // variable to store entity
	err := json.Unmarshal([]byte(data), &entity) // unmarshall from string to map type
	if err != nil {
		panic(err) //exit of error exists
		return
	}

	printJS(entity, "")// function call to print our data, with proper indentation
}
