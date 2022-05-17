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

func CallSumRateScoreFeedback(couseId string) model.ResultSum {
	url := os.Getenv(constants.HOST_FEEDBACK_PROCESSOR_API) + os.Getenv(constants.PATH_FEEDBACK_PROCESSOR_API) + couseId
	response := utils.HTTPRequest(url, constants.GET, nil)
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Panic("Read response from request error.", err.Error())
	} else {
		err = nil
		defer response.Body.Close()
	}
	var result = model.ResultSum{}
	err = json.Unmarshal([]byte(body), &result)
	if err != nil {
		log.Panic("Change byte to json response.", err.Error())
	} else {
		err = nil
	}
	return result
}
