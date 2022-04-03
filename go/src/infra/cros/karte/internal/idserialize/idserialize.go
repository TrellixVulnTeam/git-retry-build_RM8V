// Copyright 2021 The Chromium OS Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package idserialize

import (
	"bytes"
	"encoding/binary"
	"fmt"

	"infra/cros/karte/internal/lex64"
)

// Further fields may be added after Disambiguation without breaking backward-compatibility.
// However, adding a new field before a field that currently exists WILL break backward compatibility.
// If you are going to do this, please change the version.
type IDInfo struct {
	// Should be "zzzz" initially.
	Version        string
	CoarseTime     uint64
	FineTime       uint32
	Disambiguation uint32
}

// VersionlessBytes converts an IDInfo into bytes. Note that we use big-endian order so that lexicographical comparisons of IDInfo
// correspond to lexicographical byte comparisons.
func (i *IDInfo) VersionlessBytes() ([]byte, error) {
	buf := new(bytes.Buffer)
	for _, x := range []interface{}{
		i.CoarseTime,
		i.FineTime,
		i.Disambiguation,
	} {
		if err := binary.Write(buf, binary.BigEndian, x); err != nil {
			return nil, err
		}
	}
	return buf.Bytes(), nil
}

// Encoded converts an IDInfo into lex64, which preserves lexicographic order.
func (i *IDInfo) Encoded() (string, error) {
	bytes, err := i.VersionlessBytes()
	if err != nil {
		return "", err
	}
	encoded, err := lex64.Encode(bytes, false)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s%s", i.Version, encoded), nil
}
