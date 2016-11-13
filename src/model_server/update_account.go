package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/boltdb/bolt"
	"io/ioutil"
	api "json_api"
	"log"
	"net/http"
	"strings"
)

// Helpers /////////////////////////////////////////////////////////////////////

//DEBUG -- Stub for chat creds
func ProvisionNewAcount(account *AccountData) {
	account.Private.Disabled = false
	account.Private.InConversation = false
	account.Private.ChatCreds = "STUB"
}

// Handler /////////////////////////////////////////////////////////////////////

func UpdateAccount(
	serverContext ModelServerContext,
	w http.ResponseWriter,
	r *http.Request,
) {
	log.Printf("UpdateAccount")

	// Parse request data
	reqData := api.AccountRequest{}
	{
		if body, err := ioutil.ReadAll(r.Body); nil != err {
			msg := "Failed reading request body"
			WriteErrorResponse(w, http.StatusBadRequest, msg)
			return
		} else if err := json.Unmarshal(body, &reqData); nil != err {
			msg := "Failed parsing request body"
			WriteErrorResponse(w, http.StatusBadRequest, msg)
			return
		}
	}

	// Do the update
	account := api.AccountResponse{}
	if err := serverContext.DB.Update(
		func(tx *bolt.Tx) error {

			// Look up if account exists already
			accountBucket := tx.Bucket([]byte("accounts"))
			if nil == accountBucket {
				msg := "Could not retrieve 'accounts' bucket. Going down..."
				log.Fatal(msg)
				return errors.New(msg)
			}
			accountData := accountBucket.Get([]byte(reqData.Key))

			// If account exists, extract it
			removedIdtys := []string{}
			addedIdtys := []string{}
			idtyChange := false
			currentAccount := AccountData{}
			newAccount := true
			if nil != accountData {
				newAccount = false
				if err := json.Unmarshal(accountData, &currentAccount); nil != err {
					return err
				}
			}

			// Find added identities
			for _, idty := range reqData.Data.Identities {
				if !stringInSlice(idty, currentAccount.Public.Identities) {
					addedIdtys = append(addedIdtys, idty)
					idtyChange = true
				}
			}

			// Find removed identities
			for _, idty := range currentAccount.Public.Identities {
				if !stringInSlice(idty, reqData.Data.Identities) {
					removedIdtys = append(removedIdtys, idty)
					idtyChange = true
				}
			}

			// Update any necessary identity indexes
			if idtyChange {

				log.Printf("Updating identities: +[%s] -[%s]",
					strings.Join(addedIdtys, ","),
					strings.Join(removedIdtys, ","))

				idtyBucket := tx.Bucket([]byte("identities"))
				if nil == idtyBucket {
					msg := "Failed retrieving identity bucket for Idty to remove!"
					return errors.New(msg)
				}

				// Removes
				for _, idty := range removedIdtys {
					idtyIdx := IdentityIndex{}
					if val := idtyBucket.Get([]byte(idty)); nil == val {
						msg := fmt.Sprintf("Database out of sync. No entry for %s in identities", idty)
						log.Println(msg)
						return errors.New(msg)
					} else if err := json.Unmarshal(val, &idtyIdx); nil != err {
						log.Println("Failed to extract identity index: " + err.Error())
						return err
					} else {
						idtyIdx.Accounts = removeFromSet(reqData.Key, idtyIdx.Accounts)
						if updatedIdty, err := json.Marshal(idtyIdx); nil != err {
							return err
						} else if err := idtyBucket.Put([]byte(idty), updatedIdty); nil != err {
							return err
						}
					}
				}

				// Adds
				for _, idty := range addedIdtys {

					// Add the identity if needed
					AddIdentityTx(idty)(tx)

					// Add to the identity index
					idtyIdx := IdentityIndex{}
					if val := idtyBucket.Get([]byte(idty)); nil == val {
						msg := fmt.Sprintf("Failed to add index for %s correctly", idty)
						log.Println(msg)
						return errors.New(msg)
					} else if err := json.Unmarshal(val, &idtyIdx); nil != err {
						log.Println("Failed to extract identity index: " + err.Error())
						return err
					} else {
						idtyIdx.Accounts = addToSet(reqData.Key, idtyIdx.Accounts)
						if updatedIdty, err := json.Marshal(idtyIdx); nil != err {
							return err
						} else if err := idtyBucket.Put([]byte(idty), updatedIdty); nil != err {
							return err
						}
					}
				}
			}

			// Update the account entry
			currentAccount.Public = reqData.Data
			if newAccount {
				ProvisionNewAcount(&currentAccount)
			}
			if updatedAccount, err := json.Marshal(currentAccount); nil != err {
				log.Println("Failed to marshal new account")
				return err
			} else if err := accountBucket.Put([]byte(reqData.Key), updatedAccount); nil != err {
				log.Println("Failed to write updated account info")
				return err
			}

			// Success
			account.Public = currentAccount.Public
			account.Private = currentAccount.Private
			return nil

		}); err != nil {
		msg := "Failed to update account: " + err.Error()
		WriteErrorResponse(w, http.StatusInternalServerError, msg)

	} else {
		// Successfully updated account. Return the account data.
		WriteSuccessfulResponse(w, account)
	}
}
