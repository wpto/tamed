source "./curl.sh"


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
