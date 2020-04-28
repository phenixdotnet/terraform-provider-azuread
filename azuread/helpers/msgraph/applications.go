package msgraph

import (
	"context"
	"log"
	"net/http"

	"github.com/Azure/azure-sdk-for-go/services/graphrbac/1.6/graphrbac"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/date"
	msgraphapi "github.com/yaegashi/msgraph.go/beta"
)

const msGraphAPIUrl = "https://graph.microsoft.com"

// ApplicationsClient is the the Graph RBAC Management Client
type ApplicationsClient struct {
	graphrbac.BaseClient
}

/*Create Create an application by using the MS Graph API */
func (client ApplicationsClient) Create(ctx context.Context, parameters graphrbac.ApplicationCreateParameters, accessTokenAcceptedVersion *int, owners []string) (result graphrbac.Application, err error) {

	application := convertApplicationCreateParametersFromGraphRBAC(parameters, accessTokenAcceptedVersion, owners)
	req, err := client.createPreparer(ctx, application)
	if err != nil {
		log.Printf("Error in create preparer: %v", err)
		err = autorest.NewErrorWithError(err, "graphrbac.ApplicationsClient", "Create", nil, "Failure preparing request")
		return
	}

	resp, err := client.createSender(req)
	if err != nil {
		log.Printf("Error in create sender: %v", err)
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "graphrbac.ApplicationsClient", "Create", resp, "Failure sending request")
		return
	}

	//var msgraphResult msgraphapi.Application
	msgraphResult, err := client.createResponder(resp)
	if err != nil {
		log.Printf("Error in create responder: %v", err)
		err = autorest.NewErrorWithError(err, "graphrbac.ApplicationsClient", "Create", resp, "Failure responding to request")
		return
	}

	result = convertApplicationToGraphRBACApplication(*msgraphResult)
	return
}

// Get get an application by object ID.
// Parameters:
// applicationObjectID - application object ID.
func (client ApplicationsClient) Get(ctx context.Context, applicationObjectID string) (result graphrbac.Application, err error) {

	req, err := client.getPreparer(ctx, applicationObjectID)
	if err != nil {
		log.Printf("Error in Get getPreparer: %v", err)
		err = autorest.NewErrorWithError(err, "msgraph.ApplicationsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.getSender(req)
	if err != nil {
		log.Printf("Error in Get GetSender: %v", err)
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "msgraph.ApplicationsClient", "Get", resp, "Failure sending request")
		return
	}

	msgraphResult, err := client.getResponder(resp)
	if err != nil {
		log.Printf("Error in Get GetResponder: %v", err)
		err = autorest.NewErrorWithError(err, "msgraph.ApplicationsClient", "Get", resp, "Failure responding to request")
	}

	result = convertApplicationToGraphRBACApplication(msgraphResult)
	return
}

func (client ApplicationsClient) Delete(ctx context.Context, applicationObjectID string) (result autorest.Response, err error) {
	req, err := client.deletePreparer(ctx, applicationObjectID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "graphrbac.ApplicationsClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.deleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "graphrbac.ApplicationsClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.deleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "graphrbac.ApplicationsClient", "Delete", resp, "Failure responding to request")
	}

	return
}

// Patch update an existing application.
// Parameters:
// applicationObjectID - application object ID.
// parameters - parameters to update an existing application.
func (client ApplicationsClient) Patch(ctx context.Context, applicationObjectID string, parameters graphrbac.ApplicationUpdateParameters, accessTokenAcceptedVersion *int, owners []string) (result autorest.Response, err error) {
	application := convertApplicationUpdateParametersFromGraphRBAC(parameters, accessTokenAcceptedVersion, owners)
	req, err := client.patchPreparer(ctx, applicationObjectID, application)
	if err != nil {
		err = autorest.NewErrorWithError(err, "graphrbac.ApplicationsClient", "Patch", nil, "Failure preparing request")
		return
	}

	resp, err := client.patchSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "graphrbac.ApplicationsClient", "Patch", resp, "Failure sending request")
		return
	}

	result, err = client.patchResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "graphrbac.ApplicationsClient", "Patch", resp, "Failure responding to request")
	}

	return
}

