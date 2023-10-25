package tunnel

import (
	"os"
	"testing"

	"net/http"
	"net/http/httptest"
	"net/http/httputil"
	"net/url"
	"path/filepath"

	"github.com/kitabisa/teler-proxy/common"
	"github.com/kitabisa/teler-waf"
)

var (
	cwd, workspaceDir string

	dest = "http://example.com"

	handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
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
	t.Run("Valid Destination & No Config File", func(t *testing.T) {
		tun, err := NewTunnel(8080, dest, "", "", nil)
		if err != nil {
			t.Fatalf("Expected no error, but got: %v", err)
		}

		if tun == nil {
			t.Fatal("Expected Tunnel instance, but got nil")
		}
	})

	// Test case 2: invalid destination (empty)
	t.Run("Invalid Destination", func(t *testing.T) {
		_, err := NewTunnel(8080, "", "", "", nil)
		if err != common.ErrDestAddressEmpty {
			t.Fatalf("Expected %v, but got: %v", common.ErrDestAddressEmpty, err)
		}
	})

	// Test case 3: with config file but empty format
	t.Run("With Config File but Empty Format", func(t *testing.T) {
		_, err := NewTunnel(8080, dest, filepath.Join(workspaceDir, "teler-waf.conf.example.yaml"), "", nil)
		if err != common.ErrCfgFileFormatUnd {
			t.Fatalf("Expected %v, but got: %v", common.ErrCfgFileFormatUnd, err)
		}
	})

	// Test case 4: with config file and YAML format
	t.Run("With Config File and YAML Format", func(t *testing.T) {
		tun, err := NewTunnel(8080, dest, filepath.Join(workspaceDir, "teler-waf.conf.example.yaml"), "yaml", nil)
		if err != nil {
			t.Fatalf("Expected no error, but got: %v", err)
		}
		if tun == nil {
			t.Fatal("Expected Tunnel instance, but got nil")
		}
	})

	// Test case 5: with config file and JSON format
	t.Run("With Config File and JSON Format", func(t *testing.T) {
		tun, err := NewTunnel(8080, dest, filepath.Join(workspaceDir, "teler-waf.conf.example.json"), "json", nil)
		if err != nil {
			t.Fatalf("Expected no error, but got: %v", err)
		}
		if tun == nil {
			t.Fatal("Expected Tunnel instance, but got nil")
		}
	})

	// Test case 6: with config file and xml format
	t.Run("With Config File and XML Format", func(t *testing.T) {
		_, err := NewTunnel(8080, dest, filepath.Join(workspaceDir, "teler-waf.conf.example.json"), "xml", nil)
		if err != common.ErrCfgFileFormatInv {
			t.Fatalf("Expected %v, but got: %v", common.ErrCfgFileFormatInv, err)
		}
	})

	// Test case 7: invalid destination
	t.Run("Invalid Destination", func(t *testing.T) {
		tun, _ := NewTunnel(8080, "http://this is not a valid URL", "", "", nil)
		if tun != nil {
			t.Fatalf("Expected %v, but got: %v", nil, tun)
		}
	})

	// Test case 8: with invalid config file
	t.Run("Invalid YAML Config File", func(t *testing.T) {
		_, err := NewTunnel(8080, dest, "nonexistent", "yaml", nil)
		if err == nil {
			t.Fatal("Expected error, but got nil")
		}
	})

	// Test case 9: with invalid config file
	t.Run("Invalid JSON Config File", func(t *testing.T) {
		_, err := NewTunnel(8080, dest, "nonexistent", "json", nil)
		if err == nil {
			t.Fatal("Expected no error, but got nil")
		}
	})

	// Test case 10: with io.Writer
	t.Run("With Writer", func(t *testing.T) {
		_, err := NewTunnel(8080, dest, filepath.Join(workspaceDir, "teler-waf.conf.example.yaml"), "yaml", os.Stderr)
		if err != nil {
			t.Fatalf("Expected no error, but got: %v", err)
		}
	})
}

func TestServeHTTP(t *testing.T) {
	ts := httptest.NewServer(handler)
	defer ts.Close()

	parsedURL, _ := url.Parse(ts.URL)
	mockReverseProxy := httputil.NewSingleHostReverseProxy(parsedURL)

	tunnel := &Tunnel{
		Teler:        teler.New(),
		ReverseProxy: mockReverseProxy,
	}

	// Create a mock HTTP request and response recorder
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}
	recorder := httptest.NewRecorder()

	tunnel.ServeHTTP(recorder, req)

	if recorder.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, recorder.Code)
	}
}

func BenchmarkNewTunnel(b *testing.B) {
	b.Run("YAML", func(b *testing.B) {
		b.ResetTimer()
		b.ReportAllocs()

		for i := 0; i < b.N; i++ {
			_, err := NewTunnel(8080, dest, filepath.Join(workspaceDir, "teler-waf.conf.example.yaml"), "yaml", nil)
			if err != nil {
				b.Fatalf("Expected no error, but got: %v", err)
			}
		}
	})

	b.Run("JSON", func(b *testing.B) {
		b.ResetTimer()
		b.ReportAllocs()

		for i := 0; i < b.N; i++ {
			_, err := NewTunnel(8080, dest, filepath.Join(workspaceDir, "teler-waf.conf.example.json"), "json", nil)
			if err != nil {
				b.Fatalf("Expected no error, but got: %v", err)
			}
		}
	})
}

func BenchmarkServeHTTP(b *testing.B) {
	ts := httptest.NewServer(handler)
	defer ts.Close()

	parsedURL, _ := url.Parse(ts.URL)
	mockReverseProxy := httputil.NewSingleHostReverseProxy(parsedURL)

	tunnel := &Tunnel{
		Teler:        teler.New(),
		ReverseProxy: mockReverseProxy,
	}

	// Create a mock HTTP request and response recorder
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		b.Fatal(err)
	}
	recorder := httptest.NewRecorder()

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		tunnel.ServeHTTP(recorder, req)

		if recorder.Code != http.StatusOK {
			b.Errorf("Expected status code %d, but got %d", http.StatusOK, recorder.Code)
		}
	}
}
