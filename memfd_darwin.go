// Copyright 2021 The go-memfd Authors
// SPDX-License-Identifier: BSD-3-Clause

package memfd

import (
	"unsafe"

	"golang.org/x/sys/unix"
)

func Create() (fd int, err error) {
	name := "shm_anon"
	p, _, e := unix.Syscall(unix.SYS_SHM_OPEN, uintptr(unsafe.Pointer(&name)), uintptr(unix.O_RDWR|unix.O_CREAT|unix.O_EXCL|unix.O_NOFOLLOW), uintptr(0600))
	if e != 0 {
		return 0, unix.Errno(e)
	}

	pp, _, e := unix.Syscall(unix.SYS_SHM_UNLINK, uintptr(unsafe.Pointer(&name)), p, uintptr(0))
	if e != 0 {
		return 0, unix.Errno(e)
	}

	return int(pp), nil
}