// CreatePreparer prepare the request
func (client ApplicationsClient) createPreparer(ctx context.Context, parameters msgraphapi.Application) (*http.Request, error) {

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/applications"),
		autorest.WithJSON(parameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client ApplicationsClient) createSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// CreateResponder handles the response to the Create request. The method always
// closes the http.Response Body.
func (client ApplicationsClient) createResponder(resp *http.Response) (result *msgraphapi.Application, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	//res := autorest.Response{Response: resp}

	if err != nil {
		return nil, err
	}

	log.Printf("Response: %v", result.ID)

	return
}

// GetPreparer prepares the Get request.
func (client ApplicationsClient) getPreparer(ctx context.Context, applicationObjectID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"applicationObjectId": autorest.Encode("path", applicationObjectID),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/applications/{applicationObjectId}", pathParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client ApplicationsClient) getSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ApplicationsClient) getResponder(resp *http.Response) (result msgraphapi.Application, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	//result.Response = autorest.Response{Response: resp}
	return
}

// DeletePreparer prepares the Delete request.
func (client ApplicationsClient) deletePreparer(ctx context.Context, applicationObjectID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"applicationObjectId": autorest.Encode("path", applicationObjectID),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/applications/{applicationObjectId}", pathParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client ApplicationsClient) deleteSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client ApplicationsClient) deleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// PatchPreparer prepares the Patch request.
func (client ApplicationsClient) patchPreparer(ctx context.Context, applicationObjectID string, parameters msgraphapi.Application) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"applicationObjectId": autorest.Encode("path", applicationObjectID),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/applications/{applicationObjectId}", pathParameters),
		autorest.WithJSON(parameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// PatchSender sends the Patch request. The method will close the
// http.Response Body if it receives an error.
func (client ApplicationsClient) patchSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// PatchResponder handles the response to the Patch request. The method always
// closes the http.Response Body.
func (client ApplicationsClient) patchResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// ListPasswordCredentials get the passwordCredentials associated with an application.
// Parameters:
// applicationObjectID - application object ID.
func (client ApplicationsClient) ListPasswordCredentials(ctx context.Context, applicationObjectID string) (result graphrbac.PasswordCredentialListResult, err error) {

	req, err := client.listPasswordCredentialsPreparer(ctx, applicationObjectID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "graphrbac.ApplicationsClient", "ListPasswordCredentials", nil, "Failure preparing request")
		return
	}

	resp, err := client.listPasswordCredentialsSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "graphrbac.ApplicationsClient", "ListPasswordCredentials", resp, "Failure sending request")
		return
	}

	application, err := client.listPasswordCredentialsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "graphrbac.ApplicationsClient", "ListPasswordCredentials", resp, "Failure responding to request")
	}

	pcs := make([]graphrbac.PasswordCredential, len(application.PasswordCredentials))
	for i, p := range application.PasswordCredentials {
		pcs[i] = convertPasswordCredentialToRBAC(p)
	}
	result.Value = &pcs

	return
}

// ListPasswordCredentialsPreparer prepares the ListPasswordCredentials request.
func (client ApplicationsClient) listPasswordCredentialsPreparer(ctx context.Context, applicationObjectID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"applicationObjectId": autorest.Encode("path", applicationObjectID),
	}

	queryParameters := map[string]interface{}{
		"$select": "passwordCredentials",
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/applications/{applicationObjectId}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListPasswordCredentialsSender sends the ListPasswordCredentials request. The method will close the
// http.Response Body if it receives an error.
func (client ApplicationsClient) listPasswordCredentialsSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ListPasswordCredentialsResponder handles the response to the ListPasswordCredentials request. The method always
// closes the http.Response Body.
func (client ApplicationsClient) listPasswordCredentialsResponder(resp *http.Response) (result msgraphapi.Application, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())

	return
}

