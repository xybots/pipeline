// Copyright © 2018 Banzai Cloud
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package schedules

import (
	"net/http"

	"emperror.dev/errors"
	"github.com/gin-gonic/gin"

	"github.com/banzaicloud/pipeline/internal/platform/gin/correlationid"
	"github.com/banzaicloud/pipeline/src/api/ark/common"
)

// Get gets an ARK schedule
func Get(c *gin.Context) {
	scheduleName := c.Param("name")

	logger := correlationid.LogrusLogger(common.Log, c).WithField("schedule", scheduleName)
	logger.Info("getting schedule")

	schedule, err := common.GetARKService(c.Request).GetSchedulesService().GetByName(scheduleName)
	if err != nil {
		err = errors.WrapIf(err, "could not get schedule")
		common.ErrorHandler.Handle(err)
		common.ErrorResponse(c, err)
		return
	}

	c.JSON(http.StatusOK, schedule)
}