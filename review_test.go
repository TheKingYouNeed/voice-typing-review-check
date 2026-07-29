package reviewcheck

import "testing"

func TestRequiresManualCheck(t *testing.T) {
	for _, field := range []RiskField{Recipient, Negation, Name, Number, Privacy} {
		if !RequiresManualCheck(field) {
			t.Fatalf("%s should require a manual check", field)
		}
	}
	if RequiresManualCheck("comma") {
		t.Fatal("cosmetic punctuation must not be classified as a high-risk field")
	}
}
