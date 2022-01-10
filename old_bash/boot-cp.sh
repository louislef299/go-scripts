#!/bin/bash

## This command is to recover files that were previously booted

if [[ -z $1 ]]; then
    echo "Usage: boot-cp <file>"
fi

if [[ "$pwd/$1" == "/Users/lefebl4/.bash/$1" ]]; then
    echo "This works"
fi

