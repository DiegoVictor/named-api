PWD=$(dirname "$0")
ROOT="$PWD/.."
cd "$ROOT"
./scripts/test.sh
go tool cover -html=coverage-report.out