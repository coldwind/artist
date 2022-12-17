#!/bin/sh
export LC_CTYPE=C 
export LANG=C
if [ $# -eq 1 ]; then
    grep -rl --exclude=install.sh . | xargs sed -i "" "s/ARTIST_PROJECT_NAME/${1}/g"
else
    echo "项目名不能为空."
fi
