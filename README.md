# go-set
Simple Go sets that work how I want them to work, as opposed to the million other variations

# Usage

Underneath they are simply maps, so acts as first class type.

```go
mySet := set.Set[int]{}
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

# Any type
Sets work with any comparable types, include structs.
```go
type MyStruct struct {
   num int
   str string
}

a := MyStruct{num: 5, str: "thing"}
b := MyStruct{num: 5, str: "thing"}
c := MyStruct{num: 1, str: "other"}
mySet := make(set.Set[MyStruct])
mySet.Add(a)
mySet.Add(b)
mySet.Add(c)
// Set contains only two items, since a and b compare as identical.
```


# Set operations

Supports Intersection, Union, and Difference

## Difference
```go
// Returns the difference a - b That is to say:
a := MakeSet([]int{1,2,3,4,5})
b := MakeSet([]int{1,2,3})
d := a.Difference(b)
// d contains 4,5
## Union
```go
a := MakeSet([]int{1,2,3,4,5})
b := MakeSet([]int{1,2,3,10,11})
u := a.Union(b)
// u contains 1,2,3,4,5,10,11
```
## Intersection
Everything that is in both 'a' and b'
```go
a := set.FromSlice([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
b := set.FromSlice([]int{6, 7, 8, 9, 10, 20, 30, 40})
i := a.Intersection(b)
// i contains 6, 7, 8, 9, 10
```
