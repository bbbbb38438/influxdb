package influxdb

import (
	"context"
	"encoding/json"
)

// ErrSecretNotFound is the error msg for a missing secret.
const ErrSecretNotFound = "secret not found"

// SecretService a service for storing and retrieving secrets.
type SecretService interface {
	// LoadSecret retrieves the secret value v found at key k for organization orgID.
	LoadSecret(ctx context.Context, orgID ID, k string) (string, error)

	// GetSecretKeys retrieves all secret keys that are stored for the organization orgID.
	GetSecretKeys(ctx context.Context, orgID ID) ([]string, error)

	// PutSecret stores the secret pair (k,v) for the organization orgID.
	PutSecret(ctx context.Context, orgID ID, k string, v string) error

	// PutSecrets puts all provided secrets and overwrites any previous values.
	PutSecrets(ctx context.Context, orgID ID, m map[string]string) error

	// PatchSecrets patches all provided secrets and updates any previous values.
	PatchSecrets(ctx context.Context, orgID ID, m map[string]string) error

	// DeleteSecret removes a single secret from the secret store.
	DeleteSecret(ctx context.Context, orgID ID, ks ...string) error
}

// SecretField is a key string, can be used to fetch the secret value.
type SecretField string

// String returns the key of the secret.
func (s SecretField) String() string {
	return "secret: " + string(s)
}

// MarshalJSON implement the json marshaler interface.
func (s SecretField) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.String())
}

// UnmarshalJSON implement the json unmarshaler interface.
func (s *SecretField) UnmarshalJSON(b []byte) error {
	var ss string
	if err := json.Unmarshal(b, &ss); err != nil {
		return err
	}
	*s = SecretField(ss[len("secret: "):])
	return nil
}
