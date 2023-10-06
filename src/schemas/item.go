package schemas

import (
	"time"

	"github.com/google/uuid"
)

type CommonFields struct {
	OrganizationID *uuid.UUID `json:"organizationId,omitempty"`
	CollectionIds  *uuid.UUID `json:"collectionIds,omitempty"`
	FolderID       *uuid.UUID `json:"folderId,omitempty"`
	Type           *int       `json:"type,omitempty"`
	Name           *string    `json:"name,omitempty"`
	Notes          *string    `json:"notes,omitempty"`
	Favorite       *bool      `json:"favorite,omitempty"`
	Fields         *[]any     `json:"fields,omitempty"`
}

type Item struct {
	CommonFields
	Login      *any `json:"login,omitempty"`
	SecureNote *any `json:"secureNote,omitempty"`
	Card       *any `json:"card,omitempty"`
	Identity   *any `json:"identity,omitempty"`
	Reprompt   *int `json:"reprompt,omitempty"`
	Uris       []struct {
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
	ExternalID           any        `json:"externalId,omitempty"`
	Groups               []struct {
		ID            string `json:"id,omitempty"`
		ReadOnly      bool   `json:"readOnly"`
		HidePasswords bool   `json:"hidePasswords"`
	} `json:"groups,omitempty"`
}

type Template struct {
	Success bool `json:"success"`
	Data    struct {
		Object   string `json:"object"`
		Template struct {
			CommonFields
			Reprompt *int `json:"reprompt,omitempty"`
			Uris     []struct {
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
			ExternalID           any        `json:"externalId,omitempty"`
			Groups               []struct {
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
	CommonFields
	Title          *any    `json:"title,omitempty"`
	FirstName      *string `json:"firstName,omitempty"`
	MiddleName     *any    `json:"middleName,omitempty"`
	LastName       *string `json:"lastName,omitempty"`
	Address1       *string `json:"address1,omitempty"`
	Address2       *any    `json:"address2,omitempty"`
	Address3       *any    `json:"address3,omitempty"`
	City           *string `json:"city,omitempty"`
	State          *string `json:"state,omitempty"`
	PostalCode     *string `json:"postalCode,omitempty"`
	Country        *any    `json:"country,omitempty"`
	Company        *any    `json:"company,omitempty"`
	Email          *string `json:"email,omitempty"`
	Phone          *any    `json:"phone,omitempty"`
	Ssn            *any    `json:"ssn,omitempty"`
	PassportNumber *any    `json:"passportNumber,omitempty"`
	LicenseNumber  *any    `json:"licenseNumber,omitempty"`
}

type ListVaultItemsResponse struct {
	Success bool `json:"success"`
	Data    struct {
		Object string `json:"object"`
		Data   []struct {
			Object *string `json:"object"`
			ID     *string `json:"id"`
			CommonFields
			Login         *Login      `json:"login,omitempty"`
			CollectionIds *[]any      `json:"collectionIds"`
			RevisionDate  *time.Time  `json:"revisionDate"`
			CreationDate  *time.Time  `json:"creationDate"`
			DeletedDate   *any        `json:"deletedDate"`
			Identity      *Identity   `json:"identity,omitempty"`
			Card          *Card       `json:"card,omitempty"`
			SecureNote    *SecureNote `json:"secureNote,omitempty"`
		} `json:"data"`
	} `json:"data"`
}
