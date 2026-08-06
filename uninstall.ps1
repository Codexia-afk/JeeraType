# JeeraType Universal Uninstaller Script for Windows
$ErrorActionPreference = "Stop"

$BinaryName = "jeeratype.exe"
$InstallDir = Join-Path $env:USERPROFILE ".jeeratype\bin"
$BaseDir = Join-Path $env:USERPROFILE ".jeeratype"
$ConfigDir = Join-Path $env:AppData "jeeratype"
$UserConfigDir = Join-Path $env:USERPROFILE ".config\jeeratype"

Write-Host "🗑️ Uninstalling JeeraType from Windows..." -ForegroundColor Yellow

# Remove from User PATH in Windows Registry
$UserPath = [Environment]::GetEnvironmentVariable("Path", "User")
if ($UserPath -like "*$InstallDir*") {
    $NewPath = ($UserPath -split ";" | Where-Object { $_ -ne $InstallDir -and $_ -ne "" }) -join ";"
    [Environment]::SetEnvironmentVariable("Path", $NewPath, "User")
    Write-Host "  - Removed JeeraType from User PATH environment variable." -ForegroundColor Cyan
}

# Remove installation folder
if (Test-Path $BaseDir) {
    Remove-Item $BaseDir -Recurse -Force -ErrorAction SilentlyContinue
    Write-Host "  - Removed installation directory ($BaseDir)." -ForegroundColor Cyan
}

# Remove app data / stats folder
if (Test-Path $ConfigDir) {
    Remove-Item $ConfigDir -Recurse -Force -ErrorAction SilentlyContinue
    Write-Host "  - Removed AppData configuration directory ($ConfigDir)." -ForegroundColor Cyan
}
if (Test-Path $UserConfigDir) {
    Remove-Item $UserConfigDir -Recurse -Force -ErrorAction SilentlyContinue
    Write-Host "  - Removed user configuration directory ($UserConfigDir)." -ForegroundColor Cyan
}

# Remove from active memory session
$env:Path = ($env:Path -split ";" | Where-Object { $_ -ne $InstallDir -and $_ -ne "" }) -join ";"
if (Test-Path Function:\global:jeeratype) {
    Remove-Item Function:\global:jeeratype -ErrorAction SilentlyContinue
}

Write-Host ""
Write-Host "✅ JeeraType has been completely uninstalled from Windows!" -ForegroundColor Green
