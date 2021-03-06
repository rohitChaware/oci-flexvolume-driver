// Copyright 2017 Oracle and/or its affiliates. All rights reserved.
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

package integration

import (
	"log"
	"testing"

	"github.com/oracle/oci-flexvolume-driver/test/integration/framework"
)

var fw *framework.Framework

// TestMain integration test entrypoint.
func TestMain(m *testing.M) {
	err := fw.Init()
	if err != nil {
		log.Fatal(err)
	}

	fw.Run(m.Run)
}

func init() {
	fw = framework.New()
}
