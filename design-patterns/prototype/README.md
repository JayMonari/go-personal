# Prototype

This design pattern allows for creation of objects from previously instantiated
objects. The idea is to avoid the cost of the creation of new objects and to
instead clone a previously instantiated object.

When speaking of costly operations, we can think of operations that are highly
cpu bound or io bound, such as gathering values from a database or needing a
myriad of services to build the object. Cloning an object from already obtained
values is a great way to save time and resources. This is much like caching.

## Actors

* Prototype: `interface` that the clonable objects should implement.
* Concrete Prototype: `struct` or `type` that implements the Prototype.
* Client: Creates the clone from the Concrete Prototype.
