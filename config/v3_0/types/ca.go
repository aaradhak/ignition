// Copyright 2018 CoreOS, Inc.
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

package types

import (
	"github.com/ajeddeloh/vcontext/path"
	"github.com/ajeddeloh/vcontext/report"
)

func (c CaReference) Key() string {
	return c.Source
}

func (ca CaReference) Validate(c path.ContextPath) (r report.Report) {
	r.AddOnError(c.Append("source"), validateURL(ca.Source))
	return
}
