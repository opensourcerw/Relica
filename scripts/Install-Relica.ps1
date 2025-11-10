# Relica Installer (user-local, resilient to ConstrainedLanguage)

$ErrorActionPreference = "Stop"
$installDir = "$env:LOCALAPPDATA\Programs\relica"
$zipUrl    = "https://github.com/opensourcerw/relica/releases/latest/download/relica_windows_amd64.zip"
$tempZip   = Join-Path $env:TEMP "relica.zip"

Write-Host "LanguageMode: $($ExecutionContext.SessionState.LanguageMode)"

if (!(Test-Path $installDir)) {
    New-Item -ItemType Directory -Force -Path $installDir | Out-Null
}

Invoke-WebRequest -Uri $zipUrl -OutFile $tempZip
Expand-Archive -Path $tempZip -DestinationPath $installDir -Force
Remove-Item $tempZip -Force

$needsPath = ($env:Path -notlike "*$installDir*")

if ($needsPath) {
    if ($ExecutionContext.SessionState.LanguageMode -eq 'ConstrainedLanguage') {
        Write-Host "ConstrainedLanguage: skipping automatic PATH modification."
        Write-Host "Add manually or run:"
        Write-Host "cmd /c setx PATH \"$env:Path;$installDir\""
    } else {
        $userPath = [Environment]::GetEnvironmentVariable('Path','User')
        if ($userPath -notlike "*$installDir*") {
            [Environment]::SetEnvironmentVariable('Path', "$userPath;$installDir", 'User')
            Write-Host "Added to user PATH. Restart terminal."
        }
    }
} else {
    Write-Host "Already in PATH."
}

Write-Host "Installed."
Write-Host "Run: & '$installDir\\relica.exe' --version"