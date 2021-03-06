# Regex3000

Quoting Larry Page: "*Simply the most awesome regex parser in
the world!!*"

Just kidding, this is just a project I did by myself to know
more about how regex parsing works, and to learn Go.

## The parser

This is a bash executable that answers `true` or `false` if a
string matches a regex.

Examples:

```
$ regex3000 "[abc]+" "abc"
true
$ regex3000 "[abc]+" "def"
false
```

It currently supports:

- `+` `*` `?` `|` operators
- Curvy brackets operator, eg: `a{2,4}`
- Parentheses
- Bracket characters sets, eg: `[a-z]`
- Escaping special characters with `\`
- "opposite" sets with `^`, eg `[^a-z]`
- match start & end (`^` & `$`)

Still todo:

- Special sets like `\w`, `\W`, `\d`, ...
- Escaping special chars within brackets (eg: `[a-z\]]`)

The [list of integration test cases](https://github.com/marcbllv/regex3000/blob/master/tests/testcases.txt)
sums up all the supported features.

## Build sources

### Compile

Run:

```go
go build -i -o bin/regex3000 ./cmd/regex3000
```

This will create the `regex3000` executable in subdirectory `bin/`.

### Run tests

Multiple test cases are written down in file `tests/testcases.txt`.
They can be run at once with the following command:

```bash
tests/testcli.sh bin/regex3000
```
