#!/bin/bash

echo "🚂 Railway Deployment Script for Audio Room"
echo "============================================"
echo ""

# Check if Railway CLI is installed
if ! command -v railway &> /dev/null; then
    echo "❌ Railway CLI not found. Installing..."
    npm i -g @railway/cli
fi

# Login to Railway
echo "🔐 Logging into Railway..."
railway login

# Deploy Backend
echo ""
echo "📦 Deploying Backend Service..."
cd server

# Initialize and deploy backend
railway init
echo "✅ Backend deployed!"

# Get backend URL
echo "🌐 Generating backend domain..."
railway domain
echo ""
echo "⚠️  IMPORTANT: Copy the backend URL above!"
read -p "Enter your backend URL (e.g., https://server-production-xxxx.up.railway.app): " BACKEND_URL

cd ..

# Deploy Frontend
echo ""
echo "📦 Deploying Frontend Service..."
cd frontend

# Link to the same project
echo "🔗 Linking to existing project..."
railway link

# Set environment variable
echo "⚙️  Setting environment variable..."
railway variables --set NEXT_PUBLIC_BACKEND_URL=$BACKEND_URL

# Deploy frontend
railway up

# Generate frontend domain
echo "🌐 Generating frontend domain..."
railway domain

echo ""
echo "✅ Deployment Complete!"
echo "======================="
echo ""
echo "🎉 Your Audio Room app is now live!"
echo ""
echo "Backend: $BACKEND_URL"
echo "Frontend: Check the URL above"
echo ""
echo "View logs:"
echo "  Backend:  cd server && railway logs"
echo "  Frontend: cd frontend && railway logs"
