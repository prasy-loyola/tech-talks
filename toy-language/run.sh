#/bin/bash

go build -o tlang-comp 
./tlang-comp > .intermediate.asm
fasm .intermediate.asm $1
chmod a+x $1
./$1
