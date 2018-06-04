# go-hacking
A demo repository for my adventures hacking around in Go to get familiar with the language.

## How to Run:

1. Spin up the latest GoLang Docker Container:

```
$ sudo docker run -it -v $PWD:/go golang:latest /bin/bash
```

2. From inside the Docker Container, use standard go commands such as `go run`, `go build`, or `go clean`:

```
$ go run hack.go
Hello World!  I am a Go Program!
=========================================
Also - Hello World; I am another line of code!
The Environment Variable is not defined. Using default value:  TestValue
The SETTING1 configuration option is:  TestValue
The value of my variable is  2
The incremented value of my variable is 3
```

3. You can set the environment variable, `SETTING1` to equal any value and observe the output:

```
$ export SETTING1="Hello Go!"
$ go run hack.go
Hello World!  I am a Go Program!
=========================================
Also - Hello World; I am another line of code!
The Environment Variable is defined. The value is:  Hello Go!
The SETTING1 configuration option is:  Hello Go!
The value of my variable is  2
The incremented value of my variable is 3
```
4. CLI Arguments can be sent using the `--number` flag and `--text` flag after building the projec to an executable:

```
$ go build hack.go
$ chmod +x hack
$ ./hack --number 8 --text "Hello World!"
Hello World!  I am a Go Program!
=========================================
Also - Hello World; I am another line of code!
The Environment Variable is not defined. Using default value:  TestValue
The CLI option for '--number' is:  8  -  The provided value is less than 10!
The CLI option for '--text' is:  Hello World!
The SETTING1 configuration option is:  TestValue
The value of my variable is  2
The incremented value of my variable is 3
```
