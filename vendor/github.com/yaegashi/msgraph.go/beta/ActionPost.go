// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/yaegashi/msgraph.go/jsonx"
)

// PostReplyRequestParameter undocumented
type PostReplyRequestParameter struct {
	// Post undocumented
	Post *Post `json:"Post,omitempty"`
}

// PostForwardRequestParameter undocumented
type PostForwardRequestParameter struct {
	// Comment undocumented
	Comment *string `json:"Comment,omitempty"`
	// ToRecipients undocumented
	ToRecipients []Recipient `json:"ToRecipients,omitempty"`
}

// Attachments returns request builder for Attachment collection
func (b *PostRequestBuilder) Attachments() *PostAttachmentsCollectionRequestBuilder {
	bb := &PostAttachmentsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/attachments"
	return bb
}

// PostAttachmentsCollectionRequestBuilder is request builder for Attachment collection
type PostAttachmentsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for Attachment collection
func (b *PostAttachmentsCollectionRequestBuilder) Request() *PostAttachmentsCollectionRequest {
	return &PostAttachmentsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for Attachment item
func (b *PostAttachmentsCollectionRequestBuilder) ID(id string) *AttachmentRequestBuilder {
	bb := &AttachmentRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// PostAttachmentsCollectionRequest is request for Attachment collection
type PostAttachmentsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for Attachment collection
func (r *PostAttachmentsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]Attachment, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []Attachment
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			errRes := &ErrorResponse{Response: res}
			err := jsonx.Unmarshal(b, errRes)
			if err != nil {
				return nil, fmt.Errorf("%s: %s", res.Status, string(b))
			}
			return nil, errRes
		}
		var (
			paging Paging
			value  []Attachment
		)
		err := jsonx.NewDecoder(res.Body).Decode(&paging)
		if err != nil {
			return nil, err
		}
		err = jsonx.Unmarshal(paging.Value, &value)
		if err != nil {
			return nil, err
		}
		values = append(values, value...)
		if n >= 0 {
			n--
		}
		if n == 0 || len(paging.NextLink) == 0 {
			return values, nil
		}
		req, err = http.NewRequest("GET", paging.NextLink, nil)
		if ctx != nil {
			req = req.WithContext(ctx)
		}
		res, err = r.client.Do(req)
		if err != nil {
			return nil, err
		}
	}
}

