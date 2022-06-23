package cobra

import (
	"fmt"
	"secret"

	"github.com/spf13/cobra"
)

var setCmd = &cobra.Command{
	Use:   "set",
	Short: "Sets a secret in your secret storage",
	Run: func(cmd *cobra.Command, args []string) {
		v := secret.File(encryptKey, secretsPath())
		key, val := args[0], args[1]
		if err := v.Set(key, val); err != nil {
			fmt.Println(err)
		}
		fmt.Println("Value set successfully!")
	},
}

func init() {
	RootCmd.AddCommand(setCmd)
}
