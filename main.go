package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"

	"cloud.google.com/go/compute/metadata"
	kms "cloud.google.com/go/kms/apiv1"
	kmspb "google.golang.org/genproto/googleapis/cloud/kms/v1"
	"gopkg.in/alecthomas/kingpin.v2"
)

func main() {
	keyring := kingpin.Flag("keyring", "Name of keyring").Required().String()
	key := kingpin.Flag("key", "Name of key").Required().String()
	location := kingpin.Flag("location", "Location of keyring").Default("global").String()
	project := kingpin.Flag("project", "GCP project. Attempts read project from metadata server if not given.").String()
	kingpin.Parse()

	if *project == "" {
		*project, _ = metadata.ProjectID()
		if *project == "" {
			fmt.Fprintln(os.Stderr, "flag --project not provided, and unable to detect current project")
			os.Exit(1)
		}
	}

	ctx := context.Background()
	c, err := kms.NewKeyManagementClient(ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to connect to KMS: %v\n", err)
		os.Exit(1)
	}

	keyID := fmt.Sprintf(
		"projects/%s/locations/%s/keyRings/%s/cryptoKeys/%s",
		*project,
		*location,
		*keyring,
		*key,
	)

	ciphertext, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "unable to read ciphertext: %v\n", err)
		os.Exit(1)
	}
	req := &kmspb.DecryptRequest{
		Name:       keyID,
		Ciphertext: ciphertext,
	}
	resp, err := c.Decrypt(ctx, req)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to decrypt input: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(string(resp.GetPlaintext()))
}
