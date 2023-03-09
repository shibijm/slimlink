@echo off
if not "%1"=="go" (
	cd web
	call npm i
	call npx next build
	call npx next export
	cd ..
)
if "%1"=="nextjs" (
	exit /b 0
)
setlocal
set CGO_ENABLED=0
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
