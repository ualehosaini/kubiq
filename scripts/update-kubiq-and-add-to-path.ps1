# PowerShell Script: Update kubiq.exe and Add Directory to User PATH
# Usage: Run this script in PowerShell as administrator if you want to add for all users, or as your user for just yourself.

# Ensure latest kubiq.exe from build is in this directory
$kubiqDir = (Get-Location).Path
$buildKubiq = Join-Path $kubiqDir 'build' 'kubiq.exe'
$targetKubiq = Join-Path $kubiqDir 'kubiq.exe'
if (Test-Path $buildKubiq) {
    Copy-Item $buildKubiq $targetKubiq -Force
    Write-Host "Copied latest build/kubiq.exe to $kubiqDir."
} else {
    Write-Host "No build/kubiq.exe found. Skipping copy."
}

# Get current user PATH
$oldPath = [Environment]::GetEnvironmentVariable("Path", "User")

if ($oldPath -notlike "*$kubiqDir*") {
    $newPath = "$oldPath;$kubiqDir"
    [Environment]::SetEnvironmentVariable("Path", $newPath, "User")
    Write-Host "Added $kubiqDir to your user PATH. Restart your terminal to use 'kubiq' globally."
} else {
    Write-Host "$kubiqDir is already in your user PATH."
}

Write-Host "Current user PATH is now:"
Write-Host ([Environment]::GetEnvironmentVariable("Path", "User"))
