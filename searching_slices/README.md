#Background
I've seen a few places where to search for a string in an slice, a practice of joining the slice and 
calling strings.Contains against it has been used.

Something like this:
```go
func joinContains(slice []string, s string) bool {
	return strings.Contains(strings.Join(slice, "|"), s)
}
```

It strikes me that while this may be succinct code, surely it would be more efficient to just loop 
the slice and return true on the first instance of an equal string. Not to mention that simply 
checking whether a substring exists could be... sloppy?

My thinking was that unless the strings package is doing some kind of super magical stuff to do the
join and contains operations, you'd be going through at least two iterations of looping. Turns out
(inspecting the package), it's actually up to 4 sets of loops.

So the benchmarks here were written to check my bias.

#Results:
The results speak for themselves, and bear out the 4 loop operations, not to mention the additional
conditional checks that are done in strings.

The benchmarks check a worst case scenario (i.e. the element does not exist at all in the slice).

For completeness, given that there exist cases where you *would* want to check a substring, there is
a "compromise" benchmark which loops a slice, but instead of simple equality calls `strings.Contains`

Small: 100 elements
Medium: 1,000,000 elements
Large: 100,000,000 elements

Run 1
```
Benchmark_JoinASmallSliceAndCheckContains
Benchmark_JoinASmallSliceAndCheckContains-4        	  960590	      1155 ns/op
Benchmark_LoopOverASmallSliceAndCheckEquality
Benchmark_LoopOverASmallSliceAndCheckEquality-4    	 2959987	       412 ns/op
Benchmark_JoinAMediumSliceAndCheckContains
Benchmark_JoinAMediumSliceAndCheckContains-4       	      94	  12199270 ns/op
Benchmark_LoopOverAMediumSliceAndCheckEquality
Benchmark_LoopOverAMediumSliceAndCheckEquality-4   	     279	   4147637 ns/op
Benchmark_JoinALargeSliceAndCheckContains
Benchmark_JoinALargeSliceAndCheckContains-4        	       1	46215546918 ns/op
Benchmark_LoopOverALargeSliceAndCheckEquality
Benchmark_LoopOverALargeSliceAndCheckEquality-4    	       1	1002904178 ns/op
```