package msgraph

import (
	"log"
	"net/http"
	"net/url"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/adal"
	"github.com/Azure/go-autorest/autorest/azure"
)

type MSGraphOAuthTokenProvider struct {
	graphClient autorest.Client
}

func NewMSGraphAutorizer(graphClient autorest.Client) autorest.Authorizer {
	var provider MSGraphOAuthTokenProvider
	provider.graphClient = graphClient

	log.Println("[DEBUG] Creating a new GraphAuthorizer")

	authorizer := autorest.NewBearerAuthorizer(provider)

	return authorizer
}

func (provider MSGraphOAuthTokenProvider) OAuthToken() string {

	log.Println("[DEBUG][MSGRAPHAUTHORIZE] Requesting new OAuth token")

	values := url.Values{}
	values.Add("client_id", "e703529c-5f0a-4a52-89e5-71a9f8ccc1c4")
	values.Add("client_secret", "92zBed87B/AzPVET.ie_w[G.I5oP]3PF")
	values.Add("grant_type", "client_credentials")
	values.Add("scope", "https://graph.microsoft.com/.default")

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL("https://login.microsoftonline.com"),
		autorest.WithPath("/c20cebf5-7d85-41f8-9342-6e134fa7e0f1/oauth2/v2.0/token"),
		autorest.WithFormData(values))

	req, err := preparer.Prepare((&http.Request{}))
	if err != nil {
		log.Println("[MSGRAPHAUTHORIZE] Unable to preparer oauth token request")
		return ""
	}

	resp, err := provider.graphClient.Send(req, autorest.DoRetryForStatusCodes(provider.graphClient.RetryAttempts, provider.graphClient.RetryDuration, autorest.StatusCodesForRetry...))
	if err != nil {
		log.Printf("[MSGRAPHAUTHORIZE] %s", err)
		return ""
	}

	var result adal.Token
	err = autorest.Respond(
		resp,
		provider.graphClient.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())

	log.Println("[DEBUG] Ok new token")
	return result.AccessToken
}
