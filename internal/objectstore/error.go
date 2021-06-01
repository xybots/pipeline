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

package objectstore

import (
	"github.com/banzaicloud/pipeline/pkg/objectstore"
)

// IsAlreadyExistsError checks if an error indicates an already existing bucket.
func IsAlreadyExistsError(err error) bool {
	return objectstore.IsAlreadyExistsError(err)
}

// IsNotFoundError checks if an error indicates a missing bucket.
func IsNotFoundError(err error) bool {
	return objectstore.IsNotFoundError(err)
}
