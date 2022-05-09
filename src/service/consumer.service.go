package service

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"unisun/api/course-listenner/src/constants"
	"unisun/api/course-listenner/src/model"
	"unisun/api/course-listenner/src/utils"
)

func GetInformationFormStrapi(payloadRequest model.ServiceIncomeRequest) model.ServiceIncomeResponse {
	var serviceIncomeResponse = model.ServiceIncomeResponse{}
	url := os.Getenv(constants.HOST_STRAPI_SERVICE) + os.Getenv(constants.PATH_STRAPI_INFORMATION_GATEWAY)
	payload, err := json.Marshal(payloadRequest)
	if err != nil {
		log.Println("Change json to byte", err.Error())
		serviceIncomeResponse.Error = err.Error()
		return serviceIncomeResponse
	} else {
		err = nil
	}
	response := utils.HTTPRequest(url, constants.POST, payload)
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println("Read response from request error.", err.Error())
		serviceIncomeResponse.Error = err.Error()
		return serviceIncomeResponse
	} else {
		err = nil
		defer response.Body.Close()
	}
	err = json.Unmarshal([]byte(body), &serviceIncomeResponse)
	if err != nil {
		log.Println("Change byte to json response", err.Error())
		serviceIncomeResponse.Error = err.Error()
		return serviceIncomeResponse
	} else {
		err = nil

	}
	if serviceIncomeResponse.Error != "" {
		log.Println(serviceIncomeResponse.Error)
	}
	return serviceIncomeResponse
}
