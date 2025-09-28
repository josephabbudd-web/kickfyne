clear
set -e # break on error.

go build >&tools/logs/build.log
mv ./kickfyne ~/bin
