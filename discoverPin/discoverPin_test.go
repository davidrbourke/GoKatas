package discoverpin

import (
	"testing"
)

func TestDiscoverPinGetPin(t *testing.T) {

	keys := []string{"352", "384", "845", "352"}
	result := GetPin(keys)
	if result != "38452" {
		t.Errorf("Did not find correct pin: %v", result)
	}
}
