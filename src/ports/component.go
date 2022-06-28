package ports

import "unisun/api/course-listener/src/model/advisor"

type MappingAdvisorComponentPort interface {
	MappingAdvisor(value []advisor.AdvisorData) ([]advisor.AdvisorData, error)
}