// CreatePasswordCredential create a passwordCredential associated with an application.
// Parameters:
// applicationObjectID - application object ID.
// parameters - parameters to update passwordCredentials of an existing application.
func (client ApplicationsClient) CreatePasswordCredential(ctx context.Context, applicationObjectID string, parameters *graphrbac.PasswordCredential) (result *graphrbac.PasswordCredential, e error) {

	passwordCredential := convertPasswordCredentialFromRBAC(*parameters)

	req, err := client.createPasswordCredentialsPreparer(ctx, applicationObjectID, passwordCredential)
	if err != nil {
		err = autorest.NewErrorWithError(err, "graphrbac.ApplicationsClient", "UpdatePasswordCredentials", nil, "Failure preparing request")
		return
	}

	resp, err := client.createPasswordCredentialsSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "graphrbac.ApplicationsClient", "UpdatePasswordCredentials", resp, "Failure sending request")
		return
	}

	pc, err := client.createPasswordCredentialsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "graphrbac.ApplicationsClient", "UpdatePasswordCredentials", resp, "Failure responding to request")
	}

	var i = convertPasswordCredentialToRBAC(pc)
	result = &i

	return
}

// UpdatePasswordCredentialsPreparer prepares the UpdatePasswordCredentials request.
func (client ApplicationsClient) createPasswordCredentialsPreparer(ctx context.Context, applicationObjectID string, parameters msgraphapi.PasswordCredential) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"applicationObjectId": autorest.Encode("path", applicationObjectID),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/applications/{applicationObjectId}/addPassword", pathParameters),
		autorest.WithJSON(parameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdatePasswordCredentialsSender sends the UpdatePasswordCredentials request. The method will close the
// http.Response Body if it receives an error.
func (client ApplicationsClient) createPasswordCredentialsSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// UpdatePasswordCredentialsResponder handles the response to the UpdatePasswordCredentials request. The method always
// closes the http.Response Body.
func (client ApplicationsClient) createPasswordCredentialsResponder(resp *http.Response) (result msgraphapi.PasswordCredential, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	return
}

func (client ApplicationsClient) DeletePasswordCredential(ctx context.Context, applicationObjectID string, id string) (e error) {
	keyID := msgraphapi.UUID(id)
	parameters := msgraphapi.PasswordCredential{
		KeyID: &keyID,
	}

	req, err := client.deletePasswordCredentialsPreparer(ctx, applicationObjectID, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "graphrbac.ApplicationsClient", "UpdatePasswordCredentials", nil, "Failure preparing request")
		return
	}

	resp, err := client.deletePasswordCredentialsSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "graphrbac.ApplicationsClient", "UpdatePasswordCredentials", resp, "Failure sending request")
		return
	}

	err = client.deletePasswordCredentialsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "graphrbac.ApplicationsClient", "UpdatePasswordCredentials", resp, "Failure responding to request")
	}

	return
}

// UpdatePasswordCredentialsPreparer prepares the UpdatePasswordCredentials request.
func (client ApplicationsClient) deletePasswordCredentialsPreparer(ctx context.Context, applicationObjectID string, parameters msgraphapi.PasswordCredential) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"applicationObjectId": autorest.Encode("path", applicationObjectID),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/applications/{applicationObjectId}/removePassword", pathParameters),
		autorest.WithJSON(parameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdatePasswordCredentialsSender sends the UpdatePasswordCredentials request. The method will close the
// http.Response Body if it receives an error.
func (client ApplicationsClient) deletePasswordCredentialsSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// UpdatePasswordCredentialsResponder handles the response to the UpdatePasswordCredentials request. The method always
// closes the http.Response Body.
func (client ApplicationsClient) deletePasswordCredentialsResponder(resp *http.Response) (err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusNoContent),
		autorest.ByClosing())
	return
}

/* Converters */

