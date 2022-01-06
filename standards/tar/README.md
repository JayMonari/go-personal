# archive/tar

## Required Knowledge

- fs
- io

## Background

Tape archives (tar) are a file format for storing a sequence of files that can
be read and written in a streaming manner. A common way of grouping files on
UNIX-like systems is known as a "tarball", tarfile or a tape archive (tar). It
is usually compressed with gzip.

### Examples

- `recovery.tar.gz`
- `favorite-cats.tar`
- `shopping_receipts.tar`

## Dilemma

You are dealing with a file or group of files with the `.tar` file extension,
what would be the easiest way to interact with this tarball? There is always
the possibility of calling the `tar` command found on Linux and Mac, inside of
our Go program with something like `exec.Command("tar", "-czf")` but the
implementations for each differ and may vary in results each time. Not only
that but we may only want to see the contents of the tarball and pick what we
want from it or add one thing on condition of a certain file, this would
become very cumbersome, very quickly.

## Package's Solution

A way to interact with a tarfile, whether that be creating, appending, or
reading from the tarfile. It accomplishes this with the `tar.NewReader` and
`tar.NewWriter` functions which will return their respective types.
