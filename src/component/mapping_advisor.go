package component

import (
	"encoding/json"
	"log"
	"strconv"
	"strings"
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
	ids := []string{}
	for i, a := range value {
		optional := "?"
		if i > 0 {
			optional = "&"
		}
		queryFilter := strings.Join([]string{optional, "filters[id][$in][", strconv.Itoa(i), "]=", strconv.Itoa(int(a.Id))}, "")
		ids = append(ids, queryFilter)
	}
	advisorForm := advisor.ResponseAdvisors{}
	result, err := srv.Service.GetAdvisorInfomation(strings.Join(ids, ""))
	if err != nil {
		log.Panic(err)
	}
	err = json.Unmarshal([]byte(result), &advisorForm)
	if err != nil {
		log.Panic(err)
	}
	return advisorForm.Data, nil
}
