package conf

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/goccy/go-yaml"
	"github.com/imdario/mergo"
)

var (
	// ErrCouldntLocateConfig returned when config.yml cannot be found in the source code.
	ErrCouldntLocateConfig = errors.New("could not locate a config.yml in your chain. please follow the link for how-to: https://github.com/tendermint/starport/blob/develop/docs/configure/index.md")

	// FileNames holds a list of appropriate names for the config file.
	FileNames = []string{"config.yml", "config.yaml"}

	// DefaultConf holds default configuration.
	DefaultConf = Config{
		//Host: Host{
		//	// when in Docker on MacOS, it only works with 0.0.0.0.
		//	RPC:     "0.0.0.0:26657",
		//	P2P:     "0.0.0.0:26656",
		//	Prof:    "0.0.0.0:6060",
		//	GRPC:    "0.0.0.0:9090",
		//	GRPCWeb: "0.0.0.0:9091",
		//	API:     "0.0.0.0:1317",
		//},
		//Build: Build{
		//	Proto: Proto{
		//		Path: "proto",
		//		ThirdPartyPaths: []string{
		//			"third_party/proto",
		//			"proto_vendor",
		//		},
		//	},
		//},
		//Faucet: Faucet{
		//	Host: "0.0.0.0:4500",
		//},
	}
)

type Genesis struct {
	ChainID  string                 `json:"chain_id" yaml:"chain_id"`
	AppState map[string]interface{} `json:"app_state" yaml:"app_state"`
}

// Config is the user given configuration to do additional setup
// during serve.
type Config struct {
	Type         string          `json:"type" yaml:"type"`
	GlobalConfig *NodeConfig     `json:"config" yaml:"config"`
	Genesis      Genesis         `json:"genesis" yaml:"genesis"`
	Accounts     []Account       `json:"accounts" yaml:"accounts"`
	Validators   []ValidatorInfo `json:"validators" yaml:"validators"`
}

// AccountByName finds account by name.
func (c Config) AccountByName(name string) (acc Account, found bool) {
	for _, acc := range c.Accounts {
		if acc.Name == name {
			return acc, true
		}
	}
	return Account{}, false
}

// Account holds the options related to setting up Cosmos wallets.
type Account struct {
	Name     string   `yaml:"name"`
	Coins    []string `yaml:"coins,omitempty"`
	Mnemonic string   `yaml:"mnemonic,omitempty"`
	Address  string   `json:"address"`

	// The RPCAddress off the chain that account is issued at.
	RPCAddress string `yaml:"rpc_address,omitempty"`
}

//// Validator holds info related to validator settings.
//type Validator struct {
//	Name   string `yaml:"name"`
//	Staked string `yaml:"staked"`
//}

// Init overwrites sdk configurations with given values.
type Init struct {
	// App overwrites appd's config/app.toml configs.
	App map[string]interface{} `yaml:"app"`

	// Client overwrites appd's config/client.toml configs.
	Client map[string]interface{} `yaml:"client"`

	// Config overwrites appd's config/config.toml configs.
	Config map[string]interface{} `yaml:"config"`

	// Home overwrites default home directory used for the app
	Home string `yaml:"home"`

	// KeyringBackend is the default keyring backend to use for blockchain initialization
	KeyringBackend string `yaml:"keyring-backend"`
}

// Host keeps configuration related to started servers.
type Host struct {
	RPC     string `yaml:"rpc"`
	P2P     string `yaml:"p2p"`
	Prof    string `yaml:"prof"`
	GRPC    string `yaml:"grpc"`
	GRPCWeb string `yaml:"grpc-web"`
	API     string `yaml:"api"`
}

// Parse parses config.yml into UserConfig.
func Parse(r io.Reader) (Config, error) {
	var conf Config
	if err := yaml.NewDecoder(r).Decode(&conf); err != nil {
		return conf, err
	}
	if err := mergo.Merge(&conf, DefaultConf); err != nil {
		return Config{}, err
	}
	return conf, validate(conf)
}

// ParseFile parses config.yml from the path.
func ParseFile(path string) (Config, error) {
	file, err := os.Open(path)
	if err != nil {
		return Config{}, nil
	}
	defer file.Close()
	return Parse(file)
}

// validate validates user config.
func validate(conf Config) error {
	if len(conf.Accounts) == 0 {
		return &ValidationError{"at least 1 account is needed"}
	}
	if len(conf.Accounts) == 0 {
		return &ValidationError{"at least 1 validator is needed"}
	}
	return nil
}

// ValidationError is returned when a configuration is invalid.
type ValidationError struct {
	Message string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("config is not valid: %s", e.Message)
}

// LocateDefault locates the default path for the config file, if no file found returns ErrCouldntLocateConfig.
func LocateDefault(root string) (path string, err error) {
	for _, name := range FileNames {
		path = filepath.Join(root, name)
		if _, err := os.Stat(path); err == nil {
			return path, nil
		} else if !os.IsNotExist(err) {
			return "", err
		}
	}
	return "", ErrCouldntLocateConfig
}
