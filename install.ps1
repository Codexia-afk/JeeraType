# JeeraType One-Line Windows PowerShell Installer
$ErrorActionPreference = "Stop"

$Repo = "Codexia-afk/JeeraType"
$BinaryName = "jeeratype.exe"

Write-Host "🚀 Installing / Updating JeeraType for Windows..." -ForegroundColor Cyan

# Query latest version tag from GitHub API
try {
    $Tags = Invoke-RestMethod -Uri "https://api.github.com/repos/$Repo/tags"
    $Tag = $Tags[0].name
} catch {
    try {
        $Release = Invoke-RestMethod -Uri "https://api.github.com/repos/$Repo/releases/latest"
        $Tag = $Release.tag_name
    } catch {
        $Tag = "v2.4.0"
    }
}

$CleanTag = $Tag.TrimStart("v")
$ZipName = "jeeratype_${CleanTag}_windows_amd64.zip"
$DownloadUrl = "https://github.com/$Repo/releases/download/$Tag/$ZipName"

# Target Install Directory: %USERPROFILE%\.jeeratype\bin
$InstallDir = Join-Path $env:USERPROFILE ".jeeratype\bin"
if (!(Test-Path $InstallDir)) {
    New-Item -ItemType Directory -Path $InstallDir -Force | Out-Null
}

$TempZip = Join-Path $env:TEMP "jeeratype.zip"
$TempExtract = Join-Path $env:TEMP "jeeratype_extract"

Write-Host "📥 Downloading JeeraType $Tag..." -ForegroundColor Cyan

$DownloadSuccess = $false
try {
    Invoke-WebRequest -Uri $DownloadUrl -OutFile $TempZip -UseBasicParsing
    $DownloadSuccess = $true
} catch {
    Write-Host "⚠️ Direct release zip not found. Attempting Go build fallback..." -ForegroundColor Yellow
}

if ($DownloadSuccess -and (Test-Path $TempZip)) {
    Write-Host "📦 Extracting executable..." -ForegroundColor Cyan
    if (Test-Path $TempExtract) { Remove-Item $TempExtract -Recurse -Force }
    Expand-Archive -Path $TempZip -DestinationPath $TempExtract -Force

    $ExtractedExe = Get-ChildItem -Path $TempExtract -Filter "jeeratype.exe" -Recurse | Select-Object -First 1
    if ($ExtractedExe) {
        Copy-Item -Path $ExtractedExe.FullName -Destination (Join-Path $InstallDir $BinaryName) -Force
    }
} else {
    # Fallback for Windows if go compiler is installed
    if (Get-Command go -ErrorAction SilentlyContinue) {
        Write-Host "⚙️ Building latest JeeraType binary directly from main branch..." -ForegroundColor Cyan
        go install "github.com/${Repo}@main"
        $GoExe = Join-Path (go env GOPATH) "bin\jeeratype.exe"
        if (Test-Path $GoExe) {
            Copy-Item -Path $GoExe -Destination (Join-Path $InstallDir $BinaryName) -Force
        }
    }
}

$FinalExe = Join-Path $InstallDir $BinaryName
if (!(Test-Path $FinalExe)) {
    Write-Error "Could not locate or compile jeeratype.exe."
    exit 1
}

# Add to User PATH in Windows Registry
$UserPath = [Environment]::GetEnvironmentVariable("Path", "User")
if ($UserPath -notlike "*$InstallDir*") {
    Write-Host "⚙️ Adding JeeraType to User PATH..." -ForegroundColor Yellow
    [Environment]::SetEnvironmentVariable("Path", "$UserPath;$InstallDir", "User")
}

# Update current session PATH in memory
$env:Path = "$InstallDir;$env:Path"

# Create a temporary function in current session so 'jeeratype' works immediately in this window
function global:jeeratype { & "$InstallDir\jeeratype.exe" $args }

# Cleanup
Remove-Item $TempZip -Force -ErrorAction SilentlyContinue
Remove-Item $TempExtract -Recurse -Force -ErrorAction SilentlyContinue

Write-Host ""
Write-Host "✅ JeeraType updated successfully!" -ForegroundColor Green
Write-Host "Type 'jeeratype' in your terminal to start!" -ForegroundColor White
