$ROOT = $MyInvocation.MyCommand.Path.replace("scripts\win\containers-up.ps1", "")
$CWD = Get-Location

Set-Location "$ROOT"
docker-compose up -d
Set-Location $CWD
