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

* [example hello](#example-hello)
* [example bye](#example-bye)

### Documentation Commands

* [example man](#example-man)
* [example markdown](#example-markdown)

### Additional Commands

* [example help](#example-help)

# Main Commands

## `example hello`

Say hello

```text
example hello [flags]
```

### Command Flags

```text
  -c, --count int   how many times to greet? (default 1)
  -h, --help        help for hello
```

## `example bye`

Say goodbye

```text
example bye [flags]
```

### Command Flags

```text
  -n, --name string   who got the goodbye? (default "John")
  -h, --help          help for bye
```

# Documentation Commands

## `example man`

Generate manual page

```text
example man [flags]
```

### Command Flags

```text
  -h, --help   help for man
```

## `example markdown`

Generate markdown page

```text
example markdown [flags]
```

### Command Flags

```text
  -h, --help   help for markdown
```

# Additional Commands

## `example help`

Help about any command

```text
example help [command] [flags]
```

### Command Flags

```text
  -h, --help   help for help
```
