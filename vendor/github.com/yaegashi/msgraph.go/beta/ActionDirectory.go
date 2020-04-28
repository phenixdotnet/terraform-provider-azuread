// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/yaegashi/msgraph.go/jsonx"
)

// DirectoryObjectCollectionGetByIDsRequestParameter undocumented
type DirectoryObjectCollectionGetByIDsRequestParameter struct {
	// IDs undocumented
	IDs []string `json:"ids,omitempty"`
	// Types undocumented
	Types []string `json:"types,omitempty"`
}

// DirectoryObjectCollectionValidatePropertiesRequestParameter undocumented
type DirectoryObjectCollectionValidatePropertiesRequestParameter struct {
	// EntityType undocumented
	EntityType *string `json:"entityType,omitempty"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// MailNickname undocumented
	MailNickname *string `json:"mailNickname,omitempty"`
	// OnBehalfOfUserID undocumented
	OnBehalfOfUserID *UUID `json:"onBehalfOfUserId,omitempty"`
}

// DirectoryObjectCollectionGetUserOwnedObjectsRequestParameter undocumented
type DirectoryObjectCollectionGetUserOwnedObjectsRequestParameter struct {
	// UserID undocumented
	UserID *string `json:"userId,omitempty"`
	// Type undocumented
	Type *string `json:"type,omitempty"`
}

// DirectoryDefinitionDiscoverRequestParameter undocumented
type DirectoryDefinitionDiscoverRequestParameter struct {
}

// DirectoryObjectCheckMemberGroupsRequestParameter undocumented
type DirectoryObjectCheckMemberGroupsRequestParameter struct {
	// GroupIDs undocumented
	GroupIDs []string `json:"groupIds,omitempty"`
}

// DirectoryObjectCheckMemberObjectsRequestParameter undocumented
type DirectoryObjectCheckMemberObjectsRequestParameter struct {
	// IDs undocumented
	IDs []string `json:"ids,omitempty"`
}

// DirectoryObjectGetMemberGroupsRequestParameter undocumented
type DirectoryObjectGetMemberGroupsRequestParameter struct {
	// SecurityEnabledOnly undocumented
	SecurityEnabledOnly *bool `json:"securityEnabledOnly,omitempty"`
}

// DirectoryObjectGetMemberObjectsRequestParameter undocumented
type DirectoryObjectGetMemberObjectsRequestParameter struct {
	// SecurityEnabledOnly undocumented
	SecurityEnabledOnly *bool `json:"securityEnabledOnly,omitempty"`
}

// DirectoryObjectRestoreRequestParameter undocumented
type DirectoryObjectRestoreRequestParameter struct {
}

// DeletedItems returns request builder for DirectoryObject collection
func (b *DirectoryRequestBuilder) DeletedItems() *DirectoryDeletedItemsCollectionRequestBuilder {
	bb := &DirectoryDeletedItemsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/deletedItems"
	return bb
}

// DirectoryDeletedItemsCollectionRequestBuilder is request builder for DirectoryObject collection
type DirectoryDeletedItemsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for DirectoryObject collection
func (b *DirectoryDeletedItemsCollectionRequestBuilder) Request() *DirectoryDeletedItemsCollectionRequest {
	return &DirectoryDeletedItemsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for DirectoryObject item
func (b *DirectoryDeletedItemsCollectionRequestBuilder) ID(id string) *DirectoryObjectRequestBuilder {
	bb := &DirectoryObjectRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// DirectoryDeletedItemsCollectionRequest is request for DirectoryObject collection
type DirectoryDeletedItemsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for DirectoryObject collection
func (r *DirectoryDeletedItemsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]DirectoryObject, error) {
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
	var values []DirectoryObject
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
			value  []DirectoryObject
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

// GetN performs GET request for DirectoryObject collection, max N pages
func (r *DirectoryDeletedItemsCollectionRequest) GetN(ctx context.Context, n int) ([]DirectoryObject, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for DirectoryObject collection
func (r *DirectoryDeletedItemsCollectionRequest) Get(ctx context.Context) ([]DirectoryObject, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for DirectoryObject collection
func (r *DirectoryDeletedItemsCollectionRequest) Add(ctx context.Context, reqObj *DirectoryObject) (resObj *DirectoryObject, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// FeatureRolloutPolicies returns request builder for FeatureRolloutPolicy collection
func (b *DirectoryRequestBuilder) FeatureRolloutPolicies() *DirectoryFeatureRolloutPoliciesCollectionRequestBuilder {
	bb := &DirectoryFeatureRolloutPoliciesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/featureRolloutPolicies"
	return bb
}

// DirectoryFeatureRolloutPoliciesCollectionRequestBuilder is request builder for FeatureRolloutPolicy collection
type DirectoryFeatureRolloutPoliciesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for FeatureRolloutPolicy collection
func (b *DirectoryFeatureRolloutPoliciesCollectionRequestBuilder) Request() *DirectoryFeatureRolloutPoliciesCollectionRequest {
	return &DirectoryFeatureRolloutPoliciesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for FeatureRolloutPolicy item
func (b *DirectoryFeatureRolloutPoliciesCollectionRequestBuilder) ID(id string) *FeatureRolloutPolicyRequestBuilder {
	bb := &FeatureRolloutPolicyRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// DirectoryFeatureRolloutPoliciesCollectionRequest is request for FeatureRolloutPolicy collection
type DirectoryFeatureRolloutPoliciesCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for FeatureRolloutPolicy collection
func (r *DirectoryFeatureRolloutPoliciesCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]FeatureRolloutPolicy, error) {
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
	var values []FeatureRolloutPolicy
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
			value  []FeatureRolloutPolicy
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

// GetN performs GET request for FeatureRolloutPolicy collection, max N pages
func (r *DirectoryFeatureRolloutPoliciesCollectionRequest) GetN(ctx context.Context, n int) ([]FeatureRolloutPolicy, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for FeatureRolloutPolicy collection
func (r *DirectoryFeatureRolloutPoliciesCollectionRequest) Get(ctx context.Context) ([]FeatureRolloutPolicy, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for FeatureRolloutPolicy collection
func (r *DirectoryFeatureRolloutPoliciesCollectionRequest) Add(ctx context.Context, reqObj *FeatureRolloutPolicy) (resObj *FeatureRolloutPolicy, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// Members returns request builder for DirectoryObject collection
func (b *DirectoryRoleRequestBuilder) Members() *DirectoryRoleMembersCollectionRequestBuilder {
	bb := &DirectoryRoleMembersCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/members"
	return bb
}

// DirectoryRoleMembersCollectionRequestBuilder is request builder for DirectoryObject collection
type DirectoryRoleMembersCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for DirectoryObject collection
func (b *DirectoryRoleMembersCollectionRequestBuilder) Request() *DirectoryRoleMembersCollectionRequest {
	return &DirectoryRoleMembersCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for DirectoryObject item
func (b *DirectoryRoleMembersCollectionRequestBuilder) ID(id string) *DirectoryObjectRequestBuilder {
	bb := &DirectoryObjectRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// DirectoryRoleMembersCollectionRequest is request for DirectoryObject collection
type DirectoryRoleMembersCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for DirectoryObject collection
func (r *DirectoryRoleMembersCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]DirectoryObject, error) {
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
	var values []DirectoryObject
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
			value  []DirectoryObject
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

// GetN performs GET request for DirectoryObject collection, max N pages
func (r *DirectoryRoleMembersCollectionRequest) GetN(ctx context.Context, n int) ([]DirectoryObject, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for DirectoryObject collection
func (r *DirectoryRoleMembersCollectionRequest) Get(ctx context.Context) ([]DirectoryObject, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for DirectoryObject collection
func (r *DirectoryRoleMembersCollectionRequest) Add(ctx context.Context, reqObj *DirectoryObject) (resObj *DirectoryObject, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// ScopedMembers returns request builder for ScopedRoleMembership collection
func (b *DirectoryRoleRequestBuilder) ScopedMembers() *DirectoryRoleScopedMembersCollectionRequestBuilder {
	bb := &DirectoryRoleScopedMembersCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/scopedMembers"
	return bb
}

// DirectoryRoleScopedMembersCollectionRequestBuilder is request builder for ScopedRoleMembership collection
type DirectoryRoleScopedMembersCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for ScopedRoleMembership collection
func (b *DirectoryRoleScopedMembersCollectionRequestBuilder) Request() *DirectoryRoleScopedMembersCollectionRequest {
	return &DirectoryRoleScopedMembersCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for ScopedRoleMembership item
func (b *DirectoryRoleScopedMembersCollectionRequestBuilder) ID(id string) *ScopedRoleMembershipRequestBuilder {
	bb := &ScopedRoleMembershipRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// DirectoryRoleScopedMembersCollectionRequest is request for ScopedRoleMembership collection
type DirectoryRoleScopedMembersCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for ScopedRoleMembership collection
func (r *DirectoryRoleScopedMembersCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]ScopedRoleMembership, error) {
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
	var values []ScopedRoleMembership
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
			value  []ScopedRoleMembership
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

// GetN performs GET request for ScopedRoleMembership collection, max N pages
func (r *DirectoryRoleScopedMembersCollectionRequest) GetN(ctx context.Context, n int) ([]ScopedRoleMembership, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for ScopedRoleMembership collection
func (r *DirectoryRoleScopedMembersCollectionRequest) Get(ctx context.Context) ([]ScopedRoleMembership, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for ScopedRoleMembership collection
func (r *DirectoryRoleScopedMembersCollectionRequest) Add(ctx context.Context, reqObj *ScopedRoleMembership) (resObj *ScopedRoleMembership, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}
