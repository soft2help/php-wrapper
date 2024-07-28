# Get the current directory
$currentDir = Get-Location

# Create the Release directory
New-Item -ItemType Directory -Force -Path "$currentDir\Release"

# Copy files and directories to the Release directory
Copy-Item -Path "public_html" -Destination "$currentDir\Release\public_html" -Recurse -Force
Copy-Item -Path "LauncherBackground.exe" -Destination "$currentDir\Release\" -Force
Copy-Item -Path "s2hWrapper.exe" -Destination "$currentDir\Release\" -Force
Copy-Item -Path "launcher.bat" -Destination "$currentDir\Release\" -Force
Copy-Item -Path "stops2hWs.bat" -Destination "$currentDir\Release\" -Force

# Define the path for the ZIP file in the current directory
$zipPath = "$currentDir\Release.zip"

# Remove the ZIP file if it already exists
if (Test-Path $zipPath) {
    Remove-Item $zipPath -Force
}

# Create the ZIP file
Add-Type -AssemblyName "System.IO.Compression.FileSystem"
[System.IO.Compression.ZipFile]::CreateFromDirectory("$currentDir\Release", $zipPath)

# Output success message
Write-Output "Release folder successfully zipped to $zipPath"
