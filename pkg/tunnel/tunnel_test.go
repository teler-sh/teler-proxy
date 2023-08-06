package tunnel

import (
	"testing"

	"github.com/kitabisa/teler-proxy/common"
)

func TestNewTunnel(t *testing.T) {
	// Test case 1: valid destination and no configuration file
	tun, err := NewTunnel(8080, "http://example.com", "", "")
	if err != nil {
		t.Fatalf("Expected no error, but got: %v", err)
	}

	if tun == nil {
		t.Fatal("Expected Tunnel instance, but got nil")
	}

	// Test case 2: invalid destination (empty)
	_, err = NewTunnel(8080, "", "", "")
	if err != common.ErrDestAddressEmpty {
		t.Fatalf("Expected %v, but got: %v", common.ErrDestAddressEmpty, err)
	}
}
