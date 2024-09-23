/*
Copyright 2024 The Tekton Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package tektonconfig

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tektoncd/operator/pkg/reconciler/openshift"
)

func TestLoadRbacMaxConcurrentCalls(t *testing.T) {
	for _, tt := range []struct {
		desc          string
		envValue      string
		expectedValue int
	}{
		{"empty envValue", "", defaultRbacMaxConcurrentCalls},
		{"valid envValue", "10", 10},
		{"below min envValue", "-1", defaultRbacMaxConcurrentCalls},
		{"above max envValye", "60", defaultRbacMaxConcurrentCalls},
		{"invalid envValue", "xyz", defaultRbacMaxConcurrentCalls},
	} {
		t.Run(tt.desc, func(t *testing.T) {
			os.Setenv(openshift.RbacProvisioningMaxConcurrentCalls, tt.envValue)
			result := loadRbacMaxConcurrentCalls()
			require.Equal(t, result, tt.expectedValue)
		})
	}
}
