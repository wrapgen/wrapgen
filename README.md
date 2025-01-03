# wrapgen

Generates go code based on go interfaces.

It comes with a set of embedded [templates](./template) but, you can always provide you own templates
to extend it for other use cases.

Compared to gowrap or mockgen it is significantly faster in generating code, especially in large projects.

## Status

Experimental. Use at your own risk. Do not expect active support.
The tool is developed mainly for internal use. Non-conflicting contributions are welcome.
Only a small fraction of the gomock test-coverage is transferred, there may be bugs.

## Usage

Annotate an interface with `//wrapgen:generate` to activate code generation.

Example:

```go
package test

//wrapgen:generate -template prometheus -destination main_gen.go
//wrapgen:generate -template gomock -destination main_gen.go
type Reader interface {
	Read() []int
}

//wrapgen:generate -template gomock -destination main_gen.go -name Writer
type _ interface {
	Write(data []int) error
}
```

- `//wrapgen:generate` statements within in one file may to write to the same file.
- The code generated code in the output files is written in the order as the generate statements appear in the file.
- `-template` specifies the template. If it contains a slash then it's interpreted as path to a template file,
  otherwise it's treated as name of a built-in template.
- `-name` overwrites the name of the interface, as if the interface would have been defined with a different name.
- `-vars` can be used to inject arbitrary values into the templates.

## Differences to other tools

Unlike [gomock/mockgen](https://github.com/uber-go/mock) and [gowrap](https://github.com/hexdigest/gowrap) it operates
on the whole project at once, this way more similar to [buf](https://buf.build/).
Due to this difference it can utilize multiple cores on the computer as it doesn't depend on the serial
processing of `//go:generate` statements. However, there is no import-dependency tracking implemented, it's up to the
user to not create dependencies between interfaces in generated code and interfaces from which code is generated.

Similar to gowrap it can create arbitrary code based on templates.

Unlike mockgen it's not bound to a particular use-case/library. Also, it supports only "source-mode" not "
reflection-mode".

Unlike gowrap imports are handled explicitly and `golang.org/x/tools/cmd/goimports` is not invoked to generate the
`import (..)` section of the file.

### Limitations

- dot-imports are not supported. They are ignored and interfaces referencing types from dot-imports are not
  correctly generated. Supporting dot-imports correctly would require parsing those packages to properly reference
  symbols in those packages and distinguish them from symbols defined in the interfaces' own package.
- vendored dependencies are not tested/supported.
- Only go versions with support for generics are supported.

## Builtin templates

- `gomock` generates code similar to gomock/mockgen with -typed.
- `gomock_untyped` generates code similar to gomock/mockgen without -typed.
- `moq` generates code similar to https://github.com/matryer/moq.
- `prometheus` generates an generic prometheus metric middleware.

## Tests

The majority of test-coverage is through examples which the test compares to the expected output.

The expected output can be (re)generated by just running wrapgen against itself: `go run .`.

## Authors

Wrapgen is licensed under the Apache License, Version 2.0.

Wrapgen is based on code from gomock v0.4.0 https://github.com/uber-go/mock. See AUTHORS.gomock.
