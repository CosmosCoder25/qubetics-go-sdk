package core

import (
	"fmt"
	"time"

	"github.com/cometbft/cometbft/rpc/client/http"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	cosmossdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/tx"

	"github.com/qubetics/qubetics-go-sdk/config"
	"github.com/qubetics/qubetics-go-sdk/types"
)

// Client contains all necessary components for transaction handling, query management, and configuration settings.
type Client struct {
	keyring                  keyring.Keyring           // Keyring for managing private keys and signatures
	protoCodec               codec.ProtoCodecMarshaler // Used for marshaling and unmarshaling protobuf data
	queryHeight              int64                     // Query height for blockchain data
	queryProve               bool                      // Flag indicating whether to prove queries
	queryRetryAttempts       uint                      // Number of retry attempts for queries
	queryRetryDelay          time.Duration             // Delay between query retries
	rpcAddr                  string                    // RPC server address
	rpcChainID               string                    // The chain ID used to identify the blockchain network
	rpcTimeout               time.Duration             // RPC timeout duration
	txAuthzGranterAddr       cosmossdk.AccAddress      // Address that grants transaction authorization
	txBroadcastRetryAttempts uint                      // Number of retry attempts for transaction broadcast
	txBroadcastRetryDelay    time.Duration             // Delay between transaction broadcast retries
	txConfig                 client.TxConfig           // Configuration related to transactions (e.g., signing modes)
	txFeeGranterAddr         cosmossdk.AccAddress      // Address that grants transaction fees
	txFees                   cosmossdk.Coins           // Fees for transactions
	txFromName               string                    // Sender name for transactions
	txGasAdjustment          float64                   // Adjustment factor for gas estimation
	txGasPrices              cosmossdk.DecCoins        // Gas price settings for transactions
	txGas                    uint64                    // Gas limit for transactions
	txMemo                   string                    // Memo attached to transactions
	txQueryRetryAttempts     uint                      // Number of retry attempts for transaction queries
	txQueryRetryDelay        time.Duration             // Delay between transaction query retries
	txSimulateAndExecute     bool                      // Flag for simulating and executing transactions
	txTimeoutHeight          uint64                    // Transaction timeout height
}

// NewClient initializes a new Client instance.
func NewClient() *Client {
	// Create a codec for encoding/decoding protocol buffer messages.
	protoCodec := types.NewProtoCodec()
	txConfig := tx.NewTxConfig(protoCodec, tx.DefaultSignModes)

	// Initialize Client with default values and configurations.
	c := &Client{}
	c.WithProtoCodec(protoCodec)
	c.WithTxConfig(txConfig)

	return c
}

// ProtoCodec returns the protobuf codec used for marshaling and unmarshaling data.
func (c *Client) ProtoCodec() codec.ProtoCodecMarshaler {
	return c.protoCodec
}

// WithKeyring assigns the keyring to the Client and returns the updated Client.
func (c *Client) WithKeyring(keyring keyring.Keyring) *Client {
	c.keyring = keyring
	return c
}

// WithProtoCodec sets the protobuf codec and returns the updated Client.
func (c *Client) WithProtoCodec(protoCodec codec.ProtoCodecMarshaler) *Client {
	c.protoCodec = protoCodec
	return c
}

// WithQueryProve sets the prove flag for queries and returns the updated Client.
func (c *Client) WithQueryProve(prove bool) *Client {
	c.queryProve = prove
	return c
}

// WithQueryRetryAttempts sets the number of retry attempts for queries and returns the updated Client.
func (c *Client) WithQueryRetryAttempts(attempts uint) *Client {
	c.queryRetryAttempts = attempts
	return c
}

// WithQueryRetryDelay sets the retry delay duration for queries and returns the updated Client.
func (c *Client) WithQueryRetryDelay(delay time.Duration) *Client {
	c.queryRetryDelay = delay
	return c
}

// WithRPCAddr sets the RPC server address and returns the updated Client.
func (c *Client) WithRPCAddr(rpcAddr string) *Client {
	c.rpcAddr = rpcAddr
	return c
}

// WithRPCChainID sets the blockchain chain ID and returns the updated Client.
func (c *Client) WithRPCChainID(chainID string) *Client {
	c.rpcChainID = chainID
	return c
}

// WithRPCTimeout sets the RPC timeout duration and returns the updated Client.
func (c *Client) WithRPCTimeout(timeout time.Duration) *Client {
	c.rpcTimeout = timeout
	return c
}

// WithTxAuthzGranterAddr sets the transaction authorization granter address and returns the updated Client.
func (c *Client) WithTxAuthzGranterAddr(addr cosmossdk.AccAddress) *Client {
	c.txAuthzGranterAddr = addr
	return c
}

// WithTxBroadcastRetryAttempts sets the number of retry attempts for broadcasting transactions and returns the updated Client.
func (c *Client) WithTxBroadcastRetryAttempts(attempts uint) *Client {
	c.txBroadcastRetryAttempts = attempts
	return c
}

