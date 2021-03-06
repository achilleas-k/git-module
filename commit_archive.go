// Copyright 2015 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package git

import (
	"fmt"
	"path/filepath"
	"strings"
)

type ArchiveType int

const (
	ArchiveZip ArchiveType = iota + 1
	ArchiveTarGz
	ArchiveGIN
)

func (c *Commit) CreateArchive(target string, archiveType ArchiveType, cloneL string) error {
	var format string
	switch archiveType {
	case ArchiveZip:
		format = "zip"
	case ArchiveTarGz:
		format = "tar.gz"
	default:
		return fmt.Errorf("unknown format: %v", archiveType)
	}

	_, err := NewCommand("archive", "--prefix="+filepath.Base(strings.TrimSuffix(c.repo.Path, ".git"))+"/", "--format="+format, "-o", target, c.ID.String()).RunInDir(c.repo.Path)
	return err
}
