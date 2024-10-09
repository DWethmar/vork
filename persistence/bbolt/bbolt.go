package bbolt

import (
	"bytes"
	"encoding/gob"
	"fmt"

	"github.com/dwethmar/vork/component"
	bolt "go.etcd.io/bbolt"
)

// Repository is a generic repository for managing a specific component type.
type Repository[T component.Component] struct {
	factory func() T // Factory function for creating new instances of T
}

// NewRepository creates a new repository for a specific component type.
func NewRepository[T component.Component](factory func() T) *Repository[T] {
	return &Repository[T]{
		factory: factory,
	}
}

// itob converts a uint ID to a byte slice for use as a key in BoltDB.
func itob(v uint) []byte {
	return []byte{
		byte(v >> 24),
		byte(v >> 16),
		byte(v >> 8),
		byte(v),
	}
}

// Save saves a component of type T in its respective bucket, encoded using gob.
// The transaction must be passed as the first argument.
func (r *Repository[T]) Save(tx *bolt.Tx, c T) error {
	// Get the component type to determine the bucket name
	t := c.Type()

	// Serialize the component using gob
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	if err := enc.Encode(c); err != nil {
		return fmt.Errorf("failed to encode component: %w", err)
	}

	// Create or get the bucket for the component type
	bucket, err := tx.CreateBucketIfNotExists([]byte(t))
	if err != nil {
		return fmt.Errorf("failed to create or get bucket: %w", err)
	}

	// Use the component ID as the key
	id := c.ID()
	if err = bucket.Put(itob(id), buf.Bytes()); err != nil {
		return fmt.Errorf("failed to save component: %w", err)
	}

	return nil
}

// Get retrieves a specific component of type T by its ID using the provided transaction.
func (r *Repository[T]) Get(tx *bolt.Tx, id uint) (T, error) {
	// Create a new instance of the component using the factory
	c := r.factory()

	// Fetch the component by ID
	bucket := tx.Bucket([]byte(c.Type()))
	if bucket == nil {
		return c, fmt.Errorf("bucket %q not found", c.Type())
	}

	// Retrieve the component by ID
	v := bucket.Get(itob(id))
	if v == nil {
		return c, fmt.Errorf("component with ID %d not found", id)
	}

	// Decode the gob-encoded component into the new instance
	buf := bytes.NewBuffer(v)
	dec := gob.NewDecoder(buf)
	if err := dec.Decode(c); err != nil {
		return c, fmt.Errorf("failed to decode component: %w", err)
	}

	return c, nil
}

// Delete removes a component of type T by its ID using the provided transaction.
func (r *Repository[T]) Delete(tx *bolt.Tx, id uint) error {
	// Fetch the component type using the factory to get the bucket name
	c := r.factory()
	bucketName := c.Type()

	// Delete the component by ID from its corresponding bucket
	bucket := tx.Bucket([]byte(bucketName))
	if bucket == nil {
		return fmt.Errorf("bucket %q not found", bucketName)
	}

	// Delete the component by its ID
	if err := bucket.Delete(itob(id)); err != nil {
		return fmt.Errorf("failed to delete component: %w", err)
	}

	return nil
}

// List retrieves all components of type T from their bucket using the provided transaction.
func (r *Repository[T]) List(tx *bolt.Tx) ([]T, error) {
	var components []T

	// Fetch the component type using the factory to get the bucket name
	bucketName := r.factory().Type()

	// Fetch all components from the bucket
	bucket := tx.Bucket([]byte(bucketName))
	if bucket == nil {
		return components, nil
	}

	// Iterate over all components in the bucket
	err := bucket.ForEach(func(_, v []byte) error {
		// Create a new instance of the component
		c := r.factory()

		// Decode the gob-encoded component
		buf := bytes.NewBuffer(v)
		dec := gob.NewDecoder(buf)
		if err := dec.Decode(c); err != nil {
			return fmt.Errorf("failed to decode component: %w", err)
		}

		components = append(components, c)
		return nil
	})

	if err != nil {
		return nil, err
	}

	return components, nil
}
