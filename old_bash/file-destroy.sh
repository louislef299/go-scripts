#!/bin/bash                                                                                                                                                                                                 

if [[ -z "$1" ]]; then
    echo "usage: file-destroy <file>"
    exit 1
fi

echo "Start!"
while read p; do
    terraform destroy -target "${p}"
done <$1
