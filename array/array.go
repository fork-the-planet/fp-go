package array

import (
	G "github.com/ibm/fp-go/array/generic"
	F "github.com/ibm/fp-go/function"
	"github.com/ibm/fp-go/internal/array"
	M "github.com/ibm/fp-go/monoid"
	O "github.com/ibm/fp-go/option"
	"github.com/ibm/fp-go/tuple"
)

// From constructs an array from a set of variadic arguments
func From[A any](data ...A) []A {
	return G.From[[]A](data...)
}

// MakeBy returns a `Array` of length `n` with element `i` initialized with `f(i)`.
func MakeBy[A any](n int, f func(int) A) []A {
	// sanity check
	if n <= 0 {
		return Empty[A]()
	}
	// run the generator function across the input
	as := make([]A, n)
	for i := n - 1; i >= 0; i-- {
		as[i] = f(i)
	}
	return as
}

// Replicate creates a `Array` containing a value repeated the specified number of times.
func Replicate[A any](n int, a A) []A {
	return MakeBy(n, F.Constant1[int](a))
}

func MonadMap[A, B any](as []A, f func(a A) B) []B {
	return G.MonadMap[[]A, []B](as, f)
}

func MonadMapRef[A, B any](as []A, f func(a *A) B) []B {
	count := len(as)
	bs := make([]B, count)
	for i := count - 1; i >= 0; i-- {
		bs[i] = f(&as[i])
	}
	return bs
}

func Map[A, B any](f func(a A) B) func([]A) []B {
	return F.Bind2nd(MonadMap[A, B], f)
}

func MapRef[A, B any](f func(a *A) B) func([]A) []B {
	return F.Bind2nd(MonadMapRef[A, B], f)
}

func filter[A any](fa []A, pred func(A) bool) []A {
	var result []A
	count := len(fa)
	for i := 0; i < count; i++ {
		a := fa[i]
		if pred(a) {
			result = append(result, a)
		}
	}
	return result
}

func filterRef[A any](fa []A, pred func(a *A) bool) []A {
	var result []A
	count := len(fa)
	for i := 0; i < count; i++ {
		a := fa[i]
		if pred(&a) {
			result = append(result, a)
		}
	}
	return result
}

func filterMapRef[A, B any](fa []A, pred func(a *A) bool, f func(a *A) B) []B {
	var result []B
	count := len(fa)
	for i := 0; i < count; i++ {
		a := fa[i]
		if pred(&a) {
			result = append(result, f(&a))
		}
	}
	return result
}

func Filter[A any](pred func(A) bool) func([]A) []A {
	return F.Bind2nd(filter[A], pred)
}

func FilterRef[A any](pred func(*A) bool) func([]A) []A {
	return F.Bind2nd(filterRef[A], pred)
}

func MonadFilterMap[A, B any](fa []A, f func(a A) O.Option[B]) []B {
	return G.MonadFilterMap[[]A, []B](fa, f)
}

func FilterMap[A, B any](f func(a A) O.Option[B]) func([]A) []B {
	return G.FilterMap[[]A, []B](f)
}

func FilterMapRef[A, B any](pred func(a *A) bool, f func(a *A) B) func([]A) []B {
	return func(fa []A) []B {
		return filterMapRef(fa, pred, f)
	}
}

func reduceRef[A, B any](fa []A, f func(B, *A) B, initial B) B {
	current := initial
	count := len(fa)
	for i := 0; i < count; i++ {
		current = f(current, &fa[i])
	}
	return current
}

func Reduce[A, B any](f func(B, A) B, initial B) func([]A) B {
	return func(as []A) B {
		return array.Reduce(as, f, initial)
	}
}

func ReduceRef[A, B any](f func(B, *A) B, initial B) func([]A) B {
	return func(as []A) B {
		return reduceRef(as, f, initial)
	}
}

func Append[A any](as []A, a A) []A {
	return G.Append(as, a)
}

func IsEmpty[A any](as []A) bool {
	return array.IsEmpty(as)
}

func IsNonEmpty[A any](as []A) bool {
	return len(as) > 0
}

func Empty[A any]() []A {
	return G.Empty[[]A]()
}

func Zero[A any]() []A {
	return Empty[A]()
}

// Of constructs a single element array
func Of[A any](a A) []A {
	return G.Of[[]A](a)
}

func MonadChain[A, B any](fa []A, f func(a A) []B) []B {
	return array.Reduce(fa, func(bs []B, a A) []B {
		return append(bs, f(a)...)
	}, Zero[B]())
}

func Chain[A, B any](f func(a A) []B) func([]A) []B {
	return F.Bind2nd(MonadChain[A, B], f)
}

func MonadAp[B, A any](fab []func(A) B, fa []A) []B {
	return MonadChain(fab, F.Bind1st(MonadMap[A, B], fa))
}

func Ap[B, A any](fa []A) func([]func(A) B) []B {
	return F.Bind2nd(MonadAp[B, A], fa)
}

func Match[A, B any](onEmpty func() B, onNonEmpty func([]A) B) func([]A) B {
	return func(as []A) B {
		if IsEmpty(as) {
			return onEmpty()
		}
		return onNonEmpty(as)
	}
}

func Tail[A any](as []A) O.Option[[]A] {
	return G.Tail(as)
}

func Head[A any](as []A) O.Option[A] {
	return G.Head(as)
}

func First[A any](as []A) O.Option[A] {
	return G.First(as)
}

func Last[A any](as []A) O.Option[A] {
	return G.Last(as)
}

func PrependAll[A any](middle A) func([]A) []A {
	return func(as []A) []A {
		count := len(as)
		dst := count * 2
		result := make([]A, dst)
		for i := count - 1; i >= 0; i-- {
			dst--
			result[dst] = as[i]
			dst--
			result[dst] = middle
		}
		return result
	}
}

func Intersperse[A any](middle A) func([]A) []A {
	prepend := PrependAll(middle)
	return func(as []A) []A {
		if IsEmpty(as) {
			return as
		}
		return prepend(as)[1:]
	}
}

func Intercalate[A any](m M.Monoid[A]) func(A) func([]A) A {
	concatAll := ConcatAll[A](m)(m.Empty())
	return func(middle A) func([]A) A {
		return Match(m.Empty, F.Flow2(Intersperse(middle), concatAll))
	}
}

func Flatten[A any](mma [][]A) []A {
	return MonadChain(mma, F.Identity[[]A])
}

func Slice[A any](low, high int) func(as []A) []A {
	return array.Slice[[]A](low, high)
}

func Lookup[A any](idx int) func([]A) O.Option[A] {
	return G.Lookup[[]A](idx)
}

func UpsertAt[A any](a A) func([]A) []A {
	return G.UpsertAt[[]A](a)
}

func Size[A any](as []A) int {
	return G.Size(as)
}

func MonadPartition[A any](as []A, pred func(A) bool) tuple.Tuple2[[]A, []A] {
	return G.MonadPartition(as, pred)
}

// Partition creates two new arrays out of one, the left result contains the elements
// for which the predicate returns false, the right one those for which the predicate returns true
func Partition[A any](pred func(A) bool) func([]A) tuple.Tuple2[[]A, []A] {
	return G.Partition[[]A](pred)
}

// IsNil checks if the array is set to nil
func IsNil[A any](as []A) bool {
	return array.IsNil(as)
}

// IsNonNil checks if the array is set to nil
func IsNonNil[A any](as []A) bool {
	return array.IsNonNil(as)
}

// ConstNil returns a nil array
func ConstNil[A any]() []A {
	return array.ConstNil[[]A]()
}