// WithTxBroadcastRetryDelay sets the retry delay duration for broadcasting transactions and returns the updated Client.
func (c *Client) WithTxBroadcastRetryDelay(delay time.Duration) *Client {
	c.txBroadcastRetryDelay = delay
	return c
}

// WithTxConfig sets the transaction configuration and returns the updated Client.
func (c *Client) WithTxConfig(txConfig client.TxConfig) *Client {
	c.txConfig = txConfig
	return c
}

// WithTxFeeGranterAddr sets the transaction fee granter address and returns the updated Client.
func (c *Client) WithTxFeeGranterAddr(addr cosmossdk.AccAddress) *Client {
	c.txFeeGranterAddr = addr
	return c
}

// WithTxFees assigns transaction fees and returns the updated Client.
func (c *Client) WithTxFees(fees cosmossdk.Coins) *Client {
	c.txFees = fees
	return c
}

// WithTxFromName sets the "from" name for transactions and returns the updated Client.
func (c *Client) WithTxFromName(name string) *Client {
	c.txFromName = name
	return c
}

// WithTxGasAdjustment sets the gas adjustment factor for transactions and returns the updated Client.
func (c *Client) WithTxGasAdjustment(adjustment float64) *Client {
	c.txGasAdjustment = adjustment
	return c
}

// WithTxGasPrices sets the gas prices for transactions and returns the updated Client.
func (c *Client) WithTxGasPrices(prices cosmossdk.DecCoins) *Client {
	c.txGasPrices = prices
	return c
}

// WithTxGas sets the gas limit for transactions and returns the updated Client.
func (c *Client) WithTxGas(gas uint64) *Client {
	c.txGas = gas
	return c
}

// WithTxMemo sets the memo for transactions and returns the updated Client.
func (c *Client) WithTxMemo(memo string) *Client {
	c.txMemo = memo
	return c
}

// WithTxQueryRetryAttempts sets the number of retry attempts for transaction queries and returns the updated Client.
func (c *Client) WithTxQueryRetryAttempts(attempts uint) *Client {
	c.txQueryRetryAttempts = attempts
	return c
}

// WithTxQueryRetryDelay sets the retry delay duration for transaction queries and returns the updated Client.
func (c *Client) WithTxQueryRetryDelay(delay time.Duration) *Client {
	c.txQueryRetryDelay = delay
	return c
}

// WithTxSimulateAndExecute sets the simulate and execute flag and returns the updated Client.
func (c *Client) WithTxSimulateAndExecute(simulate bool) *Client {
	c.txSimulateAndExecute = simulate
	return c
}

// WithTxTimeoutHeight sets the timeout height for transactions and returns the updated Client.
func (c *Client) WithTxTimeoutHeight(height uint64) *Client {
	c.txTimeoutHeight = height
	return c
}

// HTTP creates an HTTP client for the given RPC address and timeout configuration.
// Returns the HTTP client or an error if initialization fails.
func (c *Client) HTTP() (*http.HTTP, error) {
	timeout := uint(c.rpcTimeout / time.Second)
	return http.NewWithTimeout(c.rpcAddr, "/websocket", timeout)
}

// NewClientFromConfig creates a new Client instance based on the provided configuration.
func NewClientFromConfig(c *config.Config) (*Client, error) {
	v := NewClient().
		WithQueryProve(c.Query.GetProve()).
		WithQueryRetryAttempts(c.Query.GetRetryAttempts()).
		WithQueryRetryDelay(c.Query.GetRetryDelay()).
		WithRPCAddr(c.RPC.GetAddrs()[0]).
		WithRPCChainID(c.RPC.GetChainID()).
		WithRPCTimeout(c.RPC.GetTimeout()).
		WithTxAuthzGranterAddr(c.Tx.GetAuthzGranterAddr()).
		WithTxBroadcastRetryAttempts(c.Tx.GetBroadcastRetryAttempts()).
		WithTxBroadcastRetryDelay(c.Tx.GetBroadcastRetryDelay()).
		WithTxFeeGranterAddr(c.Tx.GetFeeGranterAddr()).
		WithTxFees(nil).
		WithTxFromName(c.Tx.GetFromName()).
		WithTxGasAdjustment(c.Tx.GetGasAdjustment()).
		WithTxGas(c.Tx.GetGas()).
		WithTxGasPrices(c.Tx.GetGasPrices()).
		WithTxMemo("").
		WithTxQueryRetryAttempts(c.Tx.GetQueryRetryAttempts()).
		WithTxQueryRetryDelay(c.Tx.GetQueryRetryDelay()).
		WithTxSimulateAndExecute(c.Tx.GetSimulateAndExecute()).
		WithTxTimeoutHeight(0)

	// Setup the keyring for the client
	if err := v.SetupKeyring(c.Keyring); err != nil {
		return nil, fmt.Errorf("failed to setup keyring: %w", err)
	}

	return v, nil
}
