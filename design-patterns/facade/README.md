# Facade

Hides away the complexity of a system by providing a simple `interface` for the
client. Unifying all of the different subsystems and their interfaces into a
single `interface` that is easy to use and interact with the entire system.

e.g. The gas pedal of a car allows you to easily accelerate in a forwards and
backwards manner all at the push of a pedal, but think of all the moving parts
that create the car as a whole in order for that to be achieved. Fine
engineering, indeed.

## Example

The task is to build a process to allow a user to make a comment on a blog.

To accomplish this, the system needs to execute several processes:

- Validate that the user is logged in and active.
- Validate that the user has permission to comment.
- Post the comment.
- Notify the creator of the blog that a new comment has arrived.

This would be a lot of processes for the client to do by themselves, to which
they would never actually try to comment if all of that was required of them.
We can hide away all of these processes using a facade.

- **POST COMMENT**
