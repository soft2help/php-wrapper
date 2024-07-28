cd /d %~dp0
call stops2hWs.bat > NUL
LauncherBackground.exe s2hWrapper.exe -r "index.php"
