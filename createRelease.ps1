mkdir -p Release\public_html
Copy-Item -Path .s2hWS -Destination Release\.s2hWS -recurse -Force
Copy-Item -Path LauncherBackground.exe -Destination Release\ -recurse -Force
Copy-Item -Path s2hWrapper.exe -Destination Release\ -recurse -Force
Copy-Item -Path launcher.bat -Destination Release\ -recurse -Force
Copy-Item -Path stops2hWs.bat -Destination Release\ -recurse -Force
