// Copyright © 2019 Banzai Cloud
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

package securityscanadapter

import (
	"context"

	"github.com/banzaicloud/pipeline/internal/integratedservices/services"
	"github.com/banzaicloud/pipeline/internal/secret/secrettype"
	"github.com/banzaicloud/pipeline/src/secret"
)

// UserSecretStore stores Anchore user secrets.
type UserSecretStore struct {
	secretStore services.SecretStore
}

// NewUserSecretStore returns a new UserSecretStore.
func NewUserSecretStore(secretStore services.SecretStore) UserSecretStore {
	return UserSecretStore{
		secretStore: secretStore,
	}
}

func (s UserSecretStore) GetPasswordForUser(ctx context.Context, userName string) (string, error) {
	values, err := s.secretStore.GetSecretValues(ctx, secret.GenerateSecretIDFromName(userName))
	if err != nil {
		return "", err
	}

	return values[secrettype.Password], nil
}
