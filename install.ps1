# JeeraType One-Line Windows PowerShell Installer
$ErrorActionPreference = "Stop"

# Ensure TLS 1.2 is enabled for older PowerShell / Windows versions
[Net.ServicePointManager]::SecurityProtocol = [Net.ServicePointManager]::SecurityProtocol -bor [Net.SecurityProtocolType]::Tls12

$Repo = "Codexia-afk/JeeraType"
$BinaryName = "jeeratype.exe"

# Architecture Detection (amd64 or arm64)
$Arch = "amd64"
if ($env:PROCESSOR_ARCHITECTURE -eq "ARM64" -or $env:PROCESSOR_ARCHITEW6432 -eq "ARM64") {
    $Arch = "arm64"
}

Write-Host "🚀 Installing / Updating JeeraType for Windows ($Arch)..." -ForegroundColor Cyan

# 1. Query GitHub API for published releases (prioritizing prebuilt release assets over raw git tags)
$Release = $null
$Tag = $null

try {
    $Release = Invoke-RestMethod -Uri "https://api.github.com/repos/$Repo/releases/latest" -Headers @{"User-Agent"="PowerShell"} -ErrorAction Stop
    $Tag = $Release.tag_name
} catch {
    try {
        $Releases = Invoke-RestMethod -Uri "https://api.github.com/repos/$Repo/releases" -Headers @{"User-Agent"="PowerShell"} -ErrorAction Stop
        if ($Releases.Count -gt 0) {
            $Release = $Releases[0]
            $Tag = $Release.tag_name
        }
    } catch {
        try {
            $Tags = Invoke-RestMethod -Uri "https://api.github.com/repos/$Repo/tags" -Headers @{"User-Agent"="PowerShell"} -ErrorAction Stop
            if ($Tags.Count -gt 0) {
                $Tag = $Tags[0].name
            }
        } catch {}
    }
}

if ([string]::IsNullOrWhiteSpace($Tag)) {
    $Tag = "v2.5.0"
}

# 2. Determine Download URL
$DownloadUrl = $null

if ($Release -and $Release.assets) {
    $MatchingAsset = $Release.assets | Where-Object { $_.name -like "*windows*${Arch}*.zip" } | Select-Object -First 1
    if ($MatchingAsset) {
        $DownloadUrl = $MatchingAsset.browser_download_url
    }
}

if ([string]::IsNullOrWhiteSpace($DownloadUrl)) {
    $CleanTag = $Tag.TrimStart("v")
    $ZipName = "jeeratype_${CleanTag}_windows_${Arch}.zip"
    $DownloadUrl = "https://github.com/$Repo/releases/download/$Tag/$ZipName"
}

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
    Invoke-WebRequest -Uri $DownloadUrl -OutFile $TempZip -UseBasicParsing -ErrorAction Stop
    $DownloadSuccess = $true
} catch {
    # If initial URL failed, fallback to direct download URL of latest release
    try {
        $LatestRelease = Invoke-RestMethod -Uri "https://api.github.com/repos/$Repo/releases/latest" -Headers @{"User-Agent"="PowerShell"} -ErrorAction Stop
        $FallbackTag = $LatestRelease.tag_name
        $CleanFallbackTag = $FallbackTag.TrimStart("v")
        $FallbackUrl = "https://github.com/$Repo/releases/download/$FallbackTag/jeeratype_${CleanFallbackTag}_windows_${Arch}.zip"
        Write-Host "⚠️ Primary release zip not found. Falling back to latest release $FallbackTag..." -ForegroundColor Yellow
        Invoke-WebRequest -Uri $FallbackUrl -OutFile $TempZip -UseBasicParsing -ErrorAction Stop
        $DownloadSuccess = $true
    } catch {
        Write-Host "⚠️ Direct release zip not found. Attempting Go build fallback..." -ForegroundColor Yellow
    }
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
