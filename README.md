# go-set
Simple Go sets that work how I want them to work, as opposed to the million other variations

# Usage

Underneath they are simply maps, so acts as first class type.

```go
mySet := set.Set[int]
mySet.Add(1) // Add a single item.
mySet.AddSlice([]int{1,2,3,4}) // Add a slice to the set.
fmt.Println(mySet)
// Outputs [1 2 3 4]

myStringSet := set.FromSlice([]string{"a","b","c"})
myStringSet.Delete("c")
slice := myStringSet.ToSlice()
fmt.Println(slice)
// Outputs [a b]
```
