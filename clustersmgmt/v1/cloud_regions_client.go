/*
Copyright (c) 2019 Red Hat, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

  http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// IMPORTANT: This file has been generated automatically, refrain from modifying it manually as all
// your changes will be lost when the file is generated again.

package v1 // github.com/openshift-online/ocm-sdk-go/clustersmgmt/v1

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"path"

	"github.com/openshift-online/ocm-sdk-go/errors"
	"github.com/openshift-online/ocm-sdk-go/helpers"
)

// CloudRegionsClient is the client of the 'cloud_regions' resource.
//
// Manages the collection of regions of a cloud provider.
type CloudRegionsClient struct {
	transport http.RoundTripper
	path      string
	metric    string
}

// NewCloudRegionsClient creates a new client for the 'cloud_regions'
// resource using the given transport to sned the requests and receive the
// responses.
func NewCloudRegionsClient(transport http.RoundTripper, path string, metric string) *CloudRegionsClient {
	client := new(CloudRegionsClient)
	client.transport = transport
	client.path = path
	client.metric = metric
	return client
}

// List creates a request for the 'list' method.
//
// Retrieves the list of regions of the cloud provider.
//
// IMPORTANT: This collection doesn't currently support paging or searching, so the returned
// `page` will always be 1 and `size` and `total` will always be the total number of regions
// of the provider.
func (c *CloudRegionsClient) List() *CloudRegionsListRequest {
	request := new(CloudRegionsListRequest)
	request.transport = c.transport
	request.path = c.path
	request.metric = c.metric
	return request
}

// Region returns the target 'cloud_region' resource for the given identifier.
//
// Reference to the service that manages an specific region.
func (c *CloudRegionsClient) Region(id string) *CloudRegionClient {
	return NewCloudRegionClient(
		c.transport,
		path.Join(c.path, id),
		path.Join(c.metric, "-"),
	)
}

// CloudRegionsListRequest is the request for the 'list' method.
type CloudRegionsListRequest struct {
	transport http.RoundTripper
	path      string
	metric    string
	query     url.Values
	header    http.Header
}

// Parameter adds a query parameter.
func (r *CloudRegionsListRequest) Parameter(name string, value interface{}) *CloudRegionsListRequest {
	helpers.AddValue(&r.query, name, value)
	return r
}

// Header adds a request header.
func (r *CloudRegionsListRequest) Header(name string, value interface{}) *CloudRegionsListRequest {
	helpers.AddHeader(&r.header, name, value)
	return r
}

// Send sends this request, waits for the response, and returns it.
//
// This is a potentially lengthy operation, as it requires network communication.
// Consider using a context and the SendContext method.
func (r *CloudRegionsListRequest) Send() (result *CloudRegionsListResponse, err error) {
	return r.SendContext(context.Background())
}

// SendContext sends this request, waits for the response, and returns it.
func (r *CloudRegionsListRequest) SendContext(ctx context.Context) (result *CloudRegionsListResponse, err error) {
	query := helpers.CopyQuery(r.query)
	header := helpers.SetHeader(r.header, r.metric)
	uri := &url.URL{
		Path:     r.path,
		RawQuery: query.Encode(),
	}
	request := &http.Request{
		Method: http.MethodGet,
		URL:    uri,
		Header: header,
	}
	if ctx != nil {
		request = request.WithContext(ctx)
	}
	response, err := r.transport.RoundTrip(request)
	if err != nil {
		return
	}
	defer response.Body.Close()
	result = new(CloudRegionsListResponse)
	result.status = response.StatusCode
	result.header = response.Header
	if result.status >= 400 {
		result.err, err = errors.UnmarshalError(response.Body)
		if err != nil {
			return
		}
		err = result.err
		return
	}
	err = result.unmarshal(response.Body)
	if err != nil {
		return
	}
	return
}

// CloudRegionsListResponse is the response for the 'list' method.
type CloudRegionsListResponse struct {
	status int
	header http.Header
	err    *errors.Error
	page   *int
	size   *int
	total  *int
	items  *CloudRegionList
}

// Status returns the response status code.
func (r *CloudRegionsListResponse) Status() int {
	return r.status
}

// Header returns header of the response.
func (r *CloudRegionsListResponse) Header() http.Header {
	return r.header
}

// Error returns the response error.
func (r *CloudRegionsListResponse) Error() *errors.Error {
	return r.err
}

// Page returns the value of the 'page' parameter.
//
// Index of the returned page, where one corresponds to the first page. As this
// collection doesn't support paging the result will always be `1`.
func (r *CloudRegionsListResponse) Page() int {
	if r != nil && r.page != nil {
		return *r.page
	}
	return 0
}

// GetPage returns the value of the 'page' parameter and
// a flag indicating if the parameter has a value.
//
// Index of the returned page, where one corresponds to the first page. As this
// collection doesn't support paging the result will always be `1`.
func (r *CloudRegionsListResponse) GetPage() (value int, ok bool) {
	ok = r != nil && r.page != nil
	if ok {
		value = *r.page
	}
	return
}

// Size returns the value of the 'size' parameter.
//
// Number of items that will be contained in the returned page. As this collection
// doesn't support paging or searching the result will always be the total number of
// regions of the provider.
func (r *CloudRegionsListResponse) Size() int {
	if r != nil && r.size != nil {
		return *r.size
	}
	return 0
}

// GetSize returns the value of the 'size' parameter and
// a flag indicating if the parameter has a value.
//
// Number of items that will be contained in the returned page. As this collection
// doesn't support paging or searching the result will always be the total number of
// regions of the provider.
func (r *CloudRegionsListResponse) GetSize() (value int, ok bool) {
	ok = r != nil && r.size != nil
	if ok {
		value = *r.size
	}
	return
}

// Total returns the value of the 'total' parameter.
//
// Total number of items of the collection that match the search criteria,
// regardless of the size of the page. As this collection doesn't support paging or
// searching the result will always be the total number of regions of the provider.
func (r *CloudRegionsListResponse) Total() int {
	if r != nil && r.total != nil {
		return *r.total
	}
	return 0
}

// GetTotal returns the value of the 'total' parameter and
// a flag indicating if the parameter has a value.
//
// Total number of items of the collection that match the search criteria,
// regardless of the size of the page. As this collection doesn't support paging or
// searching the result will always be the total number of regions of the provider.
func (r *CloudRegionsListResponse) GetTotal() (value int, ok bool) {
	ok = r != nil && r.total != nil
	if ok {
		value = *r.total
	}
	return
}

// Items returns the value of the 'items' parameter.
//
// Retrieved list of cloud providers.
func (r *CloudRegionsListResponse) Items() *CloudRegionList {
	if r == nil {
		return nil
	}
	return r.items
}

// GetItems returns the value of the 'items' parameter and
// a flag indicating if the parameter has a value.
//
// Retrieved list of cloud providers.
func (r *CloudRegionsListResponse) GetItems() (value *CloudRegionList, ok bool) {
	ok = r != nil && r.items != nil
	if ok {
		value = r.items
	}
	return
}

// unmarshal is the method used internally to unmarshal responses to the
// 'list' method.
func (r *CloudRegionsListResponse) unmarshal(reader io.Reader) error {
	var err error
	decoder := json.NewDecoder(reader)
	data := new(cloudRegionsListResponseData)
	err = decoder.Decode(data)
	if err != nil {
		return err
	}
	r.page = data.Page
	r.size = data.Size
	r.total = data.Total
	r.items, err = data.Items.unwrap()
	if err != nil {
		return err
	}
	return err
}

// cloudRegionsListResponseData is the structure used internally to unmarshal
// the response of the 'list' method.
type cloudRegionsListResponseData struct {
	Page  *int                "json:\"page,omitempty\""
	Size  *int                "json:\"size,omitempty\""
	Total *int                "json:\"total,omitempty\""
	Items cloudRegionListData "json:\"items,omitempty\""
}