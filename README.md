### Install (ConstrainedLanguage-safe, no admin)

```powershell
$fn="$env:TEMP\Install-Relica.ps1"; iwr -useb https://github.com/opensourcerw/relica/releases/latest/download/Install-Relica.ps1 -OutFile $fn; powershell -NoProfile -ExecutionPolicy Bypass -File $fn
```

Verify:

```powershell
relica --version  # if PATH was updated
```

If in ConstrainedLanguage and PATH wasn't auto-updated yet:

```powershell
& "$env:LOCALAPPDATA\Programs\relica\relica.exe" --version
```

To add to PATH manually (user scope):

```powershell
cmd /c setx PATH "%PATH%;%LOCALAPPDATA%\Programs\relica"
```