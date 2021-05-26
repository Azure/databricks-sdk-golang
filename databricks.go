package databricks

import (
	"crypto/tls"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

// DBClientOption is used to configure the DataBricks Client
type DBClientOption struct {
	User               string
	Password           string
	Host               string
	Token              string
	DefaultHeaders     map[string]string
	InsecureSkipVerify bool
	TimeoutSeconds     int
	client             http.Client
}

// NewDBClientOption retruns the new DBClientOption
func NewDBClientOption(user, password, host, token string, defaultHeaders map[string]string, insecureSkipVerify bool, timeoutSeconds int) *DBClientOption {
	if timeoutSeconds == 0 {
		timeoutSeconds = 10
	}
	client := http.Client{
		Timeout: time.Duration(time.Duration(timeoutSeconds) * time.Second),
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: insecureSkipVerify,
			},
		},
	}
	return &DBClientOption{
		User:               user,
		Password:           password,
		Host:               host,
		Token:              token,
		DefaultHeaders:     defaultHeaders,
		InsecureSkipVerify: insecureSkipVerify,
		client:             client,
	}
}

func (o *DBClientOption) getHTTPClient() http.Client {
	return o.client
}

func (o *DBClientOption) getAuthHeader() map[string]string {
	auth := make(map[string]string)
	if o.User != "" && o.Password != "" {
		encodedAuth := []byte(o.User + ":" + o.Password)
		userHeaderData := "Basic " + base64.StdEncoding.EncodeToString(encodedAuth)
		auth["Authorization"] = userHeaderData
		auth["Content-Type"] = "application/json"
	} else if o.Token != "" {
		auth["Authorization"] = "Bearer " + o.Token
		auth["Content-Type"] = "application/json"
	}
	return auth
}

func (o *DBClientOption) getUserAgentHeader() map[string]string {
	return map[string]string{
		"User-Agent": fmt.Sprintf("databricks-sdk-golang-%s", SdkVersion),
	}
}

func (o *DBClientOption) getDefaultHeaders() map[string]string {
	auth := o.getAuthHeader()
	userAgent := o.getUserAgentHeader()

	defaultHeaders := make(map[string]string)
	for k, v := range auth {
		defaultHeaders[k] = v
	}
	for k, v := range o.DefaultHeaders {
		defaultHeaders[k] = v
	}
	for k, v := range userAgent {
		defaultHeaders[k] = v
	}
	return defaultHeaders
}

func (o *DBClientOption) getRequestURI(path string) (string, error) {
	parsedURI, err := url.Parse(o.Host)
	if err != nil {
		return "", err
	}
	requestURI := fmt.Sprintf("%s://%s/api/%s%s", parsedURI.Scheme, parsedURI.Host, APIVersion, path)
	return requestURI, nil
}
