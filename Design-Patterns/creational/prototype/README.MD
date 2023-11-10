# What is the Prototype pattern ? 
- Prototype pattern is used when the creation of a new object is costly than copying the props from an existing one, when the number of props are alot and actually a bunch of them is the same for all objects, we should consider cloning objects by implementing the `Clone()` method in our concrete class class.

- Prototype is used when for example in your GUI application, you created the first blue character with all its complexities of creation, and then you need to create red character with same props but the color is different, so its simplier to use cloning.


# What is the Prototype solving ? 
- The prototype is not the same as assigning objects to other variables.

- Prototype is all about creating `deep-clones` to our object, so the new created (cloned) object will not have the same reference as the old one. and thats how we avoid bugs.

# How prototype is implemented in golang ? 

first we define a clonable interface to be implemented by any class we need to apply the prototype creation pattern on it 
```go
type Cloneable interface {
	Clone() Cloneable
}
```

then we define the class (struct) that implements the clonable interface
```go
type Student struct {
    Name string
}

func (s *Student) Clone () Clonable {
    return &Student{
        Name : s.Name
    }
}
```

**HINT**: to support the deep-copy not the shallow copy, we should marshel and unmarshel the object 
```go
// apply the marshling and unmarshling in the Clone() concrete impl
func (s *Student) Clone() Clonable{
    // first we marshel the object to be cloned into bytes 
    sbytes, err := json.Marshal(s)
    if err != nil {
        panic(err)
    }

    // then we unmarshal the bytes back to the object type 
    var clonedStudent *Student 
    err = json.Unmarshal(sbytes, clonedStudent) 
    if err != nil {
        panic(err)
    }

    // finally we return the cloned object 
    return *clonedStudent
}
``` 