package util

import (
	"fmt"
	"net/http"
)

func HandleHTTPResponse(resp *http.Response) error {
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		return nil
	}

	switch resp.StatusCode {
	case http.StatusUnauthorized:
		return fmt.Errorf("request failed: server returned %d Unauthorized", resp.StatusCode)
	case http.StatusForbidden:
		return fmt.Errorf("request failed: server returned %d Forbidden", resp.StatusCode)
	default:
		return fmt.Errorf("request failed: server returned %d", resp.StatusCode)
	}
}
