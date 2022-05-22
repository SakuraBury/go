package common

import (
	"common/job/producer/impl"
)

var producer *impl.MockJobProducer

func init() {
	producer = new(impl.MockJobProducer)
}
func GetMockProducerInstance() *impl.MockJobProducer {
	return producer
}
