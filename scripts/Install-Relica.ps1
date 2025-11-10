# Relica Installer (User-local, no admin required)

$ErrorActionPreference = "Stop"

$installDir = "$env:USERPROFILE\AppData\Local\Programs\relica"
$exePath = "$installDir\relica.exe"
$releaseUrl = "https://github.com/opensourcerw/relica/releases/latest/download/relica_windows_amd64.zip"
$tempZip = "$env:TEMP\relica.zip"

Write-Host "Checking install location: $installDir"

# Create install directory if not exists
if (!(Test-Path $installDir)) {
    Write-Host "Creating install directory..."
    New-Item -ItemType Directory -Force -Path $installDir | Out-Null
}

Write-Host "Downloading latest Relica release..."
Invoke-WebRequest -Uri $releaseUrl -OutFile $tempZip

Write-Host "Extracting..."
Expand-Archive -Path $tempZip -DestinationPath $installDir -Force
Remove-Item $tempZip

# Add to user PATH if missing
$envPath = [System.Environment]::GetEnvironmentVariable("Path", "User")
if ($envPath -notlike "*$installDir*") {
    Write-Host "Adding Relica to user PATH..."
    [System.Environment]::SetEnvironmentVariable(
        "Path", "$envPath;$installDir", "User"
    )
    Write-Host "PATH updated — restart terminal to use Relica globally."
}

Write-Host "Relica installed successfully!"
Write-Host "Run 'relica --version' to verify installation."
