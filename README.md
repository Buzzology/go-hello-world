# Go Hello world tutorial
https://golang.org/doc/tutorial/getting-started

## Useful commands 
Install modules: `go mod tidy`  
Run tests: `go test` or `go test -v` for verbose  

### Building and installing
Generate binary executable: `go build`  
Executing on Windows: `.\hello.exe`  
Executing on linux: `./hello`  
Get go install path: `go list -f '{{.Target}}'` 
Add go install directory to path: `set PATH="%PATH%;C:\Users\Chris\go\bin\hello.exe"`  
Change install directory: `$ go env -w GOBIN=C:\path\to\your\bin`
Install package: `go install` _from the hello directory_

### Dependencies
https://golang.org/doc/modules/managing-dependencies  

### TODO: Investigate
#### Greetings order not preserved
Go maps don't preserve order. Normally use a slice or array instead. See https://stackoverflow.com/a/28931555/522859.