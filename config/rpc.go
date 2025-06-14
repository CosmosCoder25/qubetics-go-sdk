package config

import (
	"errors"
	"fmt"
	"net/url"
	"strconv"
	"time"

	"github.com/spf13/pflag"
)

// RPCConfig defines the configuration for RPC.
type RPCConfig struct {
	Addrs   []string `mapstructure:"addrs"`    // Addrs is a list of RPC server addresses.
	ChainID string   `mapstructure:"chain_id"` // ChainID is the identifier of the blockchain network.
	Timeout string   `mapstructure:"timeout"`  // Timeout is the duration for RPC requests.
}

// GetAddr returns the first RPC address from the list or an empty string if no addresses are available.
func (c *RPCConfig) GetAddr() string {
	if len(c.GetAddrs()) == 0 {
		return ""
	}

	return c.GetAddrs()[0]
}

// GetAddrs returns the addresses of the RPC servers.
func (c *RPCConfig) GetAddrs() []string {
	return c.Addrs
}

// GetChainID returns the ChainID field.
func (c *RPCConfig) GetChainID() string {
	return c.ChainID
}

// GetTimeout returns the maximum duration for an RPC request.
func (c *RPCConfig) GetTimeout() time.Duration {
	v, err := time.ParseDuration(c.Timeout)
	if err != nil {
		panic(err)
	}
	return v
}

// Validate ensures the RPC configuration is valid.
func (c *RPCConfig) Validate() error {
	// Validate that Addrs is not empty.
	if len(c.Addrs) == 0 {
		return errors.New("addrs cannot be empty")
	}

	// Ensure ChainID is not empty.
	if c.ChainID == "" {
		return errors.New("chain_id cannot be empty")
	}

	// Validate each address in Addrs.
	for _, addr := range c.Addrs {
		if err := validateURL(addr); err != nil {
			return fmt.Errorf("invalid addr: %w", err)
		}
	}

	// Validate that Timeout is a valid time.Duration.
	if _, err := time.ParseDuration(c.Timeout); err != nil {
		return fmt.Errorf("invalid timeout: %w", err)
	}

	return nil
}

// SetForFlags adds rpc configuration flags to the specified FlagSet.
func (c *RPCConfig) SetForFlags(f *pflag.FlagSet) {
	f.StringSliceVar(&c.Addrs, "rpc.addrs", c.Addrs, "addresses of the RPC servers")
	f.StringVar(&c.ChainID, "rpc.chain-id", c.ChainID, "identifier of the blockchain network")
	f.StringVar(&c.Timeout, "rpc.timeout", c.Timeout, "timeout for the RPC requests (e.g., 5s, 500ms)")
}

// DefaultRPCConfig creates a default RPCConfig.
func DefaultRPCConfig() *RPCConfig {
	return &RPCConfig{
		Addrs: []string{
			"https://rpc.qubetics.co:443",
		},
		ChainID: "qubetics-2",
		Timeout: "5s",
	}
}

// validateURL validates an RPC server URL.
func validateURL(s string) error {
	u, err := url.Parse(s)
	if err != nil {
		return fmt.Errorf("invalid url: %w", err)
	}
	if u.Scheme == "" {
		return errors.New("url must have a valid scheme")
	}
	if u.Host == "" {
		return errors.New("url must have a valid host")
	}

	// Check if the port is a valid number.
	port, err := strconv.Atoi(u.Port())
	if err != nil {
		return fmt.Errorf("invalid port: %w", err)
	}
	if port < 1 || port > 65535 {
		return errors.New("url must have a valid port")
	}

	return nil
}
