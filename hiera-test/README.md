# hiera test

Hiera test is very complex and not easy to understand, so let me explain that a little more.

The structure of the YAML or JSON file is declared

```go
    type Config struct ...
```

In the main function, the struct is first assigned to a variable and the Load function is called.

In the Load function, the Workaround function is called; this is necessary to read in all YAML files and to write
them to /tmp/hiera-test so that all anchors that may be set in the YAML files are released.
The reason for this is because the hiera library cannot handle anchors in YAML.

Then the function configHelper is called, which manipulates each section via the function makeFirstCharToLowerCase
the passed string in such a way that the first character is capitalized.
A lookup is started from here. This result is loaded into the stuct via a case block.

The prepareString function removes all spaces that might be contained and removes any break in the string.

Although most of the work is done here, there is still a lot of logic to do around it, especially when using YAML files with anchors and so on.
Unfortunately here still uses YAML.v2 instead of v3.

INFO [0000] Load
INFO [0000] workaround
INFO [0000] writeWorker
...
INFO [0000] configHelper
INFO [0000] makeFirstCharToLowerCase
INFO [0000] prepareString
...
