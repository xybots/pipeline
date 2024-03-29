// Copyright © 2020 Banzai Cloud
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

package clustermodel

import (
	"database/sql/driver"
	"fmt"
	"time"

	"github.com/gofrs/uuid"
	"github.com/jinzhu/gorm"

	"github.com/banzaicloud/pipeline/internal/database/sql/json"
	"github.com/banzaicloud/pipeline/src/secret"
)

const unknownLocation = "unknown"

const InstanceTypeSeparator = " "

// ClusterModel describes the common cluster model.
type ClusterModel struct {
	ID  uint   `gorm:"primary_key"`
	UID string `gorm:"unique_index:idx_clusters_uid"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `gorm:"unique_index:idx_clusters_unique_id" sql:"index"`
	StartedAt *time.Time
	CreatedBy uint

	Name           string `gorm:"unique_index:idx_clusters_unique_id"`
	Location       string
	Cloud          string
	Distribution   string
	OrganizationID uint `gorm:"unique_index:idx_clusters_unique_id"`
	SecretID       string
	ConfigSecretID string
	SSHSecretID    string
	Status         string
	RbacEnabled    bool
	OidcEnabled    bool        `gorm:"default:false;not null"`
	StatusMessage  string      `sql:"type:text;"`
	Tags           ClusterTags `gorm:"type:json"`
}

type ClusterTags map[string]string

// TableName changes the default table name.
func (ClusterModel) TableName() string {
	return "clusters"
}

func (m *ClusterModel) BeforeCreate() (err error) {
	if m.UID == "" {
		m.UID = uuid.Must(uuid.NewV4()).String()
	}

	return
}

// AfterFind converts Location field(s) to unknown if they are empty.
func (m *ClusterModel) AfterFind() error {
	if len(m.Location) == 0 {
		m.Location = unknownLocation
	}

	return nil
}

// String method prints formatted cluster fields.
func (m ClusterModel) String() string {
	return fmt.Sprintf("Id: %d, Creation date: %s, Cloud: %s, Distribution: %s", m.ID, m.CreatedAt, m.Cloud, m.Distribution)
}

// BeforeDelete should not be declared on this model.
// TODO: please move this to the cluster delete flow
// this should not have been added here in the first place!!!!!!!
func (m ClusterModel) BeforeDelete(tx *gorm.DB) (err error) {
	logger := log.WithFields(map[string]interface{}{"organization": m.OrganizationID, "cluster": m.ID})

	logger.Info("Delete unused cluster secrets")
	if err := secret.Store.DeleteByClusterUID(m.OrganizationID, m.UID); err != nil {
		logger.Error(fmt.Sprintf("Error during deleting secret: %s", err.Error()))
	}

	return
}

func (fs *ClusterTags) Scan(src interface{}) error {
	return json.Scan(src, fs)
}

func (fs ClusterTags) Value() (driver.Value, error) {
	return json.Value(fs)
}
