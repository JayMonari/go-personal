package cobra

import (
	"fmt"
	"secret"

	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Sets a secret in your secret storage",
	Run: func(cmd *cobra.Command, args []string) {
		v := secret.File(encryptKey, secretsPath())
		value, err := v.Get(args[0])
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("%s=%s\n", args[0], value)
	},
}

func init() {
	RootCmd.AddCommand(getCmd)
}
