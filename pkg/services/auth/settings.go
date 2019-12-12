package auth

import (
	"io/ioutil"
	"sync"

	"github.com/BurntSushi/toml"
	"golang.org/x/xerrors"

	"github.com/grafana/grafana/pkg/infra/log"
	"github.com/grafana/grafana/pkg/setting"
	"github.com/grafana/grafana/pkg/util/errutil"
)

// AuthConfig holds a list of Org Group Mappings
type AuthConfig struct {
	AuthMappings []*AuthOrgConfig `toml:"auth"`
}

// AuthOrgConfig Lists Groups and Roles per Organization
type AuthOrgConfig struct {
	Groups []*GroupToOrgRole `toml:"group_mappings"`
}

// GroupToOrgRole is a struct representation of
// config "group_mappings" setting
type GroupToOrgRole struct {
	GroupDN string `toml:"group_dn"`
	OrgID   int64  `toml:"org_id"`

	// This pointer specifies if setting was set (for backwards compatibility)
	IsGrafanaAdmin *bool `toml:"grafana_admin"`

	OrgRole string `toml:"org_role"`
}

// logger
var logger = log.New("auth.settings")

// loadingMutex locks the reading of the config so multiple requests for reloading are sequential.
var loadingMutex = &sync.Mutex{}

func IsEnabled() bool {
	return true // TODO revisit if oauth has a setting
}

// ReloadConfig reads the config from the disc and caches it.
func ReloadConfig(configFile string) error {
	if !IsEnabled() {
		return nil
	}

	loadingMutex.Lock()
	defer loadingMutex.Unlock()

	var err error
	config, err = readConfig(configFile)
	return err
}

// We need to define in this space so `GetConfig` fn
// could be defined as singleton
var config *AuthConfig

// GetConfig returns the OAuth Mapping config
func GetConfig(configFile string) (*AuthConfig, error) {
	if !IsEnabled() {
		return nil, nil
	}

	// Make it a singleton
	if config != nil {
		return config, nil
	}

	loadingMutex.Lock()
	defer loadingMutex.Unlock()

	var err error
	config, err = readConfig(configFile)

	return config, err
}

func readConfig(configFile string) (*AuthConfig, error) {
	result := &AuthConfig{}

	logger.Info("OAuth enabled, reading config file", "file", configFile)

	fileBytes, err := ioutil.ReadFile(configFile)
	if err != nil {
		return nil, errutil.Wrap("Failed to load OAuth config file", err)
	}

	// interpolate full toml string (it can contain ENV variables)
	stringContent := setting.EvalEnvVarExpression(string(fileBytes))

	_, err = toml.Decode(stringContent, result)
	if err != nil {
		return nil, errutil.Wrap("Failed to load OAuth config file", err)
	}

	if len(result.AuthMappings) == 0 {
		return nil, xerrors.New("OAuth enabled but no mapping defined in config file")
	}

	// set default org id
	for _, auth := range result.AuthMappings {

		for _, groupMap := range auth.Groups {
			if groupMap.OrgID == 0 {
				groupMap.OrgID = 1
			}
		}
	}

	return result, nil
}
