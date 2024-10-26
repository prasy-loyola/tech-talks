#/bin/bash

if [ "$#" -ne 2 ]; then
    echo "Usage: $(basename $0) <source-filename> <output-file-name>"
    exit 1
fi

source=$1
output=$2

go build -o tlangc
./tlangc c $source > .intermediate.asm
fasm .intermediate.asm $output > /dev/null
chmod a+x $output
# rm .intermediate.asm