func convertApplicationCreateParametersFromGraphRBAC(parameters graphrbac.ApplicationCreateParameters, accessTokenAcceptedVersion *int, owners []string) msgraphapi.Application {
	gmc := string(parameters.GroupMembershipClaims)
	if gmc == "" {
		gmc = "none"
	}

	var o []msgraphapi.DirectoryObject
	for _, i := range owners {
		var item msgraphapi.DirectoryObject
		item.ID = &i
		o = append(o, item)
	}

	var marketingURL, PrivacyStatementURL, TermsOfServiceURL *string
	if parameters.InformationalUrls != nil {
		marketingURL = parameters.InformationalUrls.Marketing
		PrivacyStatementURL = parameters.InformationalUrls.Privacy
		TermsOfServiceURL = parameters.InformationalUrls.TermsOfService
	}

	return msgraphapi.Application{
		API: &msgraphapi.APIApplication{
			RequestedAccessTokenVersion: accessTokenAcceptedVersion,
			KnownClientApplications:     convertSliceOfStringToSliceOfUUID(parameters.KnownClientApplications),
		},
		AppRoles:               convertAppRoleFromGraphRBAC(parameters.AppRoles),
		IsFallbackPublicClient: parameters.AvailableToOtherTenants,
		IdentifierUris:         *parameters.IdentifierUris,
		Owners:                 o,
		DisplayName:            parameters.DisplayName,
		GroupMembershipClaims:  &gmc,
		Info: &msgraphapi.InformationalURL{
			LogoURL:             parameters.LogoutURL,
			MarketingURL:        marketingURL,
			PrivacyStatementURL: PrivacyStatementURL,
			TermsOfServiceURL:   TermsOfServiceURL,
		},
		IsDeviceOnlyAuthSupported: parameters.IsDeviceOnlyAuthSupported,
		KeyCredentials:            convertKeyCredentialsFromRBAC(parameters.KeyCredentials),
		OptionalClaims:            convertOptionalClaimsFromRBAC(parameters.OptionalClaims),
		PasswordCredentials:       convertPasswordCredentialsFromRBAC(parameters.PasswordCredentials),
		PublicClient:              &msgraphapi.PublicClientApplication{},
		PublisherDomain:           parameters.PublisherDomain,
		RequiredResourceAccess:    convertRequiredResourceAccessFromRBAC(parameters.RequiredResourceAccess),
		SignInAudience:            parameters.SignInAudience,
		Web: &msgraphapi.WebApplication{
			HomePageURL:  parameters.Homepage,
			LogoutURL:    parameters.LogoutURL,
			RedirectUris: *parameters.ReplyUrls,
			ImplicitGrantSettings: &msgraphapi.ImplicitGrantSettings{
				EnableIDTokenIssuance:     parameters.Oauth2AllowImplicitFlow,
				EnableAccessTokenIssuance: parameters.Oauth2AllowImplicitFlow,
			},
		},
	}
}

func convertApplicationUpdateParametersFromGraphRBAC(parameters graphrbac.ApplicationUpdateParameters, accessTokenAcceptedVersion *int, owners []string) msgraphapi.Application {
	gmc := string(parameters.GroupMembershipClaims)
	if gmc == "" {
		gmc = "none"
	}

	var o []msgraphapi.DirectoryObject
	for _, i := range owners {
		var item msgraphapi.DirectoryObject
		item.ID = &i
		o = append(o, item)
	}

	var marketingURL, privacyStatementURL, termsOfServiceURL, supportURL *string
	if parameters.InformationalUrls != nil {
		marketingURL = parameters.InformationalUrls.Marketing
		privacyStatementURL = parameters.InformationalUrls.Privacy
		termsOfServiceURL = parameters.InformationalUrls.TermsOfService
	}

	var redirectUris []string
	if parameters.ReplyUrls != nil {
		redirectUris = *parameters.ReplyUrls
	}
	var identifierUris []string
	if parameters.IdentifierUris != nil {
		identifierUris = *parameters.IdentifierUris
	}

	return msgraphapi.Application{
		API: &msgraphapi.APIApplication{
			RequestedAccessTokenVersion: accessTokenAcceptedVersion,
			KnownClientApplications:     convertSliceOfStringToSliceOfUUID(parameters.KnownClientApplications),
		},
		AppRoles:               convertAppRoleFromGraphRBAC(parameters.AppRoles),
		IsFallbackPublicClient: parameters.PublicClient,
		IdentifierUris:         identifierUris,
		Owners:                 o,
		DisplayName:            parameters.DisplayName,
		GroupMembershipClaims:  &gmc,
		Info: &msgraphapi.InformationalURL{
			LogoURL:             parameters.LogoutURL,
			MarketingURL:        marketingURL,
			PrivacyStatementURL: privacyStatementURL,
			TermsOfServiceURL:   termsOfServiceURL,
			SupportURL:          supportURL,
		},
		IsDeviceOnlyAuthSupported: parameters.IsDeviceOnlyAuthSupported,
		KeyCredentials:            convertKeyCredentialsFromRBAC(parameters.KeyCredentials),
		OptionalClaims:            convertOptionalClaimsFromRBAC(parameters.OptionalClaims),
		PasswordCredentials:       convertPasswordCredentialsFromRBAC(parameters.PasswordCredentials),
		PublicClient:              &msgraphapi.PublicClientApplication{},
		PublisherDomain:           parameters.PublisherDomain,
		RequiredResourceAccess:    convertRequiredResourceAccessFromRBAC(parameters.RequiredResourceAccess),
		SignInAudience:            parameters.SignInAudience,
		Web: &msgraphapi.WebApplication{
			HomePageURL:  parameters.Homepage,
			LogoutURL:    parameters.LogoutURL,
			RedirectUris: redirectUris,
			ImplicitGrantSettings: &msgraphapi.ImplicitGrantSettings{
				EnableIDTokenIssuance:     parameters.Oauth2AllowImplicitFlow,
				EnableAccessTokenIssuance: parameters.Oauth2AllowImplicitFlow,
			},
		},
	}
}

