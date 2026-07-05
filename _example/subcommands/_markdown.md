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

### Main Commands

* [example main](#example-main)
* [example main hello](#example-main-hello)
* [example main bye](#example-main-bye)

### Documentation Commands

* [example docs](#example-docs)
* [example docs man](#example-docs-man)
* [example docs markdown](#example-docs-markdown)

# Main Commands

## `example main`

Run main commands

```text
example main [flags]
```

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

# Documentation Commands

## `example docs`

Generate documentation

```text
example docs [flags]
```

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
