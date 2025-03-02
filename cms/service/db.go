package service

import (
	"cms/config"
	. "cms/model"
	"encoding/json"
	"fmt"
	"os"
)

type FileBDService struct {
}

// Write
func (fs FileBDService) Write(data []Customer) {
	jsonBytes, _ := json.Marshal(data)
	os.WriteFile(config.CustomerJsonPath, jsonBytes, 0666)
}

func (fs FileBDService) Read() []Customer {
	var data []Customer
	//os operation
	jsonBytes, err := os.ReadFile(config.CustomerJsonPath)
	if err != nil {
		fmt.Println("err:", err)
	} else {
		//Deserialize and store
		json.Unmarshal(jsonBytes, &data)
		fmt.Println("data", data)
	}
	return data
}