func convertRequiredResourceAccessFromRBAC(requiredResourceAccess *[]graphrbac.RequiredResourceAccess) []msgraphapi.RequiredResourceAccess {
	if requiredResourceAccess == nil {
		return []msgraphapi.RequiredResourceAccess{}
	}

	result := make([]msgraphapi.RequiredResourceAccess, len(*requiredResourceAccess))
	for i := range *requiredResourceAccess {
		resourceAccess := make([]msgraphapi.ResourceAccess, len(*(*requiredResourceAccess)[i].ResourceAccess))
		for j := range *(*requiredResourceAccess)[i].ResourceAccess {
			t := (*(*requiredResourceAccess)[i].ResourceAccess)[j]

			ID := msgraphapi.UUID(*t.ID)

			resourceAccess[j] = msgraphapi.ResourceAccess{
				ID:   &ID,
				Type: t.Type,
			}
		}

		var item msgraphapi.RequiredResourceAccess
		item.ResourceAccess = resourceAccess
		item.ResourceAppID = (*requiredResourceAccess)[i].ResourceAppID

		result[i] = item
	}

	return result
}

func convertRequiredResourceAccessToRBAC(requiredResourceAccess *[]msgraphapi.RequiredResourceAccess) *[]graphrbac.RequiredResourceAccess {
	if requiredResourceAccess == nil {
		return &[]graphrbac.RequiredResourceAccess{}
	}

	result := make([]graphrbac.RequiredResourceAccess, len(*requiredResourceAccess))
	for i := range *requiredResourceAccess {
		resourceAccess := make([]graphrbac.ResourceAccess, len((*requiredResourceAccess)[i].ResourceAccess))
		for j := range (*requiredResourceAccess)[i].ResourceAccess {
			t := (*requiredResourceAccess)[i].ResourceAccess[j]

			ID := string(*t.ID)

			resourceAccess[j] = graphrbac.ResourceAccess{
				ID:   &ID,
				Type: t.Type,
			}
		}

		var item graphrbac.RequiredResourceAccess
		item.ResourceAccess = &resourceAccess
		item.ResourceAppID = (*requiredResourceAccess)[i].ResourceAppID

		result[i] = item
	}

	return &result
}

func convertSliceOfStringToSliceOfUUID(a *[]string) []msgraphapi.UUID {
	if a == nil {
		return []msgraphapi.UUID{}
	}

	b := make([]msgraphapi.UUID, len(*a))
	for i := range *a {
		b[i] = msgraphapi.UUID((*a)[i])
	}

	return b
}

func convertPasswordCredentialsFromRBAC(passwordCredentials *[]graphrbac.PasswordCredential) []msgraphapi.PasswordCredential {
	if passwordCredentials == nil {
		return []msgraphapi.PasswordCredential{}
	}

	var result []msgraphapi.PasswordCredential
	for _, p := range *passwordCredentials {
		item := convertPasswordCredentialFromRBAC(p)
		result = append(result, item)
	}

	return result
}

