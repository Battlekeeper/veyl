$ScriptDir = Split-Path -Parent $MyInvocation.MyCommand.Definition
$MainPath = Join-Path $ScriptDir "cmd\client\main.go"
go build -o (Join-Path $ScriptDir "bin/client-windows.exe") $MainPath