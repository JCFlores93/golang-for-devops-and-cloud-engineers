package api

import "testing"

func TestDoGetRequest(t *testing.T) {
	apiInstance := api{
		Options: Options{},
		Client:  http.Client,
	}
	apiInstance.DoRequest{}
}
