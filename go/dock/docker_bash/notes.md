# Official notes from Louis

## Priority one: learn how to use flags in bash

The solution is pretty simple and I found a command pretty quick online: [How to get arguments in flags bash](https://stackoverflow.com/questions/7069682/how-to-get-arguments-with-flags-in-bash). The article recommends using the `getopts` built-in command in order to gather information about flags. The implementation seems to be very similar to the go implementation, but I am going to read a bit more about it before I start using it. 

Here is a more in-depth description of how the `getopts` command differs from the underlying `getopt` posix-based system command: [learn more about `getopts`](https://www.computerhope.com/unix/bash/getopts.htm). After reading a bit more about the `getopts` function, we will only be able to define flags with a single character. Not a big deal being that writing out the entire flag can become tedious.

## Priority two: How to set up our dock commands in bash

This will work differently than how we were setting up our go function in that everything will live in one file. This kinda sucks for readability, but it will have to do. [Here is an article](https://linuxize.com/post/bash-functions/#passing-arguments-to-bash-functions) that outlines how to use bash functions. Passing parameters to functions works differently from other languages than python, but it pretty similar to how we gathered our commands before using flags, i.e. `$1`. Functions have to exist before they are called in the file however, so we will define all of our functions before the brains of the `dock` command begin to parse the flags and run functions.