//go:build !windows && !plan9
// +build !windows,!plan9

package limits

import (
	"syscall"

	"github.com/pkg/errors"
)

// SetLimits raises some process limits to values
func SetLimits(desiredLimits *DesiredLimits) error {
	var rLimit syscall.Rlimit

	err := syscall.Getrlimit(syscall.RLIMIT_NOFILE, &rLimit)
	if err != nil {
		return err
	}
	if rLimit.Cur > desiredLimits.FileLimitWant {
		return nil
	}
	if rLimit.Max < desiredLimits.FileLimitMin {
		err = errors.Errorf("need at least %d file descriptors",
			desiredLimits.FileLimitMin)
		return err
	}
	if rLimit.Max < desiredLimits.FileLimitWant {
		rLimit.Cur = rLimit.Max
	} else {
		rLimit.Cur = desiredLimits.FileLimitWant
	}
	err = syscall.Setrlimit(syscall.RLIMIT_NOFILE, &rLimit)
	if err != nil {
		// try min value
		rLimit.Cur = desiredLimits.FileLimitMin
		err = syscall.Setrlimit(syscall.RLIMIT_NOFILE, &rLimit)
		if err != nil {
			return err
		}
	}

	return nil
}
