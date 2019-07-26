package influxdb

import (
	"encoding/json"
	"testing"
)

func TestSecretFieldJSON(t *testing.T) {
	key := SecretField("some key")
	serialized, err := json.Marshal(key)
	if err != nil {
		t.Fatalf("secret key marshal err: %q" + err.Error())
	}
	if string(serialized) != "\"secret: some key\"" {
		t.Fatalf("secret key marshal result is unexpected, got %q, want \"secret: some key\"", string(serialized))
	}
	var deserialized SecretField
	if err := json.Unmarshal(serialized, &deserialized); err != nil {
		t.Fatalf("secret key unmarshal err: %q" + err.Error())
	}
	if deserialized != key {
		t.Fatalf("secret key unmarshal result is unexpected, got %q, want %q", deserialized, key)
	}

}
