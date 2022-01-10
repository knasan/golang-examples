# package reference to package

It can happen that you need a function that you export so often that you turn in a circle and
thus you get the error message "circle import".

This example shows how this can be prevented. It's actually quite easy once you understand it.

A struct is defined in the "config" package that contains everything you want to take with you to the next package.
In the main function of the application, a variable for Config struct is created, which can then be transferred to every package.
Every package that has to accept this struct is written as a small function.

```go
    var app * config.AppConfig
    func NewAppConfig (a * config.AppConfig) {
    app = a
    }
```

You can now access the struct in the package as usual.

You only have to transfer the struct to the package once via the NewAppConfig function.
Since it is a pointer, this only has to be done once in the application. Changes are recognized immediately.
