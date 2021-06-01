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

package restores

import (
	"emperror.dev/errors"
	"github.com/gin-gonic/gin"

	"github.com/banzaicloud/pipeline/internal/platform/gin/correlationid"
	ginutils "github.com/banzaicloud/pipeline/internal/platform/gin/utils"
	"github.com/banzaicloud/pipeline/src/api/ark/common"
)

// GetLogs get logs for an ARK restore
func GetLogs(c *gin.Context) {
	logger := correlationid.LogrusLogger(common.Log, c)

	restoreID, ok := ginutils.UintParam(c, IDParamName)
	if !ok {
		return
	}

	logger = logger.WithField("restore", restoreID)
	logger.Info("getting restore logs")

	svc := common.GetARKService(c.Request)

	restore, err := svc.GetRestoresService().GetByID(restoreID)
	if err != nil {
		err = errors.WrapIf(err, "could not get restore")
		common.ErrorHandler.Handle(err)
		common.ErrorResponse(c, err)
		return
	}

	if restore.Bucket == nil {
		err = errors.New("could not find the related bucket")
		common.ErrorHandler.Handle(err)
		common.ErrorResponse(c, err)
		return
	}

	err = svc.GetBucketsService().StreamRestoreLogsFromObjectStore(
		restore.Bucket,
		restore.BackupName,
		restore.Name,
		c.Writer,
	)
	if err != nil {
		err = errors.WrapIf(err, "could not stream logs")
		common.ErrorHandler.Handle(err)
		common.ErrorResponse(c, err)
		return
	}
}
