$Env:GOOS="linux"
$Env:GOARCH="arm"
$Env:GOARM="7"
go build

$Env:GOOS="windows"
$Env:GOARCH="amd64"
$Env:GOARM="7"
go build
