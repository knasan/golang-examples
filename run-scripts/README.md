# run-scripts

This example runs executable files and checks their status.

- scripts/s1.sh
  - script with exit code 0
- scripts/s2.sh
  - script with exit code 1
- scripts/s3.sh
  - script has no execution bit (permission denied)

output:

```sh
    path: scripts file: s1.sh
    path: scripts file: s2.sh
    exit status 1
    path: scripts file: s3.sh
    fork/exec scripts/s3.sh: permission denied
```
