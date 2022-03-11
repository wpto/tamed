#!/bin/bash

git branch -D testbranch
git checkout -b testbranch
npm run build --prefix client
git add client/public/build -f
git commit -m 'build client'
heroku config:set FS_LOCAL_PATH=.cache/local FS_MEDIA_PATH=.cache/mediacontent FS_POSTDB_PATH=.cache/post.json PG_RESET=true
git push -f heroku testbranch:main

