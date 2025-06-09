$env:GOOS="linux";
$env:GOARCH="amd64";
$ScriptDir = Split-Path -Parent $MyInvocation.MyCommand.Definition
$MainPath = Join-Path $ScriptDir "cmd\relay\main.go"
go build -o (Join-Path $ScriptDir "bin/relay-linux") $MainPath