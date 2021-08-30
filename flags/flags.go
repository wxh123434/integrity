package flags

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
	tmcli "github.com/tendermint/tendermint/libs/cli"

	"github.com/cosmos/cosmos-sdk/crypto/keyring"
)




func NewCompletionCmd(rootCmd *cobra.Command, hidden bool) *cobra.Command {
	flagZsh := "zsh"
	cmd := &cobra.Command{
		Use:   "completion",
		Short: "Generate Bash/Zsh completion script to STDOUT",
		Long: `To load completion script run
. <(completion_script)
To configure your bash shell to load completions for each session add to your bashrc
# ~/.bashrc or ~/.profile
. <(completion_script)
`,
		RunE: func(_ *cobra.Command, _ []string) error {
			if viper.GetBool(flagZsh) {
				return rootCmd.GenZshCompletion(os.Stdout)
			}
			return rootCmd.GenBashCompletion(os.Stdout)
		},
		Hidden: hidden,
		Args:   cobra.NoArgs,
	}

	cmd.Flags().Bool(flagZsh, false, "Generate Zsh completion script")

	return cmd
}