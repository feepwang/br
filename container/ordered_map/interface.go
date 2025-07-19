//go:build !go1.23
// +build !go1.23

package ordered_map

import (
	"cmp"

	"github.com/feepwang/br/container/pair"
)

type Interface[K cmp.Ordered, V any] interface {
	Len() int
	Cap() int
	Get(key K) (V, bool)
	GetMutable(key K) (*V, bool)
	Set(key K, value V)
	// SetReturn(key K, value V) (old V, replaced bool)
	Delete(key K) bool
	// DeleteReturn(key K) (old V, deleted bool)
	Has(key K) bool

	Keys() []K
	Values() []V
	Pairs() []pair.Pair[K, V]
}
