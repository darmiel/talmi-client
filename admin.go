package client

import (
	"context"
)

// ListAudits retrieves the latest audit entries from the server, limited to the specified number.
func (c *Client) ListAudits(ctx context.Context, limit uint) ([]AuditEntry, error) {
	var resp []AuditEntry
	err := c.get(ctx, c.url().
		setPath(ListAuditsRoute).
		addQueryParam("limit", limit).
		build(), &resp)
	return resp, err
}

// ListActiveTokens retrieves the list of currently active tokens from the server.
func (c *Client) ListActiveTokens(ctx context.Context) ([]TokenMetadata, error) {
	var resp []TokenMetadata
	err := c.get(ctx, c.url().
		setPath(ListActiveTokensRoute).
		build(), &resp)
	return resp, err
}
