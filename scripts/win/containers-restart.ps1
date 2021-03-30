$ROOT = $MyInvocation.MyCommand.Path.replace("scripts\win\containers-restart.ps1", "")
$CWD = Get-Location

Set-Location "$ROOT"
docker-compose restart
Set-Location $CWD  
