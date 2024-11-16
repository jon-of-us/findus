findus.exe
@echo off
set /p TARGET_DIR=<path.txt
del path.txt
cd /d %TARGET_DIR%
powershell -NoLogo
