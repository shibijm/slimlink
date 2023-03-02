@echo off
if not "%1"=="1" (
	cd ui
	call npm i
	call npx next build
	call npx next export
	cd ..
	echo.
)
setlocal
set GOARCH=amd64
set GOOS=windows
echo Building Windows binary...
go build -ldflags "-s -w" -trimpath -o out/
call :checkBuildStatus
set GOOS=linux
echo Building Linux binary...
go build -ldflags "-s -w" -trimpath -o out/
call :checkBuildStatus
endlocal
exit /b 0
:checkBuildStatus
if not %errorlevel%==0 (
	echo Build failed
) else (
	echo Build succeeded
)