func convertPasswordCredentialFromRBAC(passwordCredential graphrbac.PasswordCredential) msgraphapi.PasswordCredential {
	var item msgraphapi.PasswordCredential
	var cki msgraphapi.Binary
	if passwordCredential.CustomKeyIdentifier != nil {
		cki = msgraphapi.Binary(*passwordCredential.CustomKeyIdentifier)
	}
	var ID msgraphapi.UUID
	if passwordCredential.KeyID != nil {
		ID = msgraphapi.UUID(*passwordCredential.KeyID)
	}

	item.CustomKeyIdentifier = &cki
	item.KeyID = &ID
	item.SecretText = passwordCredential.Value

	if passwordCredential.StartDate != nil {
		item.StartDateTime = &passwordCredential.StartDate.Time
	}

	if passwordCredential.EndDate != nil {
		item.EndDateTime = &passwordCredential.EndDate.Time
	}

	return item
}

func convertPasswordCredentialsToRBAC(passwordCredentials *[]msgraphapi.PasswordCredential) *[]graphrbac.PasswordCredential {
	if passwordCredentials == nil {
		return &[]graphrbac.PasswordCredential{}
	}

	var result []graphrbac.PasswordCredential
	for _, p := range *passwordCredentials {
		item := convertPasswordCredentialToRBAC(p)
		result = append(result, item)
	}

	return &result
}

func convertPasswordCredentialToRBAC(passwordCredential msgraphapi.PasswordCredential) graphrbac.PasswordCredential {
	var keyID string
	if passwordCredential.KeyID != nil {
		keyID = string(*passwordCredential.KeyID)
	}

	var cki []byte
	if passwordCredential.CustomKeyIdentifier != nil {
		cki = []byte(*passwordCredential.CustomKeyIdentifier)
	}

	var startDate date.Time
	if passwordCredential.StartDateTime != nil {
		startDate = date.Time{*passwordCredential.StartDateTime}
	}

	var endDate date.Time
	if passwordCredential.EndDateTime != nil {
		endDate = date.Time{*passwordCredential.EndDateTime}
	}

	var secretText string
	if passwordCredential.SecretText != nil {
		secretText = *passwordCredential.SecretText
	}

	result := graphrbac.PasswordCredential{
		StartDate:           &startDate,
		EndDate:             &endDate,
		KeyID:               &keyID,
		Value:               &secretText,
		CustomKeyIdentifier: &cki,
	}

	return result
}

func convertAppRoleFromGraphRBAC(appRoles *[]graphrbac.AppRole) []msgraphapi.AppRole {
	if appRoles == nil {
		return nil
	}

	var results []msgraphapi.AppRole
	for _, a := range *appRoles {
		var ID msgraphapi.UUID
		ID = msgraphapi.UUID(*a.ID)

		var item msgraphapi.AppRole
		item.Description = a.Description
		item.DisplayName = a.DisplayName
		item.IsEnabled = a.IsEnabled
		item.Value = a.Value
		item.AllowedMemberTypes = *a.AllowedMemberTypes
		item.ID = &ID

		results = append(results, item)
	}

	return results
}

func convertAppRoleToGraphRBAC(appRoles *[]msgraphapi.AppRole) *[]graphrbac.AppRole {
	if appRoles == nil {
		return nil
	}

	var results []graphrbac.AppRole
	for _, a := range *appRoles {
		var ID string
		ID = string(*a.ID)

		var item graphrbac.AppRole
		item.Description = a.Description
		item.DisplayName = a.DisplayName
		item.IsEnabled = a.IsEnabled
		item.Value = a.Value
		item.AllowedMemberTypes = &a.AllowedMemberTypes
		item.ID = &ID

		results = append(results, item)
	}

	return &results
}

func convertKeyCredentialsFromRBAC(keyCredentials *[]graphrbac.KeyCredential) []msgraphapi.KeyCredential {
	if keyCredentials == nil {
		return nil
	}
	var results []msgraphapi.KeyCredential
	for _, a := range *keyCredentials {
		ID := msgraphapi.UUID(*a.KeyID)

		key := msgraphapi.Binary(*a.Value)
		ki := msgraphapi.Binary(*a.CustomKeyIdentifier)

		var item msgraphapi.KeyCredential
		item.CustomKeyIdentifier = &ki
		item.EndDateTime = &a.EndDate.Time
		item.Key = &key
		item.KeyID = &ID
		item.StartDateTime = &a.StartDate.Time
		item.Type = a.Type
		item.Usage = a.Usage

		results = append(results, item)
	}

	return results
}

