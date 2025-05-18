@echo off
setlocal enabledelayedexpansion

for /f %%i in ('git rev-parse HEAD') do set GIT_COMMIT=%%i
for /f %%i in ('git rev-parse --abbrev-ref HEAD') do set GIT_BRANCH=%%i
for /f %%i in ('powershell -NoProfile -Command "git describe --tags --abbrev=0"') do set GIT_VERSION=%%i

echo GIT_COMMIT=!GIT_COMMIT!
echo GIT_BRANCH=!GIT_BRANCH!
echo GIT_VERSION=!GIT_VERSION!

for /f %%i in ('powershell -NoProfile -Command "Get-Date -Format o"') do set BUILD_TIME=%%i

go build -trimpath -o adiff.exe ^
 -ldflags="-X main.GitCommit=!GIT_COMMIT! -X main.BuildTime=!BUILD_TIME! -X main.GitBranch=!GIT_BRANCH! -X main.GitVersion=!GIT_VERSION! -extldflags \"-static\" -s -w" ^
 .

endlocal
REM pause
