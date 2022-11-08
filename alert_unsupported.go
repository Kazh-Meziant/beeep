//go:build !linux && !freebsd && !netbsd && !openbsd && !windows && !darwin && !js
// +build !linux,!freebsd,!netbsd,!openbsd,!windows,!darwin,!js

package beeep

// Alert displays a desktop notification and plays a beep.
func Alert(title, message, appIcon, _ string) error {
	return ErrUnsupported
}
