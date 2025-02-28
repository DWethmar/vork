package persistence

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"sync"

	"github.com/dwethmar/vork/component"
	"go.etcd.io/bbolt"
)

type Store[T component.Component] interface {
	Get(id uint) (T, error)
	Add(c T) (uint, error)
}

type Persistence[T component.Component] struct {
	mut     sync.Mutex
	changed map[uint]struct{}
	deleted map[uint]struct{}
	t       component.Type
	store   Store[T]
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

func New[T component.Component](t component.Type, store Store[T]) *Persistence[T] {
	return &Persistence[T]{
		mut:     sync.Mutex{},
		changed: make(map[uint]struct{}),
		deleted: make(map[uint]struct{}),
		t:       t,
		store:   store,
	}
}

func (p *Persistence[T]) Changed(id uint, deleted bool) error {
	p.mut.Lock()
	defer p.mut.Unlock()
	if deleted {
		delete(p.changed, id)
		p.deleted[id] = struct{}{}
	} else {
		p.changed[id] = struct{}{}
	}
	return nil
}

func (p *Persistence[T]) Commit(tx *bbolt.Tx) error {
	p.mut.Lock()
	defer p.mut.Unlock()
	for id := range p.changed {
		c, err := p.store.Get(id)
		if err != nil {
			return fmt.Errorf("failed to get component: %w", err)
		}
		// Serialize the component using gob
		var buf bytes.Buffer
		enc := gob.NewEncoder(&buf)
		if err := enc.Encode(c); err != nil {
			return fmt.Errorf("failed to encode component: %w", err)
		}

		// Create or get the bucket for the component type
		bucket, err := tx.CreateBucketIfNotExists([]byte(p.t))
		if err != nil {
			return fmt.Errorf("failed to create or get bucket: %w", err)
		}

		// Use the component ID as the key
		id := c.ID()
		if err = bucket.Put(itob(id), buf.Bytes()); err != nil {
			return fmt.Errorf("failed to save component: %w", err)
		}
	}
	for id := range p.deleted {
		bucket := tx.Bucket([]byte(p.t))
		if bucket == nil {
			return fmt.Errorf("bucket not found for type %s", p.t)
		}
		if err := bucket.Delete(itob(id)); err != nil {
			return fmt.Errorf("failed to delete component: %w", err)
		}
	}
	p.changed = make(map[uint]struct{})
	p.deleted = make(map[uint]struct{})
	return nil
}

func (p *Persistence[T]) Load(tx *bbolt.Tx) error {
	bucket := tx.Bucket([]byte(p.t))
	if bucket == nil {
		return fmt.Errorf("bucket not found for type %s", p.t)
	}
	return bucket.ForEach(func(k, v []byte) error {
		// Decode the component using gob
		dec := gob.NewDecoder(bytes.NewReader(v))
		var c T
		if err := dec.Decode(&c); err != nil {
			return fmt.Errorf("failed to decode component: %w", err)
		}
		if _, err := p.store.Add(c); err != nil {
			return fmt.Errorf("failed to add component: %w", err)
		}
		return nil
	})
}
