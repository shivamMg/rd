# rd [![godoc](https://godoc.org/github.com/shivammg/rd?status.svg)](https://godoc.org/github.com/shivamMg/rd) [![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

`rd` is a small library to build hand-written recursive descent parsers. Besides exposing convenient methods to parse tokens it features automatic parse tree generation and flow tracing for debugging.

Recursive descent parsers can imitate their grammar quite well. For instance, for the following grammar:

```
A → aB
B → b
```

(where `A` and `B` are non-terminals, and `a` and `b` are terminals), a recursive descent parser might look like:

```go
func A() bool {
    if nextToken() == "a" {
        return B()
    }
    pushback("a")
    return false
}

func B() bool {
    if nextToken() == "b" {
        return true
    }
    pushback("b")
    return false
}
```

Parser for the same grammar written using `rd`:

```go
import "github.com/shivamMg/rd"

func A(b *rd.Builder) (ok bool) {
    b.Enter("A")
    defer b.Exit(&ok)

    return b.Match("a") && B(b)
}

func B(b *rd.Builder) (ok bool) {
    defer b.Enter("B").Exit(&ok)

    return b.Match("b")
}
```

A builder object keeps track of the current token and exposes convenient methods to write a parser. `Match` method, for instance, matches a token to the current token, and resets the original state in case of an unsuccessful match; there's no need for a manual `pushback`. As non-terminal functions are called, terminal matches, and calls to other non-terminal functions are done. A parse tree is generated in the builder object using these matches and calls:

In case of a successful match the terminal is added to parse tree under the current non-terminal (the one in which `Match` was called). Same goes in case of a non-terminal function call: the non-terminal, if it exits successfully, is added to parse tree under the current non-terminal. You can imagine this process being repeated recursively.

Argument to `Enter` is what is added to parse tree as a symbol for the non-terminal. Argument to `Exit` determines if function call was successful or not.

`ParseTree` method returns the parse tree which is, well, a tree data structure. It can be pretty-printed.

```go
tokens := []rd.Token{"a", "b"}
b := rd.NewBuilder(tokens)
if ok := A(b); ok {
    fmt.Print(b.ParseTree())
}
```

The above snippet will print:

```
A
├─ a
└─ B
   └─ b
```

A debug tree is also maintained which, unlike the parse tree, contains all matches and calls (not just the successful ones). It's helpful if you want to debug a parsing failure. It can be retrieved using the `DebugTree` method.

```go
fmt.Print(b.DebugTree())
```

The above snippet will print:

```
A(true)
├─ a = a
└─ B(true)
   └─ b = b
```

Parsing errors can be retrieved using `Err` method. For `tokens := []rd.Token{"a", "c"}` the following statements:

```go
fmt.Println(b.Err())
fmt.Print(b.DebugTree())
```

will print:

```
parsing error
A(false)
├─ a = a
└─ B(false)
   └─ c ≠ b
```


## Examples

### [Arithmetic expression parser](examples/arithmetic)

```bash
go get github.com/shivamMg/rd/examples/arithmetic   # requires go modules support (go1.11+)
arithmetic -expr='3.14*4*(6/3)'  # hopefully $GOPATH/bin is in $PATH
arithmetic -expr='3.14*4*(6/3)' -backtrackingparser
```

Parser and grammar for it can be found inside `examples/arithmetic/parser`. There's another parser written for a different grammar that also parses arithmetic expressions. This parser can be found inside `examples/arithmetic/backtrackingparser`. It uses backtracking - notice the use of `b.Backtrack()`.

This example lexer built using [chroma](https://github.com/alecthomas/chroma).


### [PL/0 programming language parser](examples/pl0)

```
go get github.com/shivamMg/rd/examples/pl0
cd examples/pl0/
pl0 square.pl0
pl0 multiply.pl0
pl0 prime.pl0
```

Parser and grammar can be found inside `examples/pl0/parser`. Grammar has been taken from [en.wikipedia.org/wiki/PL/0#Grammar](https://en.wikipedia.org/wiki/PL/0#Grammar). It also uses lexer built using chroma.

### [Domain name parser](examples/domainname)

```
go get github.com/shivamMg/rd/examples/domainname
domainname www.google.co.uk
```

Grammar has been taken from [www.ietf.org/rfc/rfc1035.txt](https://www.ietf.org/rfc/rfc1035.txt). The lexer is hand-written.

## Licence

MIT

## Contribute

Contribute through bug fixes, improvements, new examples, etc. Lucky PR submitters get to walk home with a brand new CRT and an Audi 1987.