// GetN performs GET request for Attachment collection, max N pages
func (r *PostAttachmentsCollectionRequest) GetN(ctx context.Context, n int) ([]Attachment, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for Attachment collection
func (r *PostAttachmentsCollectionRequest) Get(ctx context.Context) ([]Attachment, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for Attachment collection
func (r *PostAttachmentsCollectionRequest) Add(ctx context.Context, reqObj *Attachment) (resObj *Attachment, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// Extensions returns request builder for Extension collection
func (b *PostRequestBuilder) Extensions() *PostExtensionsCollectionRequestBuilder {
	bb := &PostExtensionsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/extensions"
	return bb
}

// PostExtensionsCollectionRequestBuilder is request builder for Extension collection
type PostExtensionsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for Extension collection
func (b *PostExtensionsCollectionRequestBuilder) Request() *PostExtensionsCollectionRequest {
	return &PostExtensionsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for Extension item
func (b *PostExtensionsCollectionRequestBuilder) ID(id string) *ExtensionRequestBuilder {
	bb := &ExtensionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// PostExtensionsCollectionRequest is request for Extension collection
type PostExtensionsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for Extension collection
func (r *PostExtensionsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]Extension, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []Extension
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			errRes := &ErrorResponse{Response: res}
			err := jsonx.Unmarshal(b, errRes)
			if err != nil {
				return nil, fmt.Errorf("%s: %s", res.Status, string(b))
			}
			return nil, errRes
		}
		var (
			paging Paging
			value  []Extension
		)
		err := jsonx.NewDecoder(res.Body).Decode(&paging)
		if err != nil {
			return nil, err
		}
		err = jsonx.Unmarshal(paging.Value, &value)
		if err != nil {
			return nil, err
		}
		values = append(values, value...)
		if n >= 0 {
			n--
		}
		if n == 0 || len(paging.NextLink) == 0 {
			return values, nil
		}
		req, err = http.NewRequest("GET", paging.NextLink, nil)
		if ctx != nil {
			req = req.WithContext(ctx)
		}
		res, err = r.client.Do(req)
		if err != nil {
			return nil, err
		}
	}
}

// GetN performs GET request for Extension collection, max N pages
func (r *PostExtensionsCollectionRequest) GetN(ctx context.Context, n int) ([]Extension, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for Extension collection
func (r *PostExtensionsCollectionRequest) Get(ctx context.Context) ([]Extension, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for Extension collection
func (r *PostExtensionsCollectionRequest) Add(ctx context.Context, reqObj *Extension) (resObj *Extension, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// InReplyTo is navigation property
func (b *PostRequestBuilder) InReplyTo() *PostRequestBuilder {
	bb := &PostRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/inReplyTo"
	return bb
}

// Mentions returns request builder for Mention collection
func (b *PostRequestBuilder) Mentions() *PostMentionsCollectionRequestBuilder {
	bb := &PostMentionsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/mentions"
	return bb
}

// PostMentionsCollectionRequestBuilder is request builder for Mention collection
type PostMentionsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for Mention collection
func (b *PostMentionsCollectionRequestBuilder) Request() *PostMentionsCollectionRequest {
	return &PostMentionsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for Mention item
func (b *PostMentionsCollectionRequestBuilder) ID(id string) *MentionRequestBuilder {
	bb := &MentionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// PostMentionsCollectionRequest is request for Mention collection
type PostMentionsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for Mention collection
func (r *PostMentionsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]Mention, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []Mention
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			errRes := &ErrorResponse{Response: res}
			err := jsonx.Unmarshal(b, errRes)
			if err != nil {
				return nil, fmt.Errorf("%s: %s", res.Status, string(b))
			}
			return nil, errRes
		}
		var (
			paging Paging
			value  []Mention
		)
		err := jsonx.NewDecoder(res.Body).Decode(&paging)
		if err != nil {
			return nil, err
		}
		err = jsonx.Unmarshal(paging.Value, &value)
		if err != nil {
			return nil, err
		}
		values = append(values, value...)
		if n >= 0 {
			n--
		}
		if n == 0 || len(paging.NextLink) == 0 {
			return values, nil
		}
		req, err = http.NewRequest("GET", paging.NextLink, nil)
		if ctx != nil {
			req = req.WithContext(ctx)
		}
		res, err = r.client.Do(req)
		if err != nil {
			return nil, err
		}
	}
}

// GetN performs GET request for Mention collection, max N pages
func (r *PostMentionsCollectionRequest) GetN(ctx context.Context, n int) ([]Mention, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for Mention collection
func (r *PostMentionsCollectionRequest) Get(ctx context.Context) ([]Mention, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for Mention collection
func (r *PostMentionsCollectionRequest) Add(ctx context.Context, reqObj *Mention) (resObj *Mention, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// MultiValueExtendedProperties returns request builder for MultiValueLegacyExtendedProperty collection
func (b *PostRequestBuilder) MultiValueExtendedProperties() *PostMultiValueExtendedPropertiesCollectionRequestBuilder {
	bb := &PostMultiValueExtendedPropertiesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/multiValueExtendedProperties"
	return bb
}

// PostMultiValueExtendedPropertiesCollectionRequestBuilder is request builder for MultiValueLegacyExtendedProperty collection
type PostMultiValueExtendedPropertiesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for MultiValueLegacyExtendedProperty collection
func (b *PostMultiValueExtendedPropertiesCollectionRequestBuilder) Request() *PostMultiValueExtendedPropertiesCollectionRequest {
	return &PostMultiValueExtendedPropertiesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for MultiValueLegacyExtendedProperty item
func (b *PostMultiValueExtendedPropertiesCollectionRequestBuilder) ID(id string) *MultiValueLegacyExtendedPropertyRequestBuilder {
	bb := &MultiValueLegacyExtendedPropertyRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// PostMultiValueExtendedPropertiesCollectionRequest is request for MultiValueLegacyExtendedProperty collection
type PostMultiValueExtendedPropertiesCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for MultiValueLegacyExtendedProperty collection
func (r *PostMultiValueExtendedPropertiesCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]MultiValueLegacyExtendedProperty, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []MultiValueLegacyExtendedProperty
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			errRes := &ErrorResponse{Response: res}
			err := jsonx.Unmarshal(b, errRes)
			if err != nil {
				return nil, fmt.Errorf("%s: %s", res.Status, string(b))
			}
			return nil, errRes
		}
		var (
			paging Paging
			value  []MultiValueLegacyExtendedProperty
		)
		err := jsonx.NewDecoder(res.Body).Decode(&paging)
		if err != nil {
			return nil, err
		}
		err = jsonx.Unmarshal(paging.Value, &value)
		if err != nil {
			return nil, err
		}
		values = append(values, value...)
		if n >= 0 {
			n--
		}
		if n == 0 || len(paging.NextLink) == 0 {
			return values, nil
		}
		req, err = http.NewRequest("GET", paging.NextLink, nil)
		if ctx != nil {
			req = req.WithContext(ctx)
		}
		res, err = r.client.Do(req)
		if err != nil {
			return nil, err
		}
	}
}

// GetN performs GET request for MultiValueLegacyExtendedProperty collection, max N pages
func (r *PostMultiValueExtendedPropertiesCollectionRequest) GetN(ctx context.Context, n int) ([]MultiValueLegacyExtendedProperty, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for MultiValueLegacyExtendedProperty collection
func (r *PostMultiValueExtendedPropertiesCollectionRequest) Get(ctx context.Context) ([]MultiValueLegacyExtendedProperty, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for MultiValueLegacyExtendedProperty collection
func (r *PostMultiValueExtendedPropertiesCollectionRequest) Add(ctx context.Context, reqObj *MultiValueLegacyExtendedProperty) (resObj *MultiValueLegacyExtendedProperty, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// SingleValueExtendedProperties returns request builder for SingleValueLegacyExtendedProperty collection
func (b *PostRequestBuilder) SingleValueExtendedProperties() *PostSingleValueExtendedPropertiesCollectionRequestBuilder {
	bb := &PostSingleValueExtendedPropertiesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/singleValueExtendedProperties"
	return bb
}

// PostSingleValueExtendedPropertiesCollectionRequestBuilder is request builder for SingleValueLegacyExtendedProperty collection
type PostSingleValueExtendedPropertiesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for SingleValueLegacyExtendedProperty collection
func (b *PostSingleValueExtendedPropertiesCollectionRequestBuilder) Request() *PostSingleValueExtendedPropertiesCollectionRequest {
	return &PostSingleValueExtendedPropertiesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for SingleValueLegacyExtendedProperty item
func (b *PostSingleValueExtendedPropertiesCollectionRequestBuilder) ID(id string) *SingleValueLegacyExtendedPropertyRequestBuilder {
	bb := &SingleValueLegacyExtendedPropertyRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// PostSingleValueExtendedPropertiesCollectionRequest is request for SingleValueLegacyExtendedProperty collection
type PostSingleValueExtendedPropertiesCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for SingleValueLegacyExtendedProperty collection
func (r *PostSingleValueExtendedPropertiesCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]SingleValueLegacyExtendedProperty, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []SingleValueLegacyExtendedProperty
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			errRes := &ErrorResponse{Response: res}
			err := jsonx.Unmarshal(b, errRes)
			if err != nil {
				return nil, fmt.Errorf("%s: %s", res.Status, string(b))
			}
			return nil, errRes
		}
		var (
			paging Paging
			value  []SingleValueLegacyExtendedProperty
		)
		err := jsonx.NewDecoder(res.Body).Decode(&paging)
		if err != nil {
			return nil, err
		}
		err = jsonx.Unmarshal(paging.Value, &value)
		if err != nil {
			return nil, err
		}
		values = append(values, value...)
		if n >= 0 {
			n--
		}
		if n == 0 || len(paging.NextLink) == 0 {
			return values, nil
		}
		req, err = http.NewRequest("GET", paging.NextLink, nil)
		if ctx != nil {
			req = req.WithContext(ctx)
		}
		res, err = r.client.Do(req)
		if err != nil {
			return nil, err
		}
	}
}

// GetN performs GET request for SingleValueLegacyExtendedProperty collection, max N pages
func (r *PostSingleValueExtendedPropertiesCollectionRequest) GetN(ctx context.Context, n int) ([]SingleValueLegacyExtendedProperty, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for SingleValueLegacyExtendedProperty collection
func (r *PostSingleValueExtendedPropertiesCollectionRequest) Get(ctx context.Context) ([]SingleValueLegacyExtendedProperty, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for SingleValueLegacyExtendedProperty collection
func (r *PostSingleValueExtendedPropertiesCollectionRequest) Add(ctx context.Context, reqObj *SingleValueLegacyExtendedProperty) (resObj *SingleValueLegacyExtendedProperty, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}
