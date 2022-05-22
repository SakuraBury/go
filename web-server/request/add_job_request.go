package request

import "common/models/endpoint"

type AddJobRequest struct {
	Endpoint    *endpoint.EndpointModel
	Target      string
	Params      string
	ProcessorId string
}
