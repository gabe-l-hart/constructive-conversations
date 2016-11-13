package main

import (
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

// Add an Identity to the database if it's not already there
func AddIdentityIfNeeded(idty string, serverContext ModelServerContext) error {
	if err := serverContext.DB.Update(
		func(tx *bolt.Tx) error {

			// If the bucket doesn't already exist, create it
			idBucket := tx.Bucket([]byte(idty))
			if nil == idBucket {
				if _, err := tx.CreateBucket([]byte(idty)); nil != err {
					log.Printf("Error adding bucket [%s]: "+err.Error(), idty)
					return err
				}
			}
			return nil
		}); err != nil {
		log.Printf("Error adding identity [%s]: "+err.Error(), idty)
		return err
	} else {
		return nil
	}
}
