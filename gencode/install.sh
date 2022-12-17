#!/bin/sh
export LC_CTYPE=C 
export LANG=C
if [ $# -eq 2 ]; then
    `rm -rf $2/$1`
    `cp -rf ../gencode $2/$1`
    `rm -f $2/$1/install.sh`
    `cd $2/$1 && grep -rl --exclude=install.sh . | xargs sed -i "" "s/ARTIST_PROJECT_NAME/${1}/g"`
else
    echo "项目名和项目路径不能为空."
fi
