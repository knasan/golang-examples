# unlink

unlink tests whether a link to the file is still valid.

If you start this example, the output "not exists" comes first.

Create any file and link it to "filelink".

```sh
    touch anyfilename
    ln -s anyfilename filelink
```

If you start it again now, the output should read "exists".

Now delete the file, not the link. It is checked whether the link still points to a valid file.

```sh
    rm anyfilename
```

If you start again, the output "not exists" comes.
