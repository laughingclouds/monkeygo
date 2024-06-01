# Monkey Programming Language

Major parts:

- lexer
- parser
- Abstract Syntax Tree (AST)
- internal object system
- evaluator

higher order function
functions as first class citizens
.
.
.
## Lexing

For the eventual conversion of *source code* into an *abstract syntax tree* we first convert it into *tokens*. Lexing is the process for doing so.

> Lexer/Tokenizer for *lexing* and Parser for turning tokens into AST

### Token

The **Token** data structure may have a type and a literal associated to it.

We can distinguish between tokens using the type, meanwhile the literal will allow us to know what value it holds.

> All token types are defined in token/token.go


### Lexer

Take source code as input and output the tokens that represent the source code. We will go to each token one by one using a ``NextToken()`` function call.

> In a production setting, it's common to attach file name and line number to the **Token**. I wonder if it's as easy as adding extra fields into the type declaration?

As of now, the lexer will only support ASCII for simplicity's sake. We can add support for unicode later.

## Parsing

*Conversion of input into a data structure that represents the input*. [wikipedia](https://en.wikipedia.org/wiki/Parsing#Parser).
A good example of a parser is the `JSON` parser in JavaScript.

```js
> var input = '{"name": "Hemant", "age":21}';
> var output = JSON.parse(input);
> output
{name; 'Hemant', age: 21}
> output.name
'Hemant'
> output.age
21
```

A *string* `input` is taken by the parse and an *object* `output` is returned. The same happens with a parser for a programming language. But what we're representing in the data structure is the language itself.