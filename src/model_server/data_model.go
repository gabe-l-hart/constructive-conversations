package main

import (
	api "json_api"
)

// Internal storage info for an account
type AccountData struct {
	Public  api.AccountPublic  `json:"public"`
	Private api.AccountPrivate `json:"private"`
}

// Internal storage info for an identity index
type IdentityIndex struct {
	Accounts []string `json:"accounts"`
}
