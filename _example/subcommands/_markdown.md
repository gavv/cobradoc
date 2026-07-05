# Example Manual

Example command using cobra and cobradoc.

```text
example [command] [global flags] [command flags]
```

### Global Flags

```text
  -v, --value string   string value
      --flag           boolean flag
```

### Commands
* [example main](#example-main)
  * [example main hello](#example-main-hello)
  * [example main bye](#example-main-bye)
* [example docs](#example-docs)
  * [example docs man](#example-docs-man)
  * [example docs markdown](#example-docs-markdown)


# Commands

## `example main`

Run main commands

### `example main hello`

Say hello

```text
example main hello [flags]
```

#### Command Flags

```text
  -c, --count int   how many times to greet? (default 1)
```

### `example main bye`

Say goodbye

```text
example main bye [flags]
```

#### Command Flags

```text
  -n, --name string   who got the goodbye? (default "John")
```

## `example docs`

Generate documentation

### `example docs man`

Generate manual page

```text
example docs man [flags]
```

### `example docs markdown`

Generate markdown page

```text
example docs markdown [flags]
```

# Reporting Bugs

Please report bugs at https://github.com/gavv/cobradoc
