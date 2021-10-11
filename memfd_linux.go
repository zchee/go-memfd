// Copyright 2021 The go-memfd Authors
// SPDX-License-Identifier: BSD-3-Clause

package memfd

import (
	"golang.org/x/sys/unix"
)

func Create() (fd int, err error) {
	return unix.MemfdCreate("shm_anon", unix.MFD_CLOEXEC)
}
