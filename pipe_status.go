//
// Copyright 2020 Bryan T. Meyers <root@datadrake.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package main

import (
	"os"
)

func pipeStatus() (fns []pieceFn) {
	failure := false
	// Generate a piece for each return code
	for _, content := range os.Args[1:] {
		// Assume success
		fn := func() *Piece {
			return &Piece{
				content: content,
				fg:      "0",
				bg:      "46",
			}
		}
		// override for failure
		if content != "0" {
			failure = true
			fn = func() *Piece {
				return &Piece{
					content: content,
					fg:      "15",
					bg:      "160",
				}
			}
		}
		fns = append(fns, fn)
	}
	// If no failures, don't produce any output
	if !failure {
		fns = make([]pieceFn, 0)
	}
	return
}
