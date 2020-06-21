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
	"strings"
)

func dir() []Piece {
	var dirs []string
	parts := strings.Split(WorkDir, "/")
	// Deal with HOME
	if strings.HasPrefix(WorkDir, HomeDir) {
		dirs = append(dirs, "~")
		path := strings.TrimPrefix(WorkDir, HomeDir)
		parts = strings.Split(path, "/")
		parts = parts[1:]
	}
	// Show part of the beginning of the path for context
	if len(parts) > 1 {
		lim := 1
		if parts[0] == "" {
			lim++
		}
		dirs = append(dirs, parts[0:lim]...)
		parts = parts[lim:]
	}
	// Only show the last bit of the path
	if len(parts) > 2 {
		dirs = append(dirs, "…")
		parts = parts[len(parts)-2:]
	}
	dirs = append(dirs, parts...)
	return []Piece{
		Piece{
			Content: strings.Join(dirs, "/"),
			FG:      97,
			BG:      100,
		},
	}
}
