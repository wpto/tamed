#!/bin/bash

API_ADDR="http://192.168.1.40:1314"

export ID
export ID2

test_upload_one() {
  if [ "$1" == "" ]; then
    UPLOAD_FILE="terminal.gif"
  else
    UPLOAD_FILE="$1"
  fi


  RES="$(curl -s -X POST "$API_ADDR/api/posts" \
    -F "upload[]=@$UPLOAD_FILE" \
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

export -f test_upload_one

test_upload_two() {
  RESPONSE=$(curl -s -X POST "$API_ADDR/api/posts" \
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
export -f test_upload_two

test_found() {
  # $1 postID
  RES=$(curl -s -X GET "$API_ADDR/api/posts/$1")

  if [[ "$(echo "$RES" | jq -r '.id')" -ne "$1" ]]; then
    echo "upload one. expected match id($ID), got:"
    echo "$RES" | jq
    exit
  fi
}
export -f test_found

test_not_found() {
  # $1 postID

  TMP="$(mktemp)"
  RES=$(curl -s -o "$TMP" \
    -w "%{http_code}" \
    "$API_ADDR/api/posts/$1")

  if [[ "$RES" -ne "404" ]]; then
    echo "not found. expected 404, got:"
    echo "status code: $RES"
    cat "$TMP"
    exit
  fi
}
export -f test_not_found

test_delete() {
  # $1 postID
  RES=$(curl -s -X DELETE "$API_ADDR/api/posts/$1")
  JQ=$(echo "$RES" | jq -r '.ok')

  if [[ "$JQ" -ne "deleted" ]]; then
    echo "delete. expected deleted, got:"
    echo "$RES" | jq
    exit
  fi
}
export -f test_delete

test_query_length() {
  # $1 expected length
  # $2 limit
  # $3 offset
  RESPONSE=$(curl -s -X GET "$API_ADDR/api/posts?limit=$2&offset=$3")

  LEN=$(echo "$RESPONSE" | jq -r '.posts | length')

  if [[ "$LEN" -ne "$1" ]]; then
    echo "query length for (lim=$2, off=$3). expected $1, got $LEN:"
    echo $RESPONSE
    exit
  fi
}
export -f test_query_length

test_modify() {
  # $1 postID
  # $2 modify json body

  RESPONSE=$(curl -s -X PATCH "$API_ADDR/api/posts/$1" \
    --header "Content-Type: application/json" \
    --data "$2")

  JQ=$(echo "$RESPONSE" | jq -r '.ok')

  if [[ "$JQ" -ne "changed" ]]; then
    echo "modify. expected changed, got:"
    echo "$RESPONSE"
    exit
  fi
}
export -f test_modify

test_tag_length() {
  # $1 postID
  # $2 expected length

  RESPONSE=$(curl -s -X GET "$API_ADDR/api/posts/$1")

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
export -f test_tag_length

test_tag_has() {
  # $1 postID
  # $2 contain json array
  RESPONSE=$(curl -s -X GET "$API_ADDR/api/posts/$1")

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
export -f test_tag_has
