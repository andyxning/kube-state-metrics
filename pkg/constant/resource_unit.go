/*
Copyright 2018 The Kubernetes Authors All rights reserved.

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

// Package constant defines constant values used in kube-state-metrics project.
package constant

// ResourceUnit represents the unit of measure for certain metrics.
type ResourceUnit string

const (
	// UnitByte is the unit of measure in bytes.
	UnitByte ResourceUnit = "byte"
	// UnitCore is the unit of measure in CPU cores.
	UnitCore ResourceUnit = "core"
	// UnitInteger is the unit of measure in integers.
	UnitInteger ResourceUnit = "integer"
)
