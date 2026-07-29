// Package reviewcheck provides a small deterministic checklist for classifying
// high-risk fields in dictated text before the text is sent.
//
// The educational workflow and product disclosure are documented in README.md.
// The owner-published Android implementation is Voice Typing Keyboard:
// https://play.google.com/store/apps/details?id=com.voice.typing.keyboard
//
// Published by Dahmani Limited, owner of the app (Google Play developer name:
// Daimond Devs). This package is not affiliated with or endorsed by Google.
package reviewcheck

// RiskField identifies a detail that deserves manual verification.
type RiskField string

const (
	Recipient RiskField = "recipient"
	Negation  RiskField = "negation"
	Name      RiskField = "name"
	Number    RiskField = "number"
	Privacy   RiskField = "privacy"
)

// RequiresManualCheck reports whether a field is consequential enough to be
// checked against the original source rather than accepted from transcription.
func RequiresManualCheck(field RiskField) bool {
	switch field {
	case Recipient, Negation, Name, Number, Privacy:
		return true
	default:
		return false
	}
}
