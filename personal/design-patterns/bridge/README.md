# Bridge

Allows for separation of behavior of the representations
(abstraction of the implementation)

Requires an abstract class and an interface. This can be done with interfaces
alone in Go.

This allows the combination of behavior with representations, using the least
amount of `structs`.

## Compare to Inheritance

| Representations | Implementations | Inheritance | Bridge |
| --- | --- | --- | --- |
| 2 | 2 | 4 | 4 |
| 3 | 2 | 6 | 5 |
| 4 | 4 | 16 | 8 |
| 5 | 4 | 20 | 9 |
