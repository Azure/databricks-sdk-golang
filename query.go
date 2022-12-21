// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package databricks

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/google/go-querystring/query"
)

// PerformQuery can be used in a client or directly
func PerformQuery(option DBClientOption, method, path string, data interface{}, headers map[string]string) ([]byte, error) {

	requestURL, err := option.getRequestURI(path)
	if err != nil {
		return nil, err
	}

	requestHeaders := option.getDefaultHeaders()

	if len(headers) > 0 {
		for k, v := range headers {
			requestHeaders[k] = v
		}
	}

	var requestBody []byte
	if method == "GET" {
		params, err := query.Values(data)
		if err != nil {
			return nil, err
		}
		requestURL += "?" + params.Encode()
	} else if data != nil {
		bodyBytes, err := json.Marshal(data)
		if err != nil {
			return nil, err
		}
		requestBody = bodyBytes
	}

	client := option.getHTTPClient()

	request, err := http.NewRequest(method, requestURL, bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, err
	}
	for k, v := range requestHeaders {
		request.Header.Set(k, v)
	}

	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("Response from server (%d) %s", resp.StatusCode, string(body))
	}

	return body, nil
}
