$ROOT = $MyInvocation.MyCommand.Path.replace("scripts\win\containers-down.ps1", "")
$CWD = Get-Location

Set-Location "$ROOT"
docker-compose down 
Set-Location $CWD  