func convertKeyCredentialsToRBAC(keyCredentials *[]msgraphapi.KeyCredential) *[]graphrbac.KeyCredential {
	if keyCredentials == nil {
		return nil
	}
	var results []graphrbac.KeyCredential
	for _, a := range *keyCredentials {
		ID := string(*a.KeyID)

		key := string(*a.Key)
		ki := string(*a.CustomKeyIdentifier)

		var item graphrbac.KeyCredential
		item.CustomKeyIdentifier = &ki
		item.EndDate.Time = *a.EndDateTime
		item.Value = &key
		item.KeyID = &ID
		item.StartDate.Time = *a.StartDateTime
		item.Type = a.Type
		item.Usage = a.Usage

		results = append(results, item)
	}

	return &results
}

func convertOptionalClaimsFromRBAC(optionalClaims *graphrbac.OptionalClaims) *msgraphapi.OptionalClaims {
	if optionalClaims == nil {
		return nil
	}

	var result msgraphapi.OptionalClaims
	for _, a := range *optionalClaims.AccessToken {

		result.AccessToken = append(result.AccessToken, convertOptionalClaimFromRBAC(a))
	}

	return &result
}

func convertOptionalClaimFromRBAC(optionalClaim graphrbac.OptionalClaim) msgraphapi.OptionalClaim {
	var result msgraphapi.OptionalClaim
	result.Essential = optionalClaim.Essential
	result.Name = optionalClaim.Name
	result.Source = optionalClaim.Source
	return result
}

func convertApplicationToGraphRBACApplication(application msgraphapi.Application) graphrbac.Application {
	var result graphrbac.Application
	result.AppID = application.AppID
	if application.Info != nil {
		result.AppLogoURL = application.Info.LogoURL
		result.InformationalUrls = &graphrbac.InformationalURL{
			Marketing:      application.Info.MarketingURL,
			TermsOfService: application.Info.TermsOfServiceURL,
			Support:        application.Info.SupportURL,
			Privacy:        application.Info.PrivacyStatementURL,
		}
	}
	result.AppRoles = convertAppRoleToGraphRBAC(&application.AppRoles)
	result.AvailableToOtherTenants = application.IsFallbackPublicClient
	result.DisplayName = application.DisplayName
	if application.GroupMembershipClaims != nil {
		result.GroupMembershipClaims = graphrbac.GroupMembershipClaimTypes(*application.GroupMembershipClaims)
	}
	if application.Web != nil {
		result.Homepage = application.Web.HomePageURL
		result.LogoutURL = application.Web.LogoutURL
		result.ReplyUrls = &application.Web.RedirectUris
		if application.Web.ImplicitGrantSettings != nil {
			result.Oauth2AllowImplicitFlow = application.Web.ImplicitGrantSettings.EnableAccessTokenIssuance
		}
	}
	result.KeyCredentials = convertKeyCredentialsToRBAC(&application.KeyCredentials)
	result.OptionalClaims = convertOptionalClaimsToRBAC(application.OptionalClaims)
	result.PasswordCredentials = convertPasswordCredentialsToRBAC(&application.PasswordCredentials)
	result.PublicClient = application.IsFallbackPublicClient
	result.RequiredResourceAccess = convertRequiredResourceAccessToRBAC(&application.RequiredResourceAccess)
	result.IsDeviceOnlyAuthSupported = application.IsDeviceOnlyAuthSupported
	result.PublisherDomain = application.PublisherDomain
	result.SignInAudience = application.SignInAudience
	result.ObjectID = application.ID
	//	result.DeletionTimestamp = application.DeletedDateTime

	return result
}

func convertOptionalClaimsToRBAC(optionalClaims *msgraphapi.OptionalClaims) *graphrbac.OptionalClaims {
	if optionalClaims == nil {
		return nil
	}

	var result graphrbac.OptionalClaims
	for _, a := range optionalClaims.AccessToken {
		*result.AccessToken = append(*result.AccessToken, convertOptionalClaimToRBAC(a))
	}

	return &result
}

func convertOptionalClaimToRBAC(optionalClaim msgraphapi.OptionalClaim) graphrbac.OptionalClaim {
	var result graphrbac.OptionalClaim
	result.Essential = optionalClaim.Essential
	result.Name = optionalClaim.Name
	result.Source = optionalClaim.Source
	return result
}
