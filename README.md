# CMinusMinus - The Other Side Of the Number Line

Hundreds of hours went into brainstorming that unique name for my language. Phew, who would have thought of that. CMinusMinus is just a dynamically typed, interpreted and a very unoptimized toy language that I used to learn how interpreters work. Currently it's very simple language with REPL playgroud.


## Working
The working of the interpreter will give a brief idea about how it works and it's architecture:
- Lexer: The input code is fed into the lexer that performs lexical analysis of our code and outputs our code as a list of small workable data structures called tokens.
- Parser: The output from the lexer is picked up by the parser that parses our code and generates an abstract syntax tree. We have achieved this by implementing a pratt parser.
- Pratt Parser: The appropriate parsing function is associated with very node in our AST, depending whether the token is found in a prefix or an infix expression.
- Evaluation: Traverse the AST, visit each node and do what the node signifies. It's called tree-walking interpreter.
- Object System: Every value in our code is an `Object`. Each value in our environment is wrapped inside a struct which fufills this `Object` interface. We have used an object system to represent the internal values instead of primitive types.

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
### REPL
You can exit the REPL by using `exit()` command.

### Todo Features

- Replace let with static types.
- Add Strings and character primitives
- Arrays
- Standard i/o functions for cli
- Networking Capabilities

### Installation
 - Clone the repository.
 - Navigate into the source code and run `./build.sh`.
 - If shows permission errors: `chmod a+x ./build.sh`.
 - Now run ./build.sh again the go will build the code binary in the `bin/` directory.
 - Run ./bin/cminusminus and you will enter the REPL. 
