package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"reflect"
	"strings"

	"github.com/tangowithfoxtrot/go_vault_management_api/src/schemas"
)

type ClientSettings struct {
	baseURL  string
	password string
}

type IClientSettings interface {
	// Set the base URL and password for the client
	SetClientSettings(baseURL string, password string)
	Lock() (schemas.LockResponse, error)
	Unlock() (schemas.UnlockResponse, error)
	Status() (schemas.StatusResponse, error)
	Sync() (schemas.SyncResponse, error)
	Generate() (schemas.GeneratorResponse, error)
	GetFingerprint() (schemas.FingerprintResponse, error)
}

func (c *ClientSettings) GetFingerprint() (schemas.FingerprintResponse, error) {
	resp, err := http.Get(c.baseURL + "object/fingerprint/me")
	if err != nil {
		log.Fatal(err)
	}

	var fingerprintResponse schemas.FingerprintResponse
	err = json.NewDecoder(resp.Body).Decode(&fingerprintResponse)
	if err != nil {
		log.Fatal(err)
	}

	return fingerprintResponse, nil
}

func (c *ClientSettings) Generate(params schemas.GeneratorParams) (schemas.GeneratorResponse, error) {
	// Encode the GeneratorParams struct as query parameters
	v := url.Values{}

	encodeWithCamelCase(&v, params)

	// Append the query parameters to the URL
	url := c.baseURL + "generate?" + v.Encode()

	resp, err := http.Get(url)
	if err != nil {
		return schemas.GeneratorResponse{}, err
	}

	var generatorResponse schemas.GeneratorResponse
	err = json.NewDecoder(resp.Body).Decode(&generatorResponse)
	if err != nil {
		return schemas.GeneratorResponse{}, err
	}

	return generatorResponse, nil
}

func (c *ClientSettings) Sync() (schemas.SyncResponse, error) {
	resp, err := http.Post(c.baseURL+"sync", "application/json", nil)
	if err != nil {
		log.Fatal(err)
	}

	var syncResponse schemas.SyncResponse
	err = json.NewDecoder(resp.Body).Decode(&syncResponse)
	if err != nil {
		log.Fatal(err)
	}

	return syncResponse, nil
}

func (c *ClientSettings) Status() (schemas.StatusResponse, error) {
	resp, err := http.Get(c.baseURL + "status")
	if err != nil {
		log.Fatal(err)
	}

	var statusResponse schemas.StatusResponse
	err = json.NewDecoder(resp.Body).Decode(&statusResponse)
	if err != nil {
		log.Fatal(err)
	}

	return statusResponse, nil
}

func (c *ClientSettings) Unlock(unlockRequest schemas.UnlockRequest) (schemas.UnlockResponse, error) {
	reqBody, err := json.Marshal(unlockRequest)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := http.Post(c.baseURL+"unlock", "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		log.Fatal(err)
	}

	var unlockResponse schemas.UnlockResponse
	err = json.NewDecoder(resp.Body).Decode(&unlockResponse)
	if err != nil {
		log.Fatal(err)
	}

	return unlockResponse, nil
}

func (c *ClientSettings) Lock() (schemas.LockResponse, error) {
	resp, err := http.Post(c.baseURL+"lock", "application/json", nil)
	if err != nil {
		log.Fatal(err)
	}

	var lockResponse schemas.LockResponse
	err = json.NewDecoder(resp.Body).Decode(&lockResponse)
	if err != nil {
		log.Fatal(err)
	}

	return lockResponse, nil
}

func (c *ClientSettings) SetClientSettings(baseURL string, password string) {
	if !strings.HasSuffix(baseURL, "/") {
		baseURL += "/"
	}
	c.baseURL = baseURL
	c.password = password
}

func (c *ClientSettings) GetPassword() string {
	return c.password
}

func encodeWithCamelCase(v *url.Values, data interface{}) {
	val := reflect.ValueOf(data)
	typ := reflect.TypeOf(data)

	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		value := val.Field(i)

		// Get the field's value and type
		fieldValue := value.Interface()
		fieldType := value.Type()

		// Only include fields with non-zero values;
		// the /generate endpoint will force passphrases to be true, even if they're false
		if !reflect.DeepEqual(fieldValue, reflect.Zero(fieldType).Interface()) {
			// Convert PascalCase field name to camelCase
			camelCaseFieldName := pascalToCamel(field.Name)

			// Encode the field value into the URL
			v.Add(camelCaseFieldName, fmt.Sprintf("%v", fieldValue))
		}
	}
}

// Convert PascalCase to camelCase
func pascalToCamel(s string) string {
	if len(s) == 0 {
		return ""
	}
	return strings.ToLower(s[:1]) + s[1:]
}
