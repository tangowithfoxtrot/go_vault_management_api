package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/tangowithfoxtrot/go_vault_management_api/src/client"
	"github.com/tangowithfoxtrot/go_vault_management_api/src/schemas"
)

func prettyPrintJSON(data interface{}) {
	prettyJSON, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(prettyJSON))
}

func main() {

	// Set the base URL and password for the client
	client := &client.ClientSettings{}
	client.SetClientSettings(os.Getenv("BW_BASE_URL"), os.Getenv("BW_PASSWORD"))

	// Make a lock request
	lockResponse, err := client.Lock()
	if err != nil {
		log.Fatal("Error locking vault: ", err)
	}
	prettyPrintJSON(lockResponse)

	// Make an unlock request
	unlockRequest := schemas.UnlockRequest{
		Password: client.GetPassword(),
	}
	unlockResponse, err := client.Unlock(unlockRequest)
	if err != nil {
		log.Fatal("Error unlocking vault: ", err)
	}
	prettyPrintJSON(unlockResponse)

	// Make a status request
	statusResponse, err := client.Status()
	if err != nil {
		log.Fatal("Error getting vault status: ", err)
	}
	prettyPrintJSON(statusResponse)

	// Make a sync request
	syncResponse, err := client.Sync()
	if err != nil {
		log.Fatal("Error syncing vault: ", err)
	}
	prettyPrintJSON(syncResponse)

	// Make a generate request
	params := schemas.GeneratorParams{
		Length:    24,
		Uppercase: true,
		Lowercase: true,
		Number:    true,
		Special:   true,
		// Passphrase: true,
		// Words:      4,
		// Separator:  "/",
		// Capitalize:    true,
		// IncludeNumber: true,
	}

	generatorResponse, err := client.Generate(params)
	if err != nil {
		log.Fatal("Error generating password: ", err)
	}
	prettyPrintJSON(generatorResponse)

	// Make a fingerprint request
	fingerprintResponse, err := client.GetFingerprint()
	if err != nil {
		log.Fatal("Error getting fingerprint: ", err)
	}
	prettyPrintJSON(fingerprintResponse)
}
