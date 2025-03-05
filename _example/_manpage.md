```
EXAMPLE(1)                                      Example Manual                                     EXAMPLE(1)

NAME
       example - Example command

SYNOPSIS
       example [command] [global flags] [command flags]

DESCRIPTION
       Example command using cobra and cobradoc.

       Global flags:

       -v, --value=string
              string value

       --flag boolean flag

MAIN COMMANDS
       example hello [flags]
              Say hello

              Command flags:

              -c, --count=1
                     how many times to greet?

              -h, --help
                     help for hello

       example bye [flags]
              Say goodbye

              Command flags:

              -n, --name=John
                     who got the goodbye?

              -h, --help
                     help for bye

DOCUMENTATION COMMANDS
       example man [flags]
              Generate manual page

              Command flags:

              -h, --help
                     help for man

       example markdown [flags]
              Generate markdown page

              Command flags:

              -h, --help
                     help for markdown

ADDITIONAL COMMANDS
       example help [command] [flags]
              Help about any command

              Command flags:

              -h, --help
                     help for help

REPORTING BUGS
       Please report bugs at https://github.com/gavv/cobradoc

Example Manual                                     Mar 2025                                        EXAMPLE(1)
```
