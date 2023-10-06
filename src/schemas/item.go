package schemas

import (
	"time"

	"github.com/google/uuid"
)

type Item struct {
	OrganizationID *uuid.UUID `json:"organizationId"`
	CollectionIds  *uuid.UUID `json:"collectionIds"`
	FolderID       *uuid.UUID `json:"folderId"`
	Type           *int       `json:"type"`
	Name           *string    `json:"name"`
	Notes          *string    `json:"notes"`
	Favorite       *bool      `json:"favorite"`
	Fields         *[]any     `json:"fields"`
	Login          *any       `json:"login,omitempty"`
	SecureNote     *any       `json:"secureNote,omitempty"`
	Card           *any       `json:"card,omitempty"`
	Identity       *any       `json:"identity,omitempty"`
	Reprompt       *int       `json:"reprompt"`
}

type Template struct {
	Success bool `json:"success"`
	Data    struct {
		Object   string   `json:"object"`
		Template struct { // there has to be a better way to do this, but ¯\_(ツ)_/¯
			OrganizationID *uuid.UUID `json:"organizationId,omitempty"`
			CollectionIds  *uuid.UUID `json:"collectionIds,omitempty"`
			FolderID       *uuid.UUID `json:"folderId,omitempty"`
			Type           *int       `json:"type,omitempty"`
			Name           *string    `json:"name,omitempty"`
			Notes          *string    `json:"notes,omitempty"`
			Favorite       *bool      `json:"favorite,omitempty"`
			Fields         *[]any     `json:"fields,omitempty"`
			Login          *any       `json:"login,omitempty"`
			SecureNote     *any       `json:"secureNote,omitempty"`
			Card           *any       `json:"card,omitempty"`
			Identity       *any       `json:"identity,omitempty"`
			Reprompt       *int       `json:"reprompt,omitempty"`
			Uris           []struct {
				Match any    `json:"match"`
				URI   string `json:"uri"`
			} `json:"uris,omitempty"`
			Username             *string    `json:"username,omitempty"`
			Password             *string    `json:"password,omitempty"`
			Totp                 *string    `json:"totp,omitempty"`
			PasswordRevisionDate *time.Time `json:"passwordRevisionDate,omitempty"`
			CardholderName       *string    `json:"cardholderName,omitempty"`
			Brand                *string    `json:"brand,omitempty"`
			Number               *string    `json:"number,omitempty"`
			ExpMonth             *string    `json:"expMonth,omitempty"`
			ExpYear              *string    `json:"expYear,omitempty"`
			Code                 *string    `json:"code,omitempty"`
			Title                *any       `json:"title,omitempty"`
			FirstName            *string    `json:"firstName,omitempty"`
			MiddleName           *any       `json:"middleName,omitempty"`
			LastName             *string    `json:"lastName,omitempty"`
			Address1             *string    `json:"address1,omitempty"`
			Address2             *any       `json:"address2,omitempty"`
			Address3             *any       `json:"address3,omitempty"`
			City                 *string    `json:"city,omitempty"`
			State                *string    `json:"state,omitempty"`
			PostalCode           *string    `json:"postalCode,omitempty"`
			Country              *any       `json:"country,omitempty"`
			Company              *any       `json:"company,omitempty"`
			Email                *string    `json:"email,omitempty"`
			Phone                *any       `json:"phone,omitempty"`
			Ssn                  *any       `json:"ssn,omitempty"`
			PassportNumber       *any       `json:"passportNumber,omitempty"`
			LicenseNumber        *any       `json:"licenseNumber,omitempty"`
			Match                *any       `json:"match,omitempty"` // needed for item.login.uri templates
			URI                  *string    `json:"uri,omitempty"`   // needed for item.login.uri templates
			// can't figure out item-collections;
			// Template should be []strings for this, which conflicts with the other fields
			ExternalID any `json:"externalId,omitempty"`
			Groups     []struct {
				ID            string `json:"id,omitempty"`
				ReadOnly      bool   `json:"readOnly"`
				HidePasswords bool   `json:"hidePasswords"`
			} `json:"groups,omitempty"`
		} `json:"template"`
	} `json:"data"`
}

type Login struct {
	Uris []struct {
		Match any    `json:"match"`
		URI   string `json:"uri"`
	} `json:"uris"`
	Username             *string    `json:"username"`
	Password             *string    `json:"password"`
	Totp                 *string    `json:"totp"`
	PasswordRevisionDate *time.Time `json:"passwordRevisionDate"`
}

type SecureNote struct {
	Type *int `json:"type"`
}

type Card struct {
	CardholderName *string `json:"cardholderName"`
	Brand          *string `json:"brand"`
	Number         *string `json:"number"`
	ExpMonth       *string `json:"expMonth"`
	ExpYear        *string `json:"expYear"`
	Code           *string `json:"code"`
}

type Identity struct {
	Title          *any    `json:"title"`
	FirstName      *string `json:"firstName"`
	MiddleName     *any    `json:"middleName"`
	LastName       *string `json:"lastName"`
	Address1       *string `json:"address1"`
	Address2       *any    `json:"address2"`
	Address3       *any    `json:"address3"`
	City           *string `json:"city"`
	State          *string `json:"state"`
	PostalCode     *string `json:"postalCode"`
	Country        *any    `json:"country"`
	Company        *any    `json:"company"`
	Email          *string `json:"email"`
	Phone          *any    `json:"phone"`
	Ssn            *any    `json:"ssn"`
	Username       *string `json:"username"`
	PassportNumber *any    `json:"passportNumber"`
	LicenseNumber  *any    `json:"licenseNumber"`
}
