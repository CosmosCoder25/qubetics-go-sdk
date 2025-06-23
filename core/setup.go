package core

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/crypto/keyring"

	qubeticshd "github.com/qubetics/qubetics-blockchain/v2/crypto/hd"
	"github.com/qubetics/qubetics-go-sdk/config"
)

// SetupKeyring initializes and configures a keyring for cryptographic key management.
func (c *Client) SetupKeyring(cfg *config.KeyringConfig) error {
	// Create a keyring instance using the provided configuration.
	kr, err := keyring.New(cfg.GetName(), cfg.GetBackend(), cfg.GetHomeDir(), cfg.GetInput(), c.ProtoCodec(), qubeticshd.EthSecp256k1Option())
	if err != nil {
		return fmt.Errorf("failed to create keyring: %w", err)
	}

	fmt.Println("kr======", kr)
	// Assign the created keyring to the client.
	c.WithKeyring(kr)
	return nil
}
