package utilities

import (
	"os"
	"encoding/json"
	"fmt"
)

func GetConfiguration() (Configuration ,error){

	configurationDB := Configuration{}

	file, err := os.Open("./configure.json")

	if err != nil {
		fmt.Println("nose 1")
		return configurationDB,err
	}
	defer file.Close()

	jsonDecoder := json.NewDecoder(file)

	err = jsonDecoder.Decode(&configurationDB)

	if err != nil {
		fmt.Println("nose 2")
		return  configurationDB, err
	}

	return configurationDB, nil

}
