package cobra

import (
	"fmt"
	"os"
	"path/filepath"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "secret",
	Short: "Secret is an API key and other secrets manager",
}

var encryptKey string

func init() {
	// TODO(jaymonari): Set the encryption key correctly from the environ.
	if encryptKey, ok := os.LookupEnv("SECRETS_ENCRYPTION_KEY"); ok {
		fmt.Println("ENCKEY", encryptKey)
		RootCmd.PersistentFlags().StringVarP(&encryptKey, "key", "k", encryptKey, "the key to use when encrypting and decrypting secrets")
		fmt.Println("ENCKEY AFTER", encryptKey)
	} else {
		fmt.Println("SECRETS_ENCRYPTION_KEY is not set.")
	}
}

func secretsPath() string {
	home, _ := homedir.Dir()
	return filepath.Join(home, ".secrets")
}
