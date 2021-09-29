package conf

import (
	"time"
)

//-----------------------------------------------------------------------------
// BaseConfig

// BaseConfig defines the base configuration for a Tendermint node
type BaseConfig struct {
	// A custom human readable name for this node
	//Moniker string `json:"moniker" mapstructure:"moniker"`
}

//-----------------------------------------------------------------------------
// RPCConfig

// RPCConfig defines the configuration options for the Tendermint RPC server
type RPCConfig struct {
	// TCP or UNIX socket address for the RPC server to listen on
	ListenAddress string `yaml:"laddr" mapstructure:"laddr" json:"listen_address"`

	// A list of origins a cross-domain request can be executed from.
	// If the special '*' value is present in the list, all origins will be allowed.
	// An origin may contain a wildcard (*) to replace 0 or more characters (i.e.: http://*.domain.com).
	// Only one wildcard can be used per origin.
	CORSAllowedOrigins []string `yaml:"cors_allowed_origins" mapstructure:"cors_allowed_origins" json:"cors_allowed_origins"`
}

//-----------------------------------------------------------------------------
// ConsensusConfig

// ConsensusConfig defines the configuration for the Tendermint consensus service,
// including timeouts and details about the WAL and the block structure.
type ConsensusConfig struct {
	TimeoutCommit time.Duration `json:"timeout_commit" mapstructure:"timeout_commit" json:"timeout_commit"`
	// EmptyBlocks mode and possible interval between empty blocks
	CreateEmptyBlocks         bool          `yaml:"create_empty_blocks" json:"create_empty_blocks" mapstructure:"create_empty_blocks"`
	CreateEmptyBlocksInterval time.Duration `yaml:"create_empty_blocks_interval" json:"create_empty_blocks_interval" mapstructure:"create_empty_blocks_interval"`
}

//-----------------------------------------------------------------------------
// TxIndexConfig

// TxIndexConfig defines the configuration for the transaction indexer,
// including tags to index.
type TxIndexConfig struct {
	// Comma-separated list of tags to index (by default the only tag is "tx.hash")
	//
	// You can also index transactions by height by adding "tx.height" tag here.
	//
	// It's recommended to index only a subset of tags due to possible memory
	// bloat. This is, of course, depends on the indexer's DB and the volume of
	// transactions.
	IndexTags string `yaml:"index_tags" json:"index_tags" mapstructure:"index_tags"`

	// When set to true, tells indexer to index all tags (predefined tags:
	// "tx.hash", "tx.height" and all tags from DeliverTx responses).
	//
	// Note this may be not desirable (see the comment above). IndexTags has a
	// precedence over IndexAllTags (i.e. when given both, IndexTags will be
	// indexed).
	IndexAllTags bool `yaml:"index_all_tags" json:"index_all_tags" mapstructure:"index_all_tags"`
}

type NodeConfig struct {
	// Top level options use an anonymous struct
	BaseConfig `mapstructure:",squash"`

	// Options for services
	RPC       *RPCConfig       `mapstructure:"rpc" json:"rpc" yaml:"rpc"`
	Consensus *ConsensusConfig `mapstructure:"consensus" json:"consensus" yaml:"consensus"`
	TxIndex   *TxIndexConfig   `mapstructure:"tx_index" json:"tx_index" yaml:"tx_index"`
}
