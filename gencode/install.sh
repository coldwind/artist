#!/bin/sh
find . -type f -exec sed -i -e 's/foo/xxx/g' {} \;