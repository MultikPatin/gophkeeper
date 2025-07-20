go build -ldflags="\
-X 'main.buildVersion=v1.0' \
-X 'main.buildDate=`date '+%Y-%m-%dT%H:%M:%SZ'`' \
-X 'main.buildCommit=`git rev-parse HEAD`'" \
-o build/client/gothkeeper ./cmd/client