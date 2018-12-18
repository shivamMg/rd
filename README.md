# rd

`rd` is a small library to build hand-written recursive descent parsers. Besides exposing convenient methods to parse tokens it features automatic parse tree generation and flow tracing for debugging.

Recursive descent parsers can imitate their grammar quite well. e.g. for the grammar:

```
A → aB
B → b
```

where `A` and `B` are non-terminals, and `a` and `b` are terminals, a recursive descent parser might look like:

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

`Match` method conveniently resets original state in case of failure terminal matches. `Enter` and `Exit` methods serve the purpose of marking entry and exit respectively from the non-terminal functions. Argument to `Enter`, e.g. `"A"`, is what will show up in the parse tree. The exit result (`&ok`) determines if the generated subtree for the non-terminal must be added to the parse tree. This subtree is created considering successful terminal matches and non-terminals calls done inside the non-terminal.

A debug tree is also generated that contains all non-terminal calls and terminal matches - unlike the parse tree that contains only successful ones. Debug tree is helpful if you want to debug a parsing failure.

Both parse trees and debug trees satisfy `fmt.Stringer` and can be pretty printed.

```go
tokens := []rd.Token{"a", "b"}
b := rd.NewBuilder(tokens)
if ok := A(b); ok {
	fmt.Println(b.ParseTree())
} else {
	fmt.Println(b.Err())
}
fmt.Println(b.DebugTree())
```

The above snippet will print:

```
A
├─ a
└─ B
   └─ b

A(true)
├─ a = a
└─ B(true)
   └─ b = b
```

For `tokens := []rd.Token{"a", "c"}` the snippet will print:

```
parsing error
A(false)
├─ a = a
└─ B(false)
   └─ c ≠ b
```

In the debug tree you can notice both successful and unsuccessful terminal matches, and non-terminal exit results.


## Examples

### Arithmetic expression parser

```bash
go get github.com/shivamMg/rd/examples/arithmetic   # requires go modules support (go1.11+)
arithmetic -expr='3.14*4*(6/3)'  # hopefully $GOPATH/bin is in $PATH
arithmetic -expr='3.14*4*(6/3)' -backtrackingparser
```

Parser and grammar can be found inside `examples/arithmetic/parser`. There's another parser written for a different grammar that also parses arithmetic expressions. This parser can be found inside `examples/arithmetic/backtrackingparser`. It uses backtracking - notice the use of `b.Backtrack()`. This example uses [chroma](https://github.com/alecthomas/chroma) for lexing.


### PL/0 programming language parser

```
go get github.com/shivamMg/rd/examples/pl0
cd examples/pl0/
pl0 square.pl0
pl0 multiply.pl0
pl0 prime.pl0
```

Parser and grammar can be found inside `examples/pl0/parser`. Grammar has been taken from [en.wikipedia.org/wiki/PL/0#Grammar](https://en.wikipedia.org/wiki/PL/0#Grammar). It also uses chroma for lexing.

### Domain name parser

```
go get github.com/shivamMg/rd/examples/domainname
domainname www.google.co.uk
```

Grammar has been taken from [www.ietf.org/rfc/rfc1035.txt](https://www.ietf.org/rfc/rfc1035.txt). Its lexer is hand-written.

