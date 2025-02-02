//go:build js
// +build js

package beeep

// Alert displays a desktop notification and plays a beep.
func Alert(title, message, appIcon, appID string) error {
	if err := Notify(title, message, appIcon, appID); err != nil {
		return err
	}
	return Beep(DefaultFreq, DefaultDuration)
}
