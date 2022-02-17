PWD=$(dirname "$0")
ROOT="$PWD/.."
cd "$ROOT"
go test ./tests/... -cover \
  -coverpkg=./controllers,./helpers \
  -coverprofile=coverage-report.out