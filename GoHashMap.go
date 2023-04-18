package main

import (
    "fmt"
    "reflect"
)

type node struct {
    key   interface{}
    value interface{}
    next  *node
}

type hashmap struct {
    buckets []*node
    size    int
}

func newHashmap() *hashmap {
    buckets := make([]*node, 16)
    return &hashmap{buckets: buckets, size: 0}
}

func (h *hashmap) getBucketIndex(key interface{}) int {
    hash := 0
    switch reflect.TypeOf(key).Kind() {
    case reflect.String:
        for _, ch := range key.(string) {
            hash = hash*31 + int(ch)
        }
    default:
        hash = reflect.ValueOf(key).Pointer()
    }
    return hash % len(h.buckets)
}

func (h *hashmap) Put(key interface{}, value interface{}) {
    index := h.getBucketIndex(key)
    node := &node{key: key, value: value}

    if h.buckets[index] == nil {
        h.buckets[index] = node
    } else {
        current := h.buckets[index]
        for current.next != nil {
            if reflect.DeepEqual(current.key, key) {
                current.value = value
                return
            }
            current = current.next
        }
        if reflect.DeepEqual(current.key, key) {
            current.value = value
        } else {
            current.next = node
        }
    }
    h.size++
}

func (h *hashmap) Get(key interface{}) interface{} {
    index := h.getBucketIndex(key)

    current := h.buckets[index]
    for current != nil {
        if reflect.DeepEqual(current.key, key) {
            return current.value
        }
        current = current.next
    }

    return nil
}

func main() {
    hm := newHashmap()
    hm.Put("foo", "bar")
    hm.Put([]int{1, 2, 3}, map[string]int{"a": 1, "b": 2})

    fmt.Println(hm.Get("foo"))         // Output: bar
    fmt.Println(hm.Get([]int{1, 2, 3})) // Output: map[a:1 b:2]
    fmt.Println(hm.Get("missing"))     // Output: nil
}

/*
In this implementation, the node struct represents a key-value pair, and a hashmap struct holds an array of buckets, 
each of which is a linked list of nodes. The Put method adds a new key-value pair to the hashmap or updates the value 
of an existing key, while the Get method retrieves the value associated with a given key.

The getBucketIndex function uses a switch statement to handle different types of keys. If the key is a string, 
it uses a simple hash function to map the key to an index in the hashmap's array of buckets. If the key is a complex 
data type (such as a slice or map), it uses the Pointer() method of the reflect package to get a unique identifier for the key.

Note that this implementation uses the reflect package, which can have a performance impact, so it may not be suitable 
for use in performance-critical code. If performance is a concern, you may want to consider using a more specific data 
type for your keys (e.g. a custom struct) and defining your own hash function.
*/
