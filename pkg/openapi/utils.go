package openapi

import (
	"errors"
	"fmt"
	"net/http"
)

func CheckStatus(resp *http.Response) error {

	switch code := resp.StatusCode; {
	case code == http.StatusUnauthorized:
		return fmt.Errorf("unauthorized")
	case code == http.StatusForbidden:
		return errors.New("forbidden")
	case code == 422:
		return errors.New("validation")
	case code >= 500:
		return fmt.Errorf("server error: %d", code)
	case code < 300:
		return nil
	default:
		return fmt.Errorf("unrecognized status code %d", code)
	}
}
