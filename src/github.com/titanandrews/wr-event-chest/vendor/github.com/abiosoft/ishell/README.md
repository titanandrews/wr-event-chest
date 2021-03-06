# ishell
ishell is an interactive shell library for creating interactive cli applications.

[![Documentation](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](https://godoc.org/github.com/abiosoft/ishell)
[![Go Report Card](https://goreportcard.com/badge/github.com/abiosoft/ishell)](https://goreportcard.com/report/github.com/abiosoft/ishell)

## Older version
The current master is not backward compatible with older version. Kindly change your import path to `gopkg.in/abiosoft/ishell.v1`.

Older version of this library is still available at [https://gopkg.in/abiosoft/ishell.v1](https://gopkg.in/abiosoft/ishell.v1).

However, you are advised to upgrade.

## Usage


```go
import "strings"
import "github.com/abiosoft/ishell"

func main(){
    // create new shell.
    // by default, new shell includes 'exit', 'help' and 'clear' commands.
    shell := ishell.New()

    // display welcome info.
    shell.Println("Sample Interactive Shell")

    // register a function for "greet" command.
    shell.AddCmd(&ishell.Cmd{
        Name: "greet",
        Help: "greet user",
        Func: func(c *ishell.Context) {
            c.Println("Hello", strings.Join(c.Args, " "))
        },
    })

    // start shell
    shell.Start()
}
```
Execution
```
Sample Interactive Shell
>>> help

Commands:
  clear      clear the screen
  greet      greet user
  exit       exit the program
  help       display help

>>> greet Someone Somewhere
Hello Someone Somewhere
>>> exit
$
```

### Reading input.
```go
// simulate an authentication
shell.AddCmd(&ishell.Cmd{
    Name: "login",
    Help: "simulate a login",
    Func: func(c *ishell.Context) {
        // disable the '>>>' for cleaner same line input.
        c.ShowPrompt(false)
        defer c.ShowPrompt(true) // yes, revert after login.

        // get username
        c.Print("Username: ")
        username := c.ReadLine()

        // get password.
        c.Print("Password: ")
        password := c.ReadPassword()

        ... // do something with username and password

        c.Println("Authentication Successful.")
    },
})
```
Execution
```
>>> login
Username: someusername
Password:
Authentication Successful.
```

### Multiline input.
Builtin support for multiple lines.
```
>>> This is \
... multi line

>>> Cool that << EOF
... everything here goes
... as a single argument.
... EOF
```
User defined
```go
shell.AddCmd(&ishell.Cmd{
    Name: "multi",
    Help: "input in multiple lines",
    Func: func(c *ishell.Context) {
        c.Println("Input multiple lines and end with semicolon ';'.")
        lines := c.ReadMultiLines(";")
        c.Println("Done reading. You wrote:")
        c.Println(lines)
    },
})
```
Execution
```
>>> multi
Input multiple lines and end with semicolon ';'.
>>> this is user defined
... multiline input;
You wrote:
this is user defined
multiline input;
```
### Keyboard interrupt.
Builtin interrupt handler.
```
>>> ^C
Input Ctrl-C once more to exit
>>> ^C
Interrupted
exit status 1
```
Custom
```go
shell.Interrupt(func(count int, c *ishell.Context) { ... })
```

### Progress Bar
Determinate
```go
func(c *ishell.Context) {
    c.ProgressBar().Start()
    for i := 0; i < 101; i++ {
        c.ProgressBar().Suffix(fmt.Sprint(" ", i, "%"))
        c.ProgressBar().Progress(i)
        ... // some background computation
    }
    c.ProgressBar().Stop()
}
```
Output
```
[==========>         ] 50%
```

Indeterminate
```go

func(c *ishell.Context) {
    c.ProgressBar().Indeterminate(true)
    c.ProgressBar().Start()
    ... // some background computation
    c.ProgressBar().Stop()
}
```
Output
```
[ ====               ]
```

Custom display using [briandowns/spinner](https://github.com/briandowns/spinner).
```go
display := ishell.ProgressDisplayCharSet(spinner.CharSets[11])
func(c *Context) { c.ProgressBar().Display(display) ... }

// or set it globally
ishell.ProgressBar().Display(display)
```

### Durable history.
```go
// Read and write history to $HOME/.ishell_history
shell.SetHomeHistoryPath(".ishell_history")
```

### Example
Available [here](https://github.com/abiosoft/ishell/blob/master/example/main.go).
```sh
go run example/main.go
```

## Supported Platforms
* [x] Linux
* [x] OSX
* [x] Windows [Not tested but should work]

## Note
ishell is in active development and can still change significantly.

## Roadmap (in no particular order)
* [x] Multiline inputs.
* [x] Command history.
* [x] Customizable tab completion.
* [x] Handle ^C interrupts.
* [x] Subcommands and help texts.
* [x] Scrollable paged output.
* [x] Progress bar.
* [ ] Testing, testing, testing.

## Contribution
1. Create an issue to discuss it.
2. Send in Pull Request.

## License
MIT

## Credits
Library | Use
------- | -----
[github.com/flynn-archive/go-shlex](http://github.com/flynn-archive/go-shlex) | splitting input into command and args.
[gopkg.in/readline.v1](http://gopkg.in/readline.v1) | history, tab completion and reading passwords.

## Donate
```
bitcoin:1GTHYEDiy2C7RzXn5nY4wVRaEN2GvLjwZN
```
