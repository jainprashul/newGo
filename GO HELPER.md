### Go Modules Example - Hello World

We have created a greeting module. We have created a function Hello() which will return a string "Hello, World!" in greetings module

To create a module, we need to run the following command

```bash
go mod init <OrganizationName>/<ModuleName>
```

```bash
go mod init example.com/greetings
```


./greetings/greetings.go

in hello module we have created a main function which will call Hello() function from greetings module and print the output.

./hello/hello.go

To run the code, we need to run the following command

To use the greetings module in hello module, we need to run the following command

```bash
go mod edit -replace example.com/greetings=../greetings
```

```bash
go mod tidy
```


### Go Web Server Example - Hello World

In WebServer module, we have created a web server which will return "Hello, World!" on the browser.
we have created a main function which will start the web server on port 8000.

To run the code, we need to run the following command

```bash
go run web.go
```
