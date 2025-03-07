//go:build !windows

/*
   Copyright The containerd Authors.

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

package fs

import (
	"bytes"
	"fmt"
	"os"
	"syscall"

	"github.com/containerd/continuity/sysx"
)

// compareSysStat returns whether the stats are equivalent,
// whether the files are considered the same file, and
// an error
func compareSysStat(s1, s2 interface{}) (bool, error) {
	ls1, ok := s1.(*syscall.Stat_t)
	if !ok {
		return false, nil
	}
	ls2, ok := s2.(*syscall.Stat_t)
	if !ok {
		return false, nil
	}

	return ls1.Mode == ls2.Mode && ls1.Uid == ls2.Uid && ls1.Gid == ls2.Gid && ls1.Rdev == ls2.Rdev, nil
}

func compareCapabilities(p1, p2 string) (bool, error) {
	c1, err := sysx.LGetxattr(p1, "security.capability")
	if err != nil && err != sysx.ENODATA {
		return false, fmt.Errorf("failed to get xattr for %s: %w", p1, err)
	}
	c2, err := sysx.LGetxattr(p2, "security.capability")
	if err != nil && err != sysx.ENODATA {
		return false, fmt.Errorf("failed to get xattr for %s: %w", p2, err)
	}
	return bytes.Equal(c1, c2), nil
}

func isLinked(f os.FileInfo) bool {
	s, ok := f.Sys().(*syscall.Stat_t)
	if !ok {
		return false
	}
	return !f.IsDir() && s.Nlink > 1
}
