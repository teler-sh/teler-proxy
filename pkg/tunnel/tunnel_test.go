package tunnel

import (
	"os"
	"testing"

	// "net/http"
	// "net/http/httptest"
	// "net/http/httputil"
	// "net/url"
	"path/filepath"

	"github.com/kitabisa/teler-proxy/common"
	// "github.com/kitabisa/teler-waf"
)

var (
	cwd, workspaceDir string
)

func init() {
	var err error

	cwd, err = os.Getwd()
	if err != nil {
		panic(err)
	}

	workspaceDir = filepath.Join(cwd, "..", "..")
}

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

	// Test case 3: with config file but empty format
	_, err = NewTunnel(8080, "http://example.com", filepath.Join(workspaceDir, "teler-waf.conf.example.yaml"), "")
	if err != common.ErrCfgFileFormatUnd {
		t.Fatalf("Expected %v, but got: %v", common.ErrCfgFileFormatUnd, err)
	}

	// Test case 4: with config file and YAML format
	tun = nil
	tun, err = NewTunnel(8080, "http://example.com", filepath.Join(workspaceDir, "teler-waf.conf.example.yaml"), "yaml")
	if err != nil {
		t.Fatalf("Expected no error, but got: %v", err)
	}

	if tun == nil {
		t.Fatal("Expected Tunnel instance, but got nil")
	}

	// Test case 5: with config file and JSON format
	tun = nil
	tun, err = NewTunnel(8080, "http://example.com", filepath.Join(workspaceDir, "teler-waf.conf.example.json"), "json")
	if err != nil {
		t.Fatalf("Expected no error, but got: %v", err)
	}

	if tun == nil {
		t.Fatal("Expected Tunnel instance, but got nil")
	}

	// Test case 6: with config file and xml format
	_, err = NewTunnel(8080, "http://example.com", filepath.Join(workspaceDir, "teler-waf.conf.example.json"), "xml")
	if err != common.ErrCfgFileFormatInv {
		t.Fatalf("Expected %v, but got: %v", common.ErrCfgFileFormatInv, err)
	}

	// Test case 7: invalid destination
	tun = nil
	tun, _ = NewTunnel(8080, "http://this is not a valid URL", "", "")
	if tun != nil {
		t.Fatalf("Expected %v, but got: %v", nil, tun)
	}
}

// TODO(dwisiswant0): make these test works
// func TestServeHTTP(t *testing.T) {
// 	parsedURL, _ := url.Parse("http://localhost")
// 	mockReverseProxy := httputil.NewSingleHostReverseProxy(parsedURL)

// 	tunnel := &Tunnel{
// 		Teler:        teler.New(),
// 		ReverseProxy: mockReverseProxy,
// 	}

// 	// Create a mock HTTP request and response recorder
// 	req, err := http.NewRequest("POST", "/", nil)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	recorder := httptest.NewRecorder()

// 	tunnel.ServeHTTP(recorder, req)

// 	if recorder.Code != http.StatusOK {
// 		t.Errorf("Expected status code %d, but got %d", http.StatusOK, recorder.Code)
// 	}
// }
