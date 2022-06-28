package component

import (
	"encoding/json"
	"log"
	"unisun/api/course-listener/src/model/advisor"
	"unisun/api/course-listener/src/ports"
)

type ServiceConsumerAdap struct {
	Service ports.ConsumerService
}

func NewServiceConsumerAdap(service ports.ConsumerService) *ServiceConsumerAdap {
	return &ServiceConsumerAdap{
		Service: service,
	}
}

func (srv *ServiceConsumerAdap) MappingAdvisor(value []advisor.AdvisorData) ([]advisor.AdvisorData, error) {
	resultAdvisors := []advisor.AdvisorData{}
	for _, a := range value {
		advisorForm := advisor.ResponseAdvisor{}
		result, err := srv.Service.GetAdvisorInfomation(a.Id)
		if err != nil {
			log.Panic(err)
		}
		err = json.Unmarshal([]byte(result), &advisorForm)
		if err != nil {
			log.Panic(err)
		}
		resultAdvisors = append(resultAdvisors, advisorForm.Data)
	}
	return resultAdvisors, nil
}
