// Copyright © 2018 Cisco Systems, Inc.
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

package multihelmlib

type Apps struct {
	Apps []App
}

func (a *Apps) Apply(configFile string, appsPath string, printRendered bool) {
	for _, app := range a.Apps {
		app.Apply(configFile, appsPath, printRendered)
	}
}

func (a *Apps) Destroy(purge bool) {
	for _, app := range a.Apps {
		app.Destroy(purge)
	}
}

func (a *Apps) Id() []string {
	var ids []string
	for _, app := range a.Apps {
		id := app.Id()
		ids = append(ids, id)
	}
	return ids
}

func (a *Apps) Simulate(configFile string, appsPath string, printRendered bool) {
	for _, app := range a.Apps {
		app.Simulate(configFile, appsPath, printRendered)
	}
}

func (a *Apps) Status() {
	for _, app := range a.Apps {
		app.Status()
	}
}
