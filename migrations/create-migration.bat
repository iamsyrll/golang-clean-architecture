@echo off
setlocal EnableDelayedExpansion

REM load .env
for /f "usebackq tokens=1,* delims==" %%a in (.env) do (
    set %%a=%%b
)

set CMD=%1
set NAME=%2

if "%CMD%"=="create" (
    if "%NAME%"=="" (
        echo Usage: migrate.bat create name
        exit /b 1
    )

    migrate create -ext sql -dir migrations -seq %NAME%
    exit /b 0
)

if "%DATABASE_URL%"=="" (
    echo DATABASE_URL not set
    exit /b 1
)

migrate -database "%DATABASE_URL%" -path migrations %*