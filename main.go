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

	// Make an item template request
	itemTemplateResponse, err := client.GetTemplate("item")
	if err != nil {
		log.Fatal("Error getting item template: ", err)
	}
	prettyPrintJSON(itemTemplateResponse)

	// Make a login item template request
	loginTemplateResponse, err := client.GetTemplate("item.login")
	if err != nil {
		log.Fatal("Error getting login template: ", err)
	}
	prettyPrintJSON(loginTemplateResponse)

	// Make an item.field template request
	itemFieldTemplateResponse, err := client.GetTemplate("item.field")
	if err != nil {
		log.Fatal("Error getting item field template: ", err)
	}
	prettyPrintJSON(itemFieldTemplateResponse)

	// Make an item.login.uri template request
	itemLoginUriTemplateResponse, err := client.GetTemplate("item.login.uri")
	if err != nil {
		log.Fatal("Error getting item login uri template: ", err)
	}
	prettyPrintJSON(itemLoginUriTemplateResponse)

	// Make an item.card template request
	itemCardTemplateResponse, err := client.GetTemplate("item.card")
	if err != nil {
		log.Fatal("Error getting item card template: ", err)
	}
	prettyPrintJSON(itemCardTemplateResponse)

	// Make an item.identity template request
	itemIdentityTemplateResponse, err := client.GetTemplate("item.identity")
	if err != nil {
		log.Fatal("Error getting item identity template: ", err)
	}
	prettyPrintJSON(itemIdentityTemplateResponse)

	// Make an item.securenote template request
	itemSecureNoteTemplateResponse, err := client.GetTemplate("item.securenote")
	if err != nil {
		log.Fatal("Error getting item secure note template: ", err)
	}
	prettyPrintJSON(itemSecureNoteTemplateResponse)

	// Make a folder template request
	folderTemplateResponse, err := client.GetTemplate("folder")
	if err != nil {
		log.Fatal("Error getting folder template: ", err)
	}
	prettyPrintJSON(folderTemplateResponse)

	// Make a collection template request
	collectionTemplateResponse, err := client.GetTemplate("collection")
	if err != nil {
		log.Fatal("Error getting collection template: ", err)
	}
	prettyPrintJSON(collectionTemplateResponse)

	// // Make an item-collections template request // can't figure this one out
	// itemCollectionsTemplateResponse, err := client.GetTemplate("item-collections")
	// if err != nil {
	// 	log.Fatal("Error getting item collections template: ", err)
	// }
	// prettyPrintJSON(itemCollectionsTemplateResponse)

	// Make an org-collection template request
	orgCollectionsTemplateResponse, err := client.GetTemplate("org-collection")
	if err != nil {
		log.Fatal("Error getting org collection template: ", err)
	}
	prettyPrintJSON(orgCollectionsTemplateResponse)

	// List vault items
	items, err := client.ListVaultItems()
	if err != nil {
		log.Fatal("Error getting vault items: ", err)
	}
	prettyPrintJSON(items)

	// Get vault item
	itemIds := items.Data.Data

	for _, itemId := range itemIds {
		item, err := client.GetVaultItem(itemId.ID)
		if err != nil {
			log.Fatal("Error getting vault item: ", err)
		}
		prettyPrintJSON(item)
	}
}
