# Official notes from Louis

## Priority one: learn how to use flags in bash

The solution is pretty simple and I found a command pretty quick online: [How to get arguments in flags bash](https://stackoverflow.com/questions/7069682/how-to-get-arguments-with-flags-in-bash). The article recommends using the `getopts` built-in command in order to gather information about flags. The implementation seems to be very similar to the go implementation, but I am going to read a bit more about it before I start using it. 

Here is a more in-depth description of how the `getopts` command differs from the underlying `getopt` posix-based system command: [learn more about `getopts`](https://www.computerhope.com/unix/bash/getopts.htm). After reading a bit more about the `getopts` function, we will only be able to define flags with a single character. Not a big deal being that writing out the entire flag can become tedious.