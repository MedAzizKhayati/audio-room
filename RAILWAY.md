# 🚂 Railway Deployment Guide

This guide will help you deploy the Audio Room application to Railway.

## 📋 Prerequisites

- A [Railway](https://railway.app/) account
- GitHub account (to connect your repository)
- This repository pushed to GitHub

## 🚀 Deploy to Railway

### Method 1: Using Railway CLI (Recommended)

1. **Install Railway CLI**
   ```bash
   npm i -g @railway/cli
   ```

2. **Login to Railway**
   ```bash
   railway login
   ```

3. **Initialize Project**
   ```bash
   railway init
   ```

4. **Deploy Backend**
   ```bash
   cd server
   railway up
   ```

5. **Deploy Frontend**
   ```bash
   cd ../frontend
   railway up
   ```

### Method 2: Using Railway Dashboard

1. **Go to [Railway Dashboard](https://railway.app/dashboard)**

2. **Create a New Project**
   - Click "New Project"
   - Select "Deploy from GitHub repo"
   - Choose your repository

3. **Deploy Backend Service**
   - Click "Add Service" → "GitHub Repo"
   - Select your repo
   - Set Root Directory: `server`
   - Railway will auto-detect the Dockerfile
   - Click "Deploy"

4. **Deploy Frontend Service**
   - Click "Add Service" → "GitHub Repo"
   - Select your repo again
   - Set Root Directory: `frontend`
   - Railway will auto-detect Next.js
   - Click "Deploy"

## ⚙️ Environment Variables

### Backend Service
No environment variables required! The app automatically:
- Uses Railway's `PORT` variable
- Enables CORS for cross-origin requests

### Frontend Service
Add this environment variable:
- `NEXT_PUBLIC_BACKEND_URL` = Your backend Railway URL (e.g., `https://your-backend.railway.app`)

**To set it:**
1. Go to Frontend service in Railway
2. Click "Variables" tab
3. Add `NEXT_PUBLIC_BACKEND_URL`
4. Paste your backend URL (without trailing slash)
5. Click "Redeploy"

## 🌐 Get Your URLs

After deployment:
1. Go to backend service → Settings → Generate Domain
2. Copy the URL (e.g., `https://audio-room-backend.railway.app`)
3. Add it to frontend's `NEXT_PUBLIC_BACKEND_URL`
4. Go to frontend service → Settings → Generate Domain
5. Your app is live! 🎉

## 🔧 Configuration Files

The following files are configured for Railway:

- `server/Dockerfile` - Backend container configuration
- `frontend/Dockerfile` - Frontend container configuration  
- `railway.backend.json` - Backend Railway config
- `railway.frontend.json` - Frontend Railway config
- `frontend/src/config/api.ts` - API URL configuration with Railway support

## 📱 Testing Your Deployment

1. Open your frontend Railway URL
2. Enter a username
3. Grant microphone permissions
4. Start talking!

**Note:** Make sure you're on HTTPS (Railway provides this automatically) for microphone access to work.

## 🐛 Troubleshooting

### Backend won't start
- Check Railway logs: `railway logs`
- Verify Dockerfile builds locally: `docker build -t test ./server`

### Frontend can't connect to backend
- Verify `NEXT_PUBLIC_BACKEND_URL` is set correctly
- Check backend is running and accessible
- Ensure backend URL doesn't have trailing slash

### WebSocket connection fails
- Railway supports WebSockets by default
- Check CORS is enabled in backend (already configured)
- Verify you're using `wss://` protocol (handled automatically)

### Microphone doesn't work
- Ensure you're on HTTPS (Railway provides this)
- Check browser permissions
- Try different browser

## 💰 Cost

Railway offers:
- **Free tier**: $5 of usage per month
- **Pro plan**: $20/month for more resources

This app is lightweight and should work fine on the free tier for testing!

## 🔄 Auto-Deploy

Railway automatically redeploys when you push to your GitHub repo's main branch. To disable:
1. Go to service settings
2. Uncheck "Auto-deploy"

## 📊 Monitoring

View logs in real-time:
```bash
railway logs --service backend
railway logs --service frontend
```

Or check in the Railway dashboard under the "Logs" tab.

---

Need help? Check [Railway Docs](https://docs.railway.app/) or open an issue!
