/*
Copyright 2019 The Crossplane Authors.

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

package oam

// Label key strings.
// AppConfig controller will add these labels into workloads.
const (
	// LabelAppName records the name of AppConfig
	LabelAppName = "app.oam.dev/name"
	// LabelAppComponent records the name of Component
	LabelAppComponent = "app.oam.dev/component"
	// LabelAppComponentRevision records the revision name of Component
	LabelAppComponentRevision = "app.oam.dev/revision"
)
