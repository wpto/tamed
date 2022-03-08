#!/bin/bash


test_upload_one() {
  RES="$(curl -s -X POST http://localhost:1314/api/posts \
    -F "upload[]=@terminal.gif" \
    -H "Content-Type: multipart/form-data")"

  if [[ "$(echo "$RES" | jq '. | length')" -ne "1" ]]; then
    echo "upload one file error. expected length 1, got:"
    echo "$RES" | jq
    exit
  fi

  ID="$(echo "$RES" | jq -r '.[0].id')"

  if [[ "$ID" -ne "" ]]; then 
    echo "upload one. expected id, got empty string:"
    echo "$RES" | jq
    exit
  fi
}

test_upload_two() {
  RESPONSE=$(curl -s -X POST http://localhost:1314/api/posts \
    -F "upload[]=@terminal.gif" \
    -F "upload[]=@dragon_sm.gif" \
    -H "Content-Type: multipart/form-data")

  RESULT=$(echo "$RESPONSE" | jq '. | length')

  if [[ "$RESULT" -ne "2" ]]; then
    echo "upload files: expected 2 items, got"
    echo "$RESPONSE" | jq 
    exit
  fi

  IDS="$(echo "$RESPONSE" | jq -r '.[].id')"
  ID="$(echo "$IDS" | head -n 1)"
  ID2="$(echo "$IDS" | tail -n 1)"
}

test_found() {
  # $1 postID
  RES=$(curl -s -X GET "http://localhost:1314/api/posts/$1")

  if [[ "$(echo "$RES" | jq -r '.id')" -ne "$1" ]]; then
    echo "upload one. expected match id($ID), got:"
    echo "$RES" | jq
    exit
  fi
}

test_not_found() {
  # $1 postID

  TMP="$(mktemp)"
  RES=$(curl -s -o "$TMP" \
    -w "%{http_code}" \
    "http://localhost:1314/api/posts/$1")

  if [[ "$RES" -ne "404" ]]; then
    echo "not found. expected 404, got:"
    echo "status code: $RES"
    cat "$TMP"
    exit
  fi
}

test_delete() {
  # $1 postID
  RES=$(curl -s -X DELETE "http://localhost:1314/api/posts/$1")
  JQ=$(echo "$RES" | jq -r '.ok')

  if [[ "$JQ" -ne "deleted" ]]; then
    echo "delete. expected deleted, got:"
    echo "$RES" | jq
    exit
  fi
}

test_query_length() {
  # $1 expected length
  # $2 limit
  # $3 offset
  RESPONSE=$(curl -s -X GET "http://localhost:1314/api/posts?limit=$2&offset=$3")

  LEN=$(echo "$RESPONSE" | jq -r '.posts | length')

  if [[ "$LEN" -ne "$1" ]]; then
    echo "query length for (lim=$2, off=$3). expected $1, got $LEN:"
    echo $RESPONSE
    exit
  fi
}

test_modify() {
  # $1 postID
  # $2 modify json body

  RESPONSE=$(curl -s -X PATCH "http://localhost:1314/api/posts/$1" \
    --header "Content-Type: application/json" \
    --data "$2")

  JQ=$(echo "$RESPONSE" | jq -r '.ok')

  if [[ "$JQ" -ne "changed" ]]; then
    echo "modify. expected changed, got:"
    echo "$RESPONSE"
    exit
  fi
}

test_tag_length() {
  # $1 postID
  # $2 expected length

  RESPONSE=$(curl -s -X GET "http://localhost:1314/api/posts/$1")

  if [[ "$(echo "$RESPONSE" | jq -r '.id')" -ne "$1" ]]; then
    echo "tag length. expected match id($ID), got:"
    echo "$RESPONSE"
    exit
  fi

  LEN="$(echo "$RESPONSE" | jq -r '.tags | length')"

  if [[ "$LEN" -ne "$2" ]]; then
    echo "tag length. expected $2, got $LEN:"
    echo "$RESPONSE"
    exit
  fi
}

test_tag_has() {
  # $1 postID
  # $2 contain json array
  RESPONSE=$(curl -s -X GET "http://localhost:1314/api/posts/$1")

  if [[ "$(echo "$RESPONSE" | jq -r '.id')" -ne "$1" ]]; then
    echo "tag length. expected match id($ID), got:"
    echo "$RESPONSE"
    exit
  fi

  RES="$(echo "$RESPONSE" | jq -r ".tags | contains($2)")"

  if [[ "$RES" -ne "true" ]]; then
    echo "tag length. expected $2"
    echo "$RESPONSE"
    exit
  fi
}


ID=""
ID2=""

test_upload_one
test_found "$ID"
test_delete "$ID"
test_not_found "$ID"
echo "delete success"

test_upload_two
test_found "$ID"
test_found "$ID2"
test_query_length "2" "10" "0"
test_query_length "1" "1" "0"
test_query_length "1" "1" "1"
test_query_length "0" "1" "2"
test_delete "$ID"
test_not_found "$ID"
test_delete "$ID2"
test_not_found "$ID2"
test_query_length "0" "10" "0"
echo "upload many success"


# test delete on random string

test_upload_one
test_modify "$ID" '{"rm_tags":["hello","world"]}'
test_tag_length "$ID" "0"
test_modify "$ID" '{"add_tags":["hello", "there"]}'
test_tag_length "$ID" "2"
test_tag_has "$ID" '["there","hello"]'
test_modify "$ID" '{"rm_tags":["hello","world"]}'
test_tag_length "$ID" "1"
test_tag_has "$ID" '["there"]'
test_delete "$ID"
echo "modify success"

test_upload_two
test_upload_two
test_upload_two
test_upload_two
test_upload_two
test_upload_two
test_upload_two
test_upload_two
