package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type RemoteAttributes struct {
	Name        string
	Description string
	Label       string
}
type RemoteAttributesWithNumber struct {
	Name     string
	Total    int64
	Page     int8
	IgnoreMe int8
}
type RemoteResponse struct {
	Id         string
	Attributes map[string]interface{}
}

func (v RemoteResponse) AttributesAsInterface() interface{} {
	return v.Attributes
}

func GetRemoteResponse(mocked string) *RemoteResponse {
	var responseData RemoteResponse
	err := json.Unmarshal([]byte(mocked), &responseData)
	if err != nil {
		fmt.Println("Error reading json. Exiting.")
		panic(err)
	}
	return &responseData
}

func main() {
	remoteString := `{"id":"1", "attributes": {"name": "response 1", "description": "Longer text saying more details", "label": "Rensponse"}}`
	parsedOne := GetRemoteResponse(remoteString)
	fmt.Println(fmt.Sprintf("%v", parsedOne))
	fmt.Printf("Read type %v \n", reflect.TypeOf(parsedOne))
	fmt.Printf("Read attributes type %v \n", reflect.TypeOf(parsedOne.Attributes))

	// this panics
	converted := parsedOne.AttributesAsInterface().(RemoteAttributes)

	fmt.Printf("Parsed attributes %v \n", converted)
}
