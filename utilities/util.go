package utilities

import (
	"os"
	"encoding/json"
)

func GetConfiguration() (Configuration ,error){

	configurationDB := Configuration{}

	file, err := os.Open("./configure.json")

	if err != nil {
		return configurationDB,err
	}
	defer file.Close()

	jsonDecoder := json.NewDecoder(file)

	err = jsonDecoder.Decode(&configurationDB)

	if err != nil {
		return  configurationDB, err
	}

	return configurationDB, nil

}
