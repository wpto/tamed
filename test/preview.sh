#!/bin/bash
source "./curl.sh"


ID=""
ID2=""


for FILE in $(ls art/); do
  filename="${FILE%.*}"

  test_upload_one "art/$FILE"

  for TAG in $(echo "$filename" | sed -r 's/_/\n/g'); do
    BODY='{"add_tags":["'
    BODY+="$TAG"
    BODY+='"]}'
    echo $BODY
    test_modify "$ID" "$BODY"
  done
  sleep 0.5
done