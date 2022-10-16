set -e

go build
./slideconvert -base $PWD -content ../../docs/
