package main

import (
	"context"
	"log"
	"os"
	"time"

	vault "github.com/hashicorp/vault/api"
)

const (
	password   = "password123$"
	secretPass = "super-secret-pass"
	mountPath  = "secret"
)

func main() {
	cfg := vault.DefaultConfig()
	cfg.Address = os.Getenv("VAULT_ADDR")
	client, err := vault.NewClient(cfg)
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken(os.Getenv("VAULT_TOKEN"))

	// Put secret into vault
	secretData := map[string]any{"password": password}
	ctx := context.Background()
	if _, err = client.KVv2(mountPath).
		Put(ctx, secretPass, secretData); err != nil {
		log.Fatal(err)
	}
	log.Println("Super secret password written successfully to the vault.")

	// Get secret from vault
	secret, err := client.KVv2(mountPath).Get(ctx, secretPass)
	if err != nil {
		log.Fatal(err)
	}
	val, ok := secret.Data["password"].(string)
	if !ok {
		log.Fatalf("value type assertion failed: %T %#v",
			secret.Data["password"], secret.Data["password"])
	}
	log.Println("Super secret password was retrieved:", val)

	// Get all versions of a secret
	versions, err := client.KVv2(mountPath).GetVersionsAsList(ctx, secretPass)
	if err != nil {
		log.Fatal(
			"unable to retrieve all versions of the super secret password from the vault:",
			err,
		)
	}
	for _, v := range versions {
		del := "Not deleted"
		if !v.DeletionTime.IsZero() {
			del = v.DeletionTime.Format(time.UnixDate)
		}

		secret, err := client.KVv2(mountPath).
			GetVersion(ctx, secretPass, v.Version)
		if err != nil {
			log.Fatal(
				"unable to retrieve version of the super secret password from the vault:",
				err,
			)
		}
		if val, ok := secret.Data["password"].(string); ok {
			log.Printf(
				"Version: %d, Created at: %s, Deleted at: %s, Destroyed: %t, Value: %q\n",
				v.Version, v.CreatedTime.Format(time.UnixDate), del, v.Destroyed, val,
			)
		}
	}

	// Delete single version of a secret
	if err = client.KVv2(mountPath).Delete(ctx, secretPass); err != nil {
		log.Fatalf("Unable to delete the latest version of the secret from the vault. Reason: %v", err)
	}
	log.Println("Delete the latest version of the secret from the vault")

	// Delete all versions
	if err = client.KVv2(mountPath).DeleteMetadata(ctx, secretPass); err != nil {
		log.Fatalf("Unable to entirely delete the super secret password from the vault. Reason: %v", err)
	}
	log.Println("Deleted the latest version of the super secret password from the vault")

}
