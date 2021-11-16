# Builder

Allows for the separation of the construction of a `struct`
(complete or partial) from its representation. Using the same building process
many different representations can be created.

* Product: Main `struct` built. Represents the `struct` under construction.
* Builder: defines the interface with the methods that should be implemented by
the constructors.
* Concrete Builder: Implementation of the Builder to make concrete `struct`.
* Director: Constructs the `struct` the interface.
