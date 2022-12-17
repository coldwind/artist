#!/bin/sh
export LC_CTYPE=C 
export LANG=C
grep -rl --exclude=install.sh . | xargs sed -i "" 's/ARTIST_PROJECT_NAME/xxx/g'