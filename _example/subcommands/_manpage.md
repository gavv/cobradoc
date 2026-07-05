```
EXAMPLE(1)                            Example Manual                           EXAMPLE(1)

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
       example main hello [flags]
              Say hello

              Command flags:

              -c, --count=1
                     how many times to greet?

       example main bye [flags]
              Say goodbye

              Command flags:

              -n, --name=John
                     who got the goodbye?

DOCUMENTATION COMMANDS
       example docs man [flags]
              Generate manual page

       example docs markdown [flags]
              Generate markdown page

REPORTING BUGS
       Please report bugs at https://github.com/gavv/cobradoc

Example Manual                           Jul 2026                              EXAMPLE(1)
```
