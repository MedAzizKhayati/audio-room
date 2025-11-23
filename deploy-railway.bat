@echo off
echo 🚂 Railway Deployment Script for Audio Room
echo ============================================
echo.

REM Check if Railway CLI is installed
where railway >nul 2>nul
if %errorlevel% neq 0 (
    echo ❌ Railway CLI not found. Installing...
    npm i -g @railway/cli
)

REM Login to Railway
echo 🔐 Logging into Railway...
railway login

REM Deploy Backend
echo.
echo 📦 Deploying Backend Service...
cd server

REM Initialize and deploy backend
railway init
echo ✅ Backend deployed!

REM Get backend URL
echo 🌐 Generating backend domain...
railway domain
echo.
echo ⚠️  IMPORTANT: Copy the backend URL above!
set /p BACKEND_URL="Enter your backend URL (e.g., https://server-production-xxxx.up.railway.app): "

cd ..

REM Deploy Frontend
echo.
echo 📦 Deploying Frontend Service...
cd frontend

REM Link to the same project
echo 🔗 Linking to existing project...
railway link

REM Set environment variable
echo ⚙️  Setting environment variable...
railway variables --set NEXT_PUBLIC_BACKEND_URL=%BACKEND_URL%

REM Deploy frontend
railway up

REM Generate frontend domain
echo 🌐 Generating frontend domain...
railway domain

echo.
echo ✅ Deployment Complete!
echo =======================
echo.
echo 🎉 Your Audio Room app is now live!
echo.
echo Backend: %BACKEND_URL%
echo Frontend: Check the URL above
echo.
echo View logs:
echo   Backend:  cd server ^&^& railway logs
echo   Frontend: cd frontend ^&^& railway logs
