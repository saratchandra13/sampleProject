package rest

import (
	"context"
	"errors"
	"fmt"
	"github.com/saratchandra13/sampleProject/utilities"
	"io/ioutil"
	"net/http"
	netUrl "net/url"
	"time"
)

type request struct {
	contentType string // Default value is 'application/json'
	headers     http.Header
	qs          netUrl.Values
	body        interface{}
}

type statusCode int

func makeRequest(
	ctx context.Context, client *http.Client, method,
	url string, r request, timeout time.Duration,
) ([]byte, statusCode, error) {
	ioReader, err := utilities.IoReaderFromBody(r.body)
	if err != nil {
		return nil, 0, err
	}

	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, method, url, ioReader)
	if err != nil {
		return nil, 0, err
	}

	utilities.UpdateQueryParams(req, r.qs)

	utilities.UpdateHeaders(req, r.headers)

	req.Header.Set("Content-Type", utilities.OrAssignment(r.contentType, "application/json"))

	resp, err := client.Do(req)
	if err != nil {
		return nil, 0, handleError(err, req)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, statusCode(resp.StatusCode), err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return body, statusCode(resp.StatusCode), errors.New(string(body))
	}

	return body, statusCode(resp.StatusCode), nil
}

func SimpleGetRequest(ctx context.Context, client *http.Client, url string, timeout time.Duration) ([]byte, statusCode, error) {

	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, 0, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, 0, handleError(err, req)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, statusCode(resp.StatusCode), err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return body, statusCode(resp.StatusCode), errors.New(string(body))
	}

	return body, statusCode(resp.StatusCode), nil

}

func handleError(err error, req *http.Request) error {
	if err == nil {
		return nil
	}
	var netURLError *netUrl.Error
	switch {
	case errors.As(err, &netURLError):
		return fmt.Errorf("%s %s %+v", netURLError.Op, utilities.OrAssignment(req.Host, netURLError.URL), netURLError.Err)
	default:
		return err
	}
}
