package service

import (
	"encoding/json"
	"io/ioutil"
	"strings"
	"unisun/api/course-listener/src/constants"
	"unisun/api/course-listener/src/model"
	"unisun/api/course-listener/src/ports"

	"github.com/spf13/viper"
)

type ServiceConsumerAdapter struct {
	HTTPRequest ports.HTTPRequest
}

func NewServiceConsumerAdapter(httpRequest ports.HTTPRequest) *ServiceConsumerAdapter {
	return &ServiceConsumerAdapter{
		HTTPRequest: httpRequest,
	}
}

func (srv *ServiceConsumerAdapter) GetInformationFormStrapi(payloadRequest model.ServiceIncomeRequest) (string, error) {
	var serviceIncomeResponse = model.ServiceIncomeResponse{}
	url := strings.Join([]string{viper.GetString("endpoint.strapi-information-gateway.host"), viper.GetString("endpoint.strapi-information-gateway.path")}, "")
	payload, err := json.Marshal(payloadRequest)
	if err != nil {
		return "", err
	}
	response, err := srv.HTTPRequest.HTTPRequest(url, constants.POST, payload)
	if err != nil {
		return "", err
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()
	err = json.Unmarshal([]byte(body), &serviceIncomeResponse)
	if err != nil {
		return "", err
	}
	return serviceIncomeResponse.Payload, nil
}

func (srv *ServiceConsumerAdapter) GetAdvisorInfomation(ids string) (string, error) {
	url := strings.Join([]string{viper.GetString("endpoint.advisor.host"), viper.GetString("endpoint.advisor.path"), ids}, "")
	response, err := srv.HTTPRequest.HTTPRequest(url, constants.GET, nil)
	if err != nil {
		return "", err
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()
	return string(body), nil
}
