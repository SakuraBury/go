package response

import (
	"common/models"
	"common/models/job"
)

type HeartbeatResponse struct {
	models.BaseResult
	Cmd *job.JobCmd
}
