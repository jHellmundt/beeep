//go:build windows && !linux && !freebsd && !netbsd && !openbsd && !darwin && !js
// +build windows,!linux,!freebsd,!netbsd,!openbsd,!darwin,!js

package beeep

import (
	toast "github.com/go-toast/toast"
)

// Alert displays a desktop notification and plays a default system sound.
func Alert(title, message, appIcon, appID string) error {
	if isWindows10 {
		note := toastNotification(title, message, pathAbs(appIcon), appID)
		note.Audio = toast.Default
		return note.Push()
	}

	if err := Notify(title, message, appIcon, appID); err != nil {
		return err
	}
	return Beep(DefaultFreq, DefaultDuration)
}
