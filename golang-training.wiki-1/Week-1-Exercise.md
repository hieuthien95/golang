# Exercise 

We will consider to porting [Underscore](https://github.com/jashkenas/underscore) in Javascript language to Go language

# Most features

Some features such as `IsEmpty`, `Last` ... are missed in Go so we will try first

## Helpers
###  IsEmpty

We can consider some of case:
- `slice`, `map` with `len` is zero
- `string` is `""`
- `boolean` is false
- `number` is 0
- ...

So the simplest code could implemented is:

```go
func IsEmpty(obj interface{}) bool {
  if obj == nil || obj == "" || obj == false || obj == 0 {
    return true
  }
  return false
}

func TestIsEmpty(t *testing.T) {
  table := []interface{}{false, "", 0, nil}
  for _, v := range table {
    expected := IsEmpty(v)
    if expected != true {
      t.Error("IsEmpty of \"", v, "\" is failed.")
    }
  }
}
```

## Slice

More stricks of slice you can follow [here](https://github.com/golang/go/wiki/SliceTricks)

### Last

For just reading the last element of a slice:

```go
sl[len(sl)-1]
```

> So we really don't need to reinvent the wheel for this, but in other way to think about readable code, we could consider this function

```go
func Last(arr interface{}) interface{} {
  value := reflect.ValueOf(arr)

  return value.Index(value.Len() - 1).Interface()
}

func TestLast(t *testing.T) {
  arr := []int{1, 2, 3}
  v := Last(arr)
  var expected int
  expected = 3
  if v != expected {
    t.Error("Value should be ", expected)
  }
}
```

