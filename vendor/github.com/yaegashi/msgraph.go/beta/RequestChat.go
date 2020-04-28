// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "context"

// ChatRequestBuilder is request builder for Chat
type ChatRequestBuilder struct{ BaseRequestBuilder }

// Request returns ChatRequest
func (b *ChatRequestBuilder) Request() *ChatRequest {
	return &ChatRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ChatRequest is request for Chat
type ChatRequest struct{ BaseRequest }

// Get performs GET request for Chat
func (r *ChatRequest) Get(ctx context.Context) (resObj *Chat, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for Chat
func (r *ChatRequest) Update(ctx context.Context, reqObj *Chat) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for Chat
func (r *ChatRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// ChatMessageRequestBuilder is request builder for ChatMessage
type ChatMessageRequestBuilder struct{ BaseRequestBuilder }

// Request returns ChatMessageRequest
func (b *ChatMessageRequestBuilder) Request() *ChatMessageRequest {
	return &ChatMessageRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ChatMessageRequest is request for ChatMessage
type ChatMessageRequest struct{ BaseRequest }

// Get performs GET request for ChatMessage
func (r *ChatMessageRequest) Get(ctx context.Context) (resObj *ChatMessage, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ChatMessage
func (r *ChatMessageRequest) Update(ctx context.Context, reqObj *ChatMessage) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ChatMessage
func (r *ChatMessageRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// ChatMessageHostedContentRequestBuilder is request builder for ChatMessageHostedContent
type ChatMessageHostedContentRequestBuilder struct{ BaseRequestBuilder }

// Request returns ChatMessageHostedContentRequest
func (b *ChatMessageHostedContentRequestBuilder) Request() *ChatMessageHostedContentRequest {
	return &ChatMessageHostedContentRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ChatMessageHostedContentRequest is request for ChatMessageHostedContent
type ChatMessageHostedContentRequest struct{ BaseRequest }

// Get performs GET request for ChatMessageHostedContent
func (r *ChatMessageHostedContentRequest) Get(ctx context.Context) (resObj *ChatMessageHostedContent, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ChatMessageHostedContent
func (r *ChatMessageHostedContentRequest) Update(ctx context.Context, reqObj *ChatMessageHostedContent) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ChatMessageHostedContent
func (r *ChatMessageHostedContentRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}
