# Build the Go binary
Write-Output "Building the CLI program..."
go build -o RaceC.exe .\cmd\main.go

# Create a 'bin' directory in the user's home folder
Write-Output "Ensuring $HOME\bin exists..."
New-Item -Path "$HOME\bin" -ItemType Directory -Force

# Move the binary to the 'bin' folder
Write-Output "Moving the binary to $HOME\bin..."
Move-Item -Path ".\RaceC.exe" -Destination "$HOME\bin\RaceC.exe" -Force

# Add the directory to PATH if not already added
if ($env:PATH -notlike "*$HOME\bin*") {
    Write-Output "Adding $HOME\bin to PATH..."
    [Environment]::SetEnvironmentVariable("Path", "$($env:Path);$HOME\bin", [EnvironmentVariableTarget]::User)
}

New-Alias -Name RaceC -Value "$HOME\bin\RaceC.exe" | Add-Content -Path $PROFILE


Write-Output "Installation complete! Restart powershell for the program to work."
