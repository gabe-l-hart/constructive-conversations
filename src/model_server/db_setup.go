package main

import (
	"encoding/json"
	"errors"
	"github.com/boltdb/bolt"
	"log"
)

// Ensure that a bucket exists
func SetupBucket(name string, serverContext ModelServerContext) error {
	if err := serverContext.DB.Update(
		func(tx *bolt.Tx) error {
			if _, err := tx.CreateBucketIfNotExists([]byte(name)); nil != err {
				log.Printf("Error creating %s bucket: "+err.Error(), name)
				return err
			} else {
				return nil
			}
		}); err != nil {
		log.Printf("Error initializing %s bucket", name)
		return err
	} else {
		return nil
	}
}

// Helper function for adding an identity as part of an open bolt transaction
func AddIdentityTx(idty string) func(tx *bolt.Tx) error {
	return func(tx *bolt.Tx) error {

		// Get the identity bucket
		idBucket := tx.Bucket([]byte("identities"))
		if nil == idBucket {
			msg := "Error looking up 'identities' bucket"
			log.Fatal(msg)
			return errors.New(msg)
		}

		// Add to the bucket if needed
		idtyList := idBucket.Get([]byte(idty))
		if nil == idtyList {
			if newIdty, err := json.Marshal(IdentityIndex{}); nil != err {
				return err
			} else if err := idBucket.Put([]byte(idty), newIdty); nil != err {
				return err
			}
		}

		return nil
	}
}

// Add an Identity to the database if it's not already there
func AddIdentityIfNeeded(idty string, serverContext ModelServerContext) error {
	if err := serverContext.DB.Update(AddIdentityTx(idty)); err != nil {
		log.Printf("Error adding identity [%s]: "+err.Error(), idty)
		return err
	} else {
		return nil
	}
}
