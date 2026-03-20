@echo off
setlocal EnableDelayedExpansion

REM load .env
for /f "usebackq tokens=1,* delims==" %%a in (.env) do (
    set %%a=%%b
)

if "%DATABASE_URL%"=="" (
    echo DATABASE_URL not set
    exit /b 1
)

echo Running migration...

migrate -database "%DATABASE_URL%" -path migrations %*

if %errorlevel% neq 0 (
    echo Migration failed!
    exit /b %errorlevel%
)

echo Done!