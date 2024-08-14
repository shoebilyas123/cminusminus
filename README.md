# CMinusMinus - The Other Side Of the Number Line

Hundreds of hours went into brainstorming that unique name for my language. Phew, who would have thought of that. CMinusMinus is just a dynamically typed, interpreted and a very unoptimized toy language that I used to learn how interpreters work. Currently it's very simple language with REPL playgroud.

### Variables

You can use `let` keyword declare let statements. The supported primitives are `int64` and `boolean`. The interpreter will dynamically assign the types with let.

```
  let x = 4;
```

Everything is an expression in CMM. In the above example, let x = 4; produces a value of 4;

### Functions

You can declare functions with the `fn` keyword.

```
let add = fn(a,b) {return a+b;}
```

### Todo Features

- Replace let with static types.
- Add Strings and character primitives
- Arrays
- Standard i/o functions for cli
- Networking Capabilities

### Installation

The installation process is quite straightforward.

- Just run `go install https://github.com/shoebilyas123/cminusminus`.
- Enter `cmm` to enter the REPL.
