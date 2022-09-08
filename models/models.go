// Account represents an account in the form3 org section.
// See https://api-docs.form3.tech/api.html#organisation-accounts for
// more information about fields.
package models

import "gorm.io/gorm"

type AccountData struct {
	gorm.Model
	Attributes     *AccountAttributes `json:"attributes,omitempty" gorm:"not null"`
	ID             string             `json:"id,omitempty"  gorm:"primaryKey"`
	OrganisationID string             `json:"organisation_id,omitempty" gorm:"unique; not null"`
	Type           string             `json:"type,omitempty" gorm:"not null"`
	Version        int64             `json:"version,omitempty" gorm:"not null"`
}

type AccountAttributes struct {
	gorm.Model
	AccountClassification   string  `json:"account_classification,omitempty" gorm:"not null"`
	AccountMatchingOptOut   bool    `json:"account_matching_opt_out,omitempty" gorm:"not null"`
	AccountNumber           string   `json:"account_number,omitempty" gorm:"not null"`
	AlternativeNames        []string `json:"alternative_names,omitempty" gorm:"not null"`
	BankID                  string   `json:"bank_id,omitempty" gorm:"primaryKey"`
	BankIDCode              string   `json:"bank_id_code,omitempty" gorm:"unique; not null"`
	BaseCurrency            string   `json:"base_currency,omitempty"  gorm:"unique; not null"`
	Bic                     string   `json:"bic,omitempty" gorm:"not null"`
	Country                 string  `json:"country,omitempty" gorm:"not null"`
	Iban                    string   `json:"iban,omitempty" gorm:"not null"`
	JointAccount            bool    `json:"joint_account,omitempty" gorm:"not null"`
	Name                    []string `json:"name,omitempty" gorm:"not null"`
	SecondaryIdentification string   `json:"secondary_identification,omitempty" gorm:"not null"`
	Status                  string  `json:"status,omitempty" gorm:"not null"`
	Switched                bool    `json:"switched,omitempty" gorm:"not null"`
}