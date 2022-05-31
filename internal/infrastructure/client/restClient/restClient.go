// Package restClient
//======================================================================================================
// File: internal/infrastructure/client/restClient/restClient.go
// Authors: Everton Barbosa Kopec (Radar)
// Date: 30 may 2022
// Brief: RestClient client module - Conversion CSV to NDJSON
//
// COPYRIGHT © 2019 all rights reserved to Radar
//#======================================================================================================
package restClient

import (
	"context"
	"net/http"
)

func Get(ctx context.Context, URL string) (*http.Response, error) {
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, URL, nil)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
