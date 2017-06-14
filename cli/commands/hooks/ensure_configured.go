package hooks

import (
	"errors"
	"github.com/sensu/sensu-go/cli"
	"github.com/spf13/cobra"
)

const (
	// ConfigurationRequirement used to identify the annotation flag for this handler
	//
	// Usage:
	//
	//	my_cmd := cobra.Command{
	//		Use: "Setup",
	//		Annotations: map[string]string{
	//			ConfigurationRequirement: ConfigurationNotRequired,
	//		}
	//	}
	ConfigurationRequirement = "CREDENTIALS_REQUIREMENT"

	// ConfigurationNotRequired specifies that the command does not require
	// credentials to be configured to complete operations
	ConfigurationNotRequired = "NO"
)

// ConfigurationPresent - unless the given command specifies that configuration
// is not required, func checks that host & secret have been configured.
func ConfigurationPresent(cmd *cobra.Command, cli *cli.SensuCli) error {
	// If the command was configured to ignore whether or not the CLI has been
	// configured stop execution.
	if cmd.Annotations[ConfigurationRequirement] == ConfigurationNotRequired {
		return nil
	}

	// Check that both a URL and a secret are present
	if cli.Config.GetString("api-url") == "" || cli.Config.GetString("secret") == "" {
		return errors.New("Unable to locate credentials. You can configure credentials by running \"sensu-cli configure\"")
	}

	return nil
}