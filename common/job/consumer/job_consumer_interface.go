package consumer

import (
	"common/models/endpoint"
	"common/models/job"
)

type JobConsumerInterface interface {
	GetJob(endpoint endpoint.EndpointModel) *job.JobCmd
}
