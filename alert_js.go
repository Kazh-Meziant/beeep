//go:build js
// +build js

package beeep

// Alert displays a desktop notification and plays a beep.
func Alert(title, message, appIcon, _ string) error {
	if err := Notify(title, message, appIcon, _); err != nil {
		return err
	}
	return Beep(DefaultFreq, DefaultDuration)
}
