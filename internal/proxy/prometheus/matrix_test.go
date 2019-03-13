/**
* Copyright 2018 Comcast Cable Communications Management, LLC
* Licensed under the Apache License, Version 2.0 (the "License");
* you may not use this file except in compliance with the License.
* You may obtain a copy of the License at
* http://www.apache.org/licenses/LICENSE-2.0
* Unless required by applicable law or agreed to in writing, software
* distributed under the License is distributed on an "AS IS" BASIS,
* WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
* See the License for the specific language governing permissions and
* limitations under the License.
 */

package prometheus

import (
	"testing"
	"time"
)

func TestSetStep(t *testing.T) {
	me := MatrixEnvelope{}
	const step = time.Duration(300) * time.Minute
	me.SetStep(step)
	if me.StepDuration != step {
		t.Errorf(`wanted "%s". got "%s"`, step, me.StepDuration)
	}
}

func TestStep(t *testing.T) {
	me := MatrixEnvelope{}
	const step = time.Duration(300) * time.Minute
	me.SetStep(step)
	if me.Step() != step {
		t.Errorf(`wanted "%s". got "%s"`, step, me.Step())
	}
}
