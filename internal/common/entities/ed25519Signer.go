package entities

import (
	"crypto/ed25519"

	"github.com/trustbloc/kms-go/doc/jose"
)

// Ed25519Signer is a Jose compliant signer.
type Ed25519Signer struct {
	privKey []byte
}

// Sign data.
func (s Ed25519Signer) Sign(data []byte) ([]byte, error) {
	return ed25519.Sign(s.privKey, data), nil
}

// Headers defined to be compatible with jose signer. TODO: remove after jose refactoring.
func (s Ed25519Signer) Headers() jose.Headers {
	return jose.Headers{
		jose.HeaderAlgorithm: "EdDSA",
	}
}

// NewEd25519Signer creates Ed25519Signer.
func NewEd25519Signer(ed25519PK []byte) *Ed25519Signer {
	return &Ed25519Signer{privKey: ed25519PK}
}
