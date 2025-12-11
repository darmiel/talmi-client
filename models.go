package client

import "time"

type AuditEntry struct {
	// ID is the unique request ID (X-Correlation-ID)
	ID string `json:"id"`

	// Time is the timestamp of the event
	Time time.Time `json:"time"`

	// Action describing what happened (e.g. "token.mint", "auth.success")
	Action string `json:"action"`

	// Principal identifies who made the request
	Principal *Principal `json:"principal"`

	// RequestedProvider that was targeted
	RequestedProvider string `json:"requested_provider,omitempty"`
	// RequestedIssuer that was used
	RequestedIssuer string `json:"issuer,omitempty"`

	// Decision details
	PolicyName string `json:"policy_name,omitempty"`
	Provider   string `json:"provider,omitempty"`
	Granted    bool   `json:"granted"`
	Error      string `json:"error,omitempty"`

	// Metadata contains artifact details
	Metadata map[string]any `json:"metadata,omitempty"`
}

// Principal represents the authenticated identity of the caller.
// It is produced by an Issuer after verifying an upstream token.
type Principal struct {
	// ID is the unique subject identifier (e.g., email, sub claim).
	ID string
	// Issuer is the name of the trusted issuer that verified this principal.
	Issuer string
	// Attributes are the claims extracted from the upstream token.
	Attributes map[string]string
}

// TokenMetadata represents the state of an issued token.
type TokenMetadata struct {
	// CorrelationID is the unique identifier for the token and ID of the request that created (requested) it.
	CorrelationID string

	// PrincipalID is the unique identifier of the principal who owns this token.
	PrincipalID string

	// Provider is the name of the downstream provider for which this token was issued.
	Provider string

	// ExpiresAt is the expiration time of the issued token.
	// It is used to check if the token is "active".
	ExpiresAt time.Time

	// IssuedAt is the time when the token was issued.
	IssuedAt time.Time

	// Metadata contains extra metadata (like scope, installation_id for GitHub, ...)
	Metadata map[string]any
}

type ErrorResponse struct {
	Error         string `json:"error"`
	CorrelationID string `json:"correlation_id"`
}

// TokenArtifact is the result of a successful Mint operation.
type TokenArtifact struct {
	// Value is the actual secret/token string (e.g., the GitHub Installation Token).
	Value string `json:"value"`

	// ExpiresAt indicates when this token becomes invalid.
	ExpiresAt time.Time `json:"expires_at"`

	// Metadata contains extra information (e.g., "git_user": "x-access-token").
	Metadata map[string]any `json:"metadata,omitempty"`
}
