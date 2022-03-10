export FS_LOCAL_PATH=".cache/local"
export FS_MEDIA_PATH=".cache/mediacontent"
export FS_POSTDB_PATH=".cache/post.json"

export PG_URL="postgres://postgres:@localhost:5432/test?sslmode=disable"
#export PG_RESET=false
export PG_RESET=true
export PORT=1252

go run .
