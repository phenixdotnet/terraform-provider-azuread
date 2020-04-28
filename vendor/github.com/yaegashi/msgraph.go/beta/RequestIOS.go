// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/yaegashi/msgraph.go/jsonx"
)

// IOSCertificateProfileRequestBuilder is request builder for IOSCertificateProfile
type IOSCertificateProfileRequestBuilder struct{ BaseRequestBuilder }

// Request returns IOSCertificateProfileRequest
func (b *IOSCertificateProfileRequestBuilder) Request() *IOSCertificateProfileRequest {
	return &IOSCertificateProfileRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// IOSCertificateProfileRequest is request for IOSCertificateProfile
type IOSCertificateProfileRequest struct{ BaseRequest }

// Get performs GET request for IOSCertificateProfile
func (r *IOSCertificateProfileRequest) Get(ctx context.Context) (resObj *IOSCertificateProfile, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for IOSCertificateProfile
func (r *IOSCertificateProfileRequest) Update(ctx context.Context, reqObj *IOSCertificateProfile) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for IOSCertificateProfile
func (r *IOSCertificateProfileRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// IOSCertificateProfileBaseRequestBuilder is request builder for IOSCertificateProfileBase
type IOSCertificateProfileBaseRequestBuilder struct{ BaseRequestBuilder }

// Request returns IOSCertificateProfileBaseRequest
func (b *IOSCertificateProfileBaseRequestBuilder) Request() *IOSCertificateProfileBaseRequest {
	return &IOSCertificateProfileBaseRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// IOSCertificateProfileBaseRequest is request for IOSCertificateProfileBase
type IOSCertificateProfileBaseRequest struct{ BaseRequest }

// Get performs GET request for IOSCertificateProfileBase
func (r *IOSCertificateProfileBaseRequest) Get(ctx context.Context) (resObj *IOSCertificateProfileBase, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for IOSCertificateProfileBase
func (r *IOSCertificateProfileBaseRequest) Update(ctx context.Context, reqObj *IOSCertificateProfileBase) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for IOSCertificateProfileBase
func (r *IOSCertificateProfileBaseRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// IOSDerivedCredentialAuthenticationConfigurationRequestBuilder is request builder for IOSDerivedCredentialAuthenticationConfiguration
type IOSDerivedCredentialAuthenticationConfigurationRequestBuilder struct{ BaseRequestBuilder }

// Request returns IOSDerivedCredentialAuthenticationConfigurationRequest
func (b *IOSDerivedCredentialAuthenticationConfigurationRequestBuilder) Request() *IOSDerivedCredentialAuthenticationConfigurationRequest {
	return &IOSDerivedCredentialAuthenticationConfigurationRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// IOSDerivedCredentialAuthenticationConfigurationRequest is request for IOSDerivedCredentialAuthenticationConfiguration
type IOSDerivedCredentialAuthenticationConfigurationRequest struct{ BaseRequest }

// Get performs GET request for IOSDerivedCredentialAuthenticationConfiguration
func (r *IOSDerivedCredentialAuthenticationConfigurationRequest) Get(ctx context.Context) (resObj *IOSDerivedCredentialAuthenticationConfiguration, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for IOSDerivedCredentialAuthenticationConfiguration
func (r *IOSDerivedCredentialAuthenticationConfigurationRequest) Update(ctx context.Context, reqObj *IOSDerivedCredentialAuthenticationConfiguration) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for IOSDerivedCredentialAuthenticationConfiguration
func (r *IOSDerivedCredentialAuthenticationConfigurationRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// IOSDeviceFeaturesConfigurationRequestBuilder is request builder for IOSDeviceFeaturesConfiguration
type IOSDeviceFeaturesConfigurationRequestBuilder struct{ BaseRequestBuilder }

// Request returns IOSDeviceFeaturesConfigurationRequest
func (b *IOSDeviceFeaturesConfigurationRequestBuilder) Request() *IOSDeviceFeaturesConfigurationRequest {
	return &IOSDeviceFeaturesConfigurationRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// IOSDeviceFeaturesConfigurationRequest is request for IOSDeviceFeaturesConfiguration
type IOSDeviceFeaturesConfigurationRequest struct{ BaseRequest }

// Get performs GET request for IOSDeviceFeaturesConfiguration
func (r *IOSDeviceFeaturesConfigurationRequest) Get(ctx context.Context) (resObj *IOSDeviceFeaturesConfiguration, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for IOSDeviceFeaturesConfiguration
func (r *IOSDeviceFeaturesConfigurationRequest) Update(ctx context.Context, reqObj *IOSDeviceFeaturesConfiguration) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for IOSDeviceFeaturesConfiguration
func (r *IOSDeviceFeaturesConfigurationRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// IOSEasEmailProfileConfigurationRequestBuilder is request builder for IOSEasEmailProfileConfiguration
type IOSEasEmailProfileConfigurationRequestBuilder struct{ BaseRequestBuilder }

// Request returns IOSEasEmailProfileConfigurationRequest
func (b *IOSEasEmailProfileConfigurationRequestBuilder) Request() *IOSEasEmailProfileConfigurationRequest {
	return &IOSEasEmailProfileConfigurationRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// IOSEasEmailProfileConfigurationRequest is request for IOSEasEmailProfileConfiguration
type IOSEasEmailProfileConfigurationRequest struct{ BaseRequest }

// Get performs GET request for IOSEasEmailProfileConfiguration
func (r *IOSEasEmailProfileConfigurationRequest) Get(ctx context.Context) (resObj *IOSEasEmailProfileConfiguration, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for IOSEasEmailProfileConfiguration
func (r *IOSEasEmailProfileConfigurationRequest) Update(ctx context.Context, reqObj *IOSEasEmailProfileConfiguration) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for IOSEasEmailProfileConfiguration
func (r *IOSEasEmailProfileConfigurationRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// IOSEnterpriseWiFiConfigurationRequestBuilder is request builder for IOSEnterpriseWiFiConfiguration
type IOSEnterpriseWiFiConfigurationRequestBuilder struct{ BaseRequestBuilder }

// Request returns IOSEnterpriseWiFiConfigurationRequest
func (b *IOSEnterpriseWiFiConfigurationRequestBuilder) Request() *IOSEnterpriseWiFiConfigurationRequest {
	return &IOSEnterpriseWiFiConfigurationRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// IOSEnterpriseWiFiConfigurationRequest is request for IOSEnterpriseWiFiConfiguration
type IOSEnterpriseWiFiConfigurationRequest struct{ BaseRequest }

// Get performs GET request for IOSEnterpriseWiFiConfiguration
func (r *IOSEnterpriseWiFiConfigurationRequest) Get(ctx context.Context) (resObj *IOSEnterpriseWiFiConfiguration, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for IOSEnterpriseWiFiConfiguration
func (r *IOSEnterpriseWiFiConfigurationRequest) Update(ctx context.Context, reqObj *IOSEnterpriseWiFiConfiguration) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for IOSEnterpriseWiFiConfiguration
func (r *IOSEnterpriseWiFiConfigurationRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// IOSImportedPFXCertificateProfileRequestBuilder is request builder for IOSImportedPFXCertificateProfile
type IOSImportedPFXCertificateProfileRequestBuilder struct{ BaseRequestBuilder }

// Request returns IOSImportedPFXCertificateProfileRequest
func (b *IOSImportedPFXCertificateProfileRequestBuilder) Request() *IOSImportedPFXCertificateProfileRequest {
	return &IOSImportedPFXCertificateProfileRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// IOSImportedPFXCertificateProfileRequest is request for IOSImportedPFXCertificateProfile
type IOSImportedPFXCertificateProfileRequest struct{ BaseRequest }

// Get performs GET request for IOSImportedPFXCertificateProfile
func (r *IOSImportedPFXCertificateProfileRequest) Get(ctx context.Context) (resObj *IOSImportedPFXCertificateProfile, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for IOSImportedPFXCertificateProfile
func (r *IOSImportedPFXCertificateProfileRequest) Update(ctx context.Context, reqObj *IOSImportedPFXCertificateProfile) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for IOSImportedPFXCertificateProfile
func (r *IOSImportedPFXCertificateProfileRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// IOSLobAppProvisioningConfigurationRequestBuilder is request builder for IOSLobAppProvisioningConfiguration
type IOSLobAppProvisioningConfigurationRequestBuilder struct{ BaseRequestBuilder }

// Request returns IOSLobAppProvisioningConfigurationRequest
func (b *IOSLobAppProvisioningConfigurationRequestBuilder) Request() *IOSLobAppProvisioningConfigurationRequest {
	return &IOSLobAppProvisioningConfigurationRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// IOSLobAppProvisioningConfigurationRequest is request for IOSLobAppProvisioningConfiguration
type IOSLobAppProvisioningConfigurationRequest struct{ BaseRequest }

// Get performs GET request for IOSLobAppProvisioningConfiguration
func (r *IOSLobAppProvisioningConfigurationRequest) Get(ctx context.Context) (resObj *IOSLobAppProvisioningConfiguration, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for IOSLobAppProvisioningConfiguration
func (r *IOSLobAppProvisioningConfigurationRequest) Update(ctx context.Context, reqObj *IOSLobAppProvisioningConfiguration) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for IOSLobAppProvisioningConfiguration
func (r *IOSLobAppProvisioningConfigurationRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// IOSLobAppProvisioningConfigurationAssignmentRequestBuilder is request builder for IOSLobAppProvisioningConfigurationAssignment
type IOSLobAppProvisioningConfigurationAssignmentRequestBuilder struct{ BaseRequestBuilder }

// Request returns IOSLobAppProvisioningConfigurationAssignmentRequest
func (b *IOSLobAppProvisioningConfigurationAssignmentRequestBuilder) Request() *IOSLobAppProvisioningConfigurationAssignmentRequest {
	return &IOSLobAppProvisioningConfigurationAssignmentRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// IOSLobAppProvisioningConfigurationAssignmentRequest is request for IOSLobAppProvisioningConfigurationAssignment
type IOSLobAppProvisioningConfigurationAssignmentRequest struct{ BaseRequest }

// Get performs GET request for IOSLobAppProvisioningConfigurationAssignment
func (r *IOSLobAppProvisioningConfigurationAssignmentRequest) Get(ctx context.Context) (resObj *IOSLobAppProvisioningConfigurationAssignment, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for IOSLobAppProvisioningConfigurationAssignment
func (r *IOSLobAppProvisioningConfigurationAssignmentRequest) Update(ctx context.Context, reqObj *IOSLobAppProvisioningConfigurationAssignment) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for IOSLobAppProvisioningConfigurationAssignment
func (r *IOSLobAppProvisioningConfigurationAssignmentRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// IOSManagedAppProtectionRequestBuilder is request builder for IOSManagedAppProtection
type IOSManagedAppProtectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns IOSManagedAppProtectionRequest
func (b *IOSManagedAppProtectionRequestBuilder) Request() *IOSManagedAppProtectionRequest {
	return &IOSManagedAppProtectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// IOSManagedAppProtectionRequest is request for IOSManagedAppProtection
type IOSManagedAppProtectionRequest struct{ BaseRequest }

// Get performs GET request for IOSManagedAppProtection
func (r *IOSManagedAppProtectionRequest) Get(ctx context.Context) (resObj *IOSManagedAppProtection, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for IOSManagedAppProtection
func (r *IOSManagedAppProtectionRequest) Update(ctx context.Context, reqObj *IOSManagedAppProtection) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for IOSManagedAppProtection
func (r *IOSManagedAppProtectionRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// IOSPkcsCertificateProfileRequestBuilder is request builder for IOSPkcsCertificateProfile
type IOSPkcsCertificateProfileRequestBuilder struct{ BaseRequestBuilder }

// Request returns IOSPkcsCertificateProfileRequest
func (b *IOSPkcsCertificateProfileRequestBuilder) Request() *IOSPkcsCertificateProfileRequest {
	return &IOSPkcsCertificateProfileRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// IOSPkcsCertificateProfileRequest is request for IOSPkcsCertificateProfile
type IOSPkcsCertificateProfileRequest struct{ BaseRequest }

// Get performs GET request for IOSPkcsCertificateProfile
func (r *IOSPkcsCertificateProfileRequest) Get(ctx context.Context) (resObj *IOSPkcsCertificateProfile, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for IOSPkcsCertificateProfile
func (r *IOSPkcsCertificateProfileRequest) Update(ctx context.Context, reqObj *IOSPkcsCertificateProfile) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for IOSPkcsCertificateProfile
func (r *IOSPkcsCertificateProfileRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// IOSScepCertificateProfileRequestBuilder is request builder for IOSScepCertificateProfile
type IOSScepCertificateProfileRequestBuilder struct{ BaseRequestBuilder }

// Request returns IOSScepCertificateProfileRequest
func (b *IOSScepCertificateProfileRequestBuilder) Request() *IOSScepCertificateProfileRequest {
	return &IOSScepCertificateProfileRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// IOSScepCertificateProfileRequest is request for IOSScepCertificateProfile
type IOSScepCertificateProfileRequest struct{ BaseRequest }

// Get performs GET request for IOSScepCertificateProfile
func (r *IOSScepCertificateProfileRequest) Get(ctx context.Context) (resObj *IOSScepCertificateProfile, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for IOSScepCertificateProfile
func (r *IOSScepCertificateProfileRequest) Update(ctx context.Context, reqObj *IOSScepCertificateProfile) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for IOSScepCertificateProfile
func (r *IOSScepCertificateProfileRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// IOSTrustedRootCertificateRequestBuilder is request builder for IOSTrustedRootCertificate
type IOSTrustedRootCertificateRequestBuilder struct{ BaseRequestBuilder }

// Request returns IOSTrustedRootCertificateRequest
func (b *IOSTrustedRootCertificateRequestBuilder) Request() *IOSTrustedRootCertificateRequest {
	return &IOSTrustedRootCertificateRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// IOSTrustedRootCertificateRequest is request for IOSTrustedRootCertificate
type IOSTrustedRootCertificateRequest struct{ BaseRequest }

// Get performs GET request for IOSTrustedRootCertificate
func (r *IOSTrustedRootCertificateRequest) Get(ctx context.Context) (resObj *IOSTrustedRootCertificate, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for IOSTrustedRootCertificate
func (r *IOSTrustedRootCertificateRequest) Update(ctx context.Context, reqObj *IOSTrustedRootCertificate) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for IOSTrustedRootCertificate
func (r *IOSTrustedRootCertificateRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// IOSUpdateDeviceStatusRequestBuilder is request builder for IOSUpdateDeviceStatus
type IOSUpdateDeviceStatusRequestBuilder struct{ BaseRequestBuilder }

// Request returns IOSUpdateDeviceStatusRequest
func (b *IOSUpdateDeviceStatusRequestBuilder) Request() *IOSUpdateDeviceStatusRequest {
	return &IOSUpdateDeviceStatusRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// IOSUpdateDeviceStatusRequest is request for IOSUpdateDeviceStatus
type IOSUpdateDeviceStatusRequest struct{ BaseRequest }

// Get performs GET request for IOSUpdateDeviceStatus
func (r *IOSUpdateDeviceStatusRequest) Get(ctx context.Context) (resObj *IOSUpdateDeviceStatus, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for IOSUpdateDeviceStatus
func (r *IOSUpdateDeviceStatusRequest) Update(ctx context.Context, reqObj *IOSUpdateDeviceStatus) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for IOSUpdateDeviceStatus
func (r *IOSUpdateDeviceStatusRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// IOSVPNConfigurationRequestBuilder is request builder for IOSVPNConfiguration
type IOSVPNConfigurationRequestBuilder struct{ BaseRequestBuilder }

// Request returns IOSVPNConfigurationRequest
func (b *IOSVPNConfigurationRequestBuilder) Request() *IOSVPNConfigurationRequest {
	return &IOSVPNConfigurationRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// IOSVPNConfigurationRequest is request for IOSVPNConfiguration
type IOSVPNConfigurationRequest struct{ BaseRequest }

// Get performs GET request for IOSVPNConfiguration
func (r *IOSVPNConfigurationRequest) Get(ctx context.Context) (resObj *IOSVPNConfiguration, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for IOSVPNConfiguration
func (r *IOSVPNConfigurationRequest) Update(ctx context.Context, reqObj *IOSVPNConfiguration) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for IOSVPNConfiguration
func (r *IOSVPNConfigurationRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// IOSVPPAppRequestBuilder is request builder for IOSVPPApp
type IOSVPPAppRequestBuilder struct{ BaseRequestBuilder }

// Request returns IOSVPPAppRequest
func (b *IOSVPPAppRequestBuilder) Request() *IOSVPPAppRequest {
	return &IOSVPPAppRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// IOSVPPAppRequest is request for IOSVPPApp
type IOSVPPAppRequest struct{ BaseRequest }

// Get performs GET request for IOSVPPApp
func (r *IOSVPPAppRequest) Get(ctx context.Context) (resObj *IOSVPPApp, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for IOSVPPApp
func (r *IOSVPPAppRequest) Update(ctx context.Context, reqObj *IOSVPPApp) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for IOSVPPApp
func (r *IOSVPPAppRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// IOSVPPAppAssignedLicenseRequestBuilder is request builder for IOSVPPAppAssignedLicense
type IOSVPPAppAssignedLicenseRequestBuilder struct{ BaseRequestBuilder }

// Request returns IOSVPPAppAssignedLicenseRequest
func (b *IOSVPPAppAssignedLicenseRequestBuilder) Request() *IOSVPPAppAssignedLicenseRequest {
	return &IOSVPPAppAssignedLicenseRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// IOSVPPAppAssignedLicenseRequest is request for IOSVPPAppAssignedLicense
type IOSVPPAppAssignedLicenseRequest struct{ BaseRequest }

// Get performs GET request for IOSVPPAppAssignedLicense
func (r *IOSVPPAppAssignedLicenseRequest) Get(ctx context.Context) (resObj *IOSVPPAppAssignedLicense, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for IOSVPPAppAssignedLicense
func (r *IOSVPPAppAssignedLicenseRequest) Update(ctx context.Context, reqObj *IOSVPPAppAssignedLicense) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for IOSVPPAppAssignedLicense
func (r *IOSVPPAppAssignedLicenseRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

//
type IOSLobAppProvisioningConfigurationCollectionHasPayloadLinksRequestBuilder struct{ BaseRequestBuilder }

// HasPayloadLinks action undocumented
func (b *DeviceAppManagementIOSLobAppProvisioningConfigurationsCollectionRequestBuilder) HasPayloadLinks(reqObj *IOSLobAppProvisioningConfigurationCollectionHasPayloadLinksRequestParameter) *IOSLobAppProvisioningConfigurationCollectionHasPayloadLinksRequestBuilder {
	bb := &IOSLobAppProvisioningConfigurationCollectionHasPayloadLinksRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/hasPayloadLinks"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type IOSLobAppProvisioningConfigurationCollectionHasPayloadLinksRequest struct{ BaseRequest }

//
func (b *IOSLobAppProvisioningConfigurationCollectionHasPayloadLinksRequestBuilder) Request() *IOSLobAppProvisioningConfigurationCollectionHasPayloadLinksRequest {
	return &IOSLobAppProvisioningConfigurationCollectionHasPayloadLinksRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *IOSLobAppProvisioningConfigurationCollectionHasPayloadLinksRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]HasPayloadLinkResultItem, error) {
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
	var values []HasPayloadLinkResultItem
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
			value  []HasPayloadLinkResultItem
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

//
func (r *IOSLobAppProvisioningConfigurationCollectionHasPayloadLinksRequest) PostN(ctx context.Context, n int) ([]HasPayloadLinkResultItem, error) {
	return r.Paging(ctx, "POST", "", r.requestObject, n)
}

//
func (r *IOSLobAppProvisioningConfigurationCollectionHasPayloadLinksRequest) Post(ctx context.Context) ([]HasPayloadLinkResultItem, error) {
	return r.Paging(ctx, "POST", "", r.requestObject, 0)
}

//
type IOSManagedAppProtectionCollectionHasPayloadLinksRequestBuilder struct{ BaseRequestBuilder }

// HasPayloadLinks action undocumented
func (b *DeviceAppManagementIOSManagedAppProtectionsCollectionRequestBuilder) HasPayloadLinks(reqObj *IOSManagedAppProtectionCollectionHasPayloadLinksRequestParameter) *IOSManagedAppProtectionCollectionHasPayloadLinksRequestBuilder {
	bb := &IOSManagedAppProtectionCollectionHasPayloadLinksRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/hasPayloadLinks"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type IOSManagedAppProtectionCollectionHasPayloadLinksRequest struct{ BaseRequest }

//
func (b *IOSManagedAppProtectionCollectionHasPayloadLinksRequestBuilder) Request() *IOSManagedAppProtectionCollectionHasPayloadLinksRequest {
	return &IOSManagedAppProtectionCollectionHasPayloadLinksRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *IOSManagedAppProtectionCollectionHasPayloadLinksRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]HasPayloadLinkResultItem, error) {
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
	var values []HasPayloadLinkResultItem
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
			value  []HasPayloadLinkResultItem
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

//
func (r *IOSManagedAppProtectionCollectionHasPayloadLinksRequest) PostN(ctx context.Context, n int) ([]HasPayloadLinkResultItem, error) {
	return r.Paging(ctx, "POST", "", r.requestObject, n)
}

//
func (r *IOSManagedAppProtectionCollectionHasPayloadLinksRequest) Post(ctx context.Context) ([]HasPayloadLinkResultItem, error) {
	return r.Paging(ctx, "POST", "", r.requestObject, 0)
}

//
type IOSLobAppProvisioningConfigurationAssignRequestBuilder struct{ BaseRequestBuilder }

// Assign action undocumented
func (b *IOSLobAppProvisioningConfigurationRequestBuilder) Assign(reqObj *IOSLobAppProvisioningConfigurationAssignRequestParameter) *IOSLobAppProvisioningConfigurationAssignRequestBuilder {
	bb := &IOSLobAppProvisioningConfigurationAssignRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/assign"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type IOSLobAppProvisioningConfigurationAssignRequest struct{ BaseRequest }

//
func (b *IOSLobAppProvisioningConfigurationAssignRequestBuilder) Request() *IOSLobAppProvisioningConfigurationAssignRequest {
	return &IOSLobAppProvisioningConfigurationAssignRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *IOSLobAppProvisioningConfigurationAssignRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}

//
type IOSVPPAppRevokeAllLicensesRequestBuilder struct{ BaseRequestBuilder }

// RevokeAllLicenses action undocumented
func (b *IOSVPPAppRequestBuilder) RevokeAllLicenses(reqObj *IOSVPPAppRevokeAllLicensesRequestParameter) *IOSVPPAppRevokeAllLicensesRequestBuilder {
	bb := &IOSVPPAppRevokeAllLicensesRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/revokeAllLicenses"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type IOSVPPAppRevokeAllLicensesRequest struct{ BaseRequest }

//
func (b *IOSVPPAppRevokeAllLicensesRequestBuilder) Request() *IOSVPPAppRevokeAllLicensesRequest {
	return &IOSVPPAppRevokeAllLicensesRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *IOSVPPAppRevokeAllLicensesRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}

//
type IOSVPPAppRevokeUserLicenseRequestBuilder struct{ BaseRequestBuilder }

// RevokeUserLicense action undocumented
func (b *IOSVPPAppRequestBuilder) RevokeUserLicense(reqObj *IOSVPPAppRevokeUserLicenseRequestParameter) *IOSVPPAppRevokeUserLicenseRequestBuilder {
	bb := &IOSVPPAppRevokeUserLicenseRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/revokeUserLicense"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type IOSVPPAppRevokeUserLicenseRequest struct{ BaseRequest }

//
func (b *IOSVPPAppRevokeUserLicenseRequestBuilder) Request() *IOSVPPAppRevokeUserLicenseRequest {
	return &IOSVPPAppRevokeUserLicenseRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *IOSVPPAppRevokeUserLicenseRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}

//
type IOSVPPAppRevokeDeviceLicenseRequestBuilder struct{ BaseRequestBuilder }

// RevokeDeviceLicense action undocumented
func (b *IOSVPPAppRequestBuilder) RevokeDeviceLicense(reqObj *IOSVPPAppRevokeDeviceLicenseRequestParameter) *IOSVPPAppRevokeDeviceLicenseRequestBuilder {
	bb := &IOSVPPAppRevokeDeviceLicenseRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/revokeDeviceLicense"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type IOSVPPAppRevokeDeviceLicenseRequest struct{ BaseRequest }

//
func (b *IOSVPPAppRevokeDeviceLicenseRequestBuilder) Request() *IOSVPPAppRevokeDeviceLicenseRequest {
	return &IOSVPPAppRevokeDeviceLicenseRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *IOSVPPAppRevokeDeviceLicenseRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}
