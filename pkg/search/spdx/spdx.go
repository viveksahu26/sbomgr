// Copyright 2023 Interlynk.io
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

package spdx

import (
	"fmt"

	"github.com/interlynk-io/sbomgr/pkg/search/options"
	"github.com/interlynk-io/sbomgr/pkg/search/results"
)

type SpdxModule struct {
	ro *options.RuntimeOptions
	so options.SearchOptions
}

func (s *SpdxModule) SetRuntimeOptions(ro *options.RuntimeOptions) {
	s.ro = ro
}

func (s *SpdxModule) SetSearchOptions(so options.SearchOptions) {
	s.so = so
}

func (s *SpdxModule) Search() (*results.Result, error) {
	doc, err := loadDoc(s)
	if err != nil {
		return nil, err
	}
	pkgIdx := doc.searchPackages(s)
	fmt.Printf("file:%s spec:%s pkgIdx: %v", s.ro.CurrentPath, s.ro.SbomSpecType, pkgIdx)
	return doc.constructResults(s, pkgIdx)
}