package runner

import (
	"net"
	"time"
)

func isReachable(dest string, timeout time.Duration) bool {
	conn, err := net.DialTimeout("tcp", dest, timeout)
	if err != nil {
		return false
	}

	defer conn.Close()

	return true
}
