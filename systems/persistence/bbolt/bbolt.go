package bbolt

import (
	"bytes"
	"encoding/gob"
	"fmt"

	"github.com/dwethmar/vork/component"
	bolt "go.etcd.io/bbolt"
)

type Repository struct {
	db *bolt.DB
}

// itob converts a uint32 ID to a byte slice for use as a key in BoltDB.
func itob(v uint32) []byte {
	return []byte{
		byte(v >> 24),
		byte(v >> 16),
		byte(v >> 8),
		byte(v),
	}
}

// Save saves a component in its respective bucket, encoded using gob.
func (r *Repository) Save(c component.Component) error {
	// Get the component type to determine the bucket name
	t := c.Type()

	// Serialize the component using gob
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	if err := enc.Encode(c); err != nil {
		return fmt.Errorf("failed to encode component: %w", err)
	}

	// Save the component in the bucket
	return r.db.Update(func(tx *bolt.Tx) error {
		// Create or get the bucket for the component type
		bucket, err := tx.CreateBucketIfNotExists([]byte(t))
		if err != nil {
			return fmt.Errorf("failed to create or get bucket: %w", err)
		}

		// Use the component ID as the key
		id := c.ID()
		if err := bucket.Put(itob(id), buf.Bytes()); err != nil {
			return fmt.Errorf("failed to save component: %w", err)
		}

		return nil
	})
}

// Delete removes a component from its respective bucket by ID.
func (r *Repository) Delete(t string, id uint32) error {
	// Delete the component by ID from its corresponding bucket
	return r.db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(t))
		if bucket == nil {
			return fmt.Errorf("bucket %s not found", t)
		}

		// Delete the component by its ID
		if err := bucket.Delete(itob(id)); err != nil {
			return fmt.Errorf("failed to delete component: %w", err)
		}

		return nil
	})
}

// List retrieves all components of a given type from their bucket.
func (r *Repository) List(t string) ([]component.Component, error) {
	var components []component.Component

	// Fetch all components from the bucket
	err := r.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(t))
		if bucket == nil {
			return fmt.Errorf("bucket %s not found", t)
		}

		// Iterate over all components in the bucket
		return bucket.ForEach(func(k, v []byte) error {
			var c component.Component

			// Decode the gob-encoded component
			buf := bytes.NewBuffer(v)
			dec := gob.NewDecoder(buf)
			if err := dec.Decode(&c); err != nil {
				return fmt.Errorf("failed to decode component: %w", err)
			}

			components = append(components, c)
			return nil
		})
	})

	if err != nil {
		return nil, err
	}

	return components, nil
}

func NewRepository(db *bolt.DB) *Repository {
	return &Repository{
		db: db,
	}
}
