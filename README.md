# 🎙️ Audio Room

A real-time audio communication application built with Go, Next.js, and WebSockets. Multiple users can join a voice room and communicate with each other in real-time with audio visualization.

![Audio Room](https://img.shields.io/badge/Next.js-15-black?logo=next.js)
![Go](https://img.shields.io/badge/Go-1.21-00ADD8?logo=go)
![TypeScript](https://img.shields.io/badge/TypeScript-5-3178C6?logo=typescript)
![WebSocket](https://img.shields.io/badge/WebSocket-Real--time-orange)

## ✨ Features

- 🎤 **Real-time Voice Chat** - Communicate with multiple users simultaneously
- 🎨 **Audio Visualization** - Real-time frequency visualization with Canvas API
- 👥 **Active Speaker Indicators** - See who's talking in real-time
- 🔐 **Username Persistence** - Your username is saved in localStorage
- 📱 **Mobile Responsive** - Works on desktop and mobile devices
- 🎯 **Modern Stack** - Built with the latest web technologies
- 🐳 **Docker Ready** - Easy deployment with Docker Compose

## 🏗️ Architecture

### Frontend
- **Framework**: Next.js 15 (App Router)
- **Language**: TypeScript
- **Styling**: Tailwind CSS
- **Audio Processing**: Web Audio API with AudioWorklet
- **State Management**: React Hooks + Context API
- **Icons**: Lucide React

### Backend
- **Language**: Go 1.21
- **Framework**: Gin
- **WebSocket**: Gorilla WebSocket
- **Architecture**: Clean Architecture with dependency injection

### Communication
- **Protocol**: WebSocket for real-time bidirectional communication
- **Audio Format**: Mono-channel 16-bit Linear PCM (Base64 encoded)
- **Buffer Size**: 8192 samples for efficient streaming

## 🚀 Quick Start

### Prerequisites

- Node.js 20+ and npm/yarn
- Go 1.21+
- Docker (optional)

### Development Mode

#### Backend
```bash
cd server
go mod download
go run cmd/api/main.go
```

The backend will start on `http://localhost:8080`

#### Frontend
```bash
cd frontend
npm install
npm run dev
```

The frontend will start on `http://localhost:3000`

### Production with Docker

```bash
# Build and run both services
docker-compose up --build

# Run in detached mode
docker-compose up -d

# Stop services
docker-compose down
```

## 📁 Project Structure

```
go-project/
├── frontend/                 # Next.js frontend application
│   ├── public/
│   │   └── audio-processor.js    # AudioWorklet processor
│   ├── src/
│   │   ├── app/             # Next.js app router pages
│   │   ├── components/      # React components
│   │   │   ├── AudioRoom.tsx
│   │   │   ├── JoinScreen.tsx
│   │   │   └── HomeContent.tsx
│   │   ├── hooks/           # Custom React hooks
│   │   │   ├── useAudioPlayer.ts
│   │   │   ├── useAudioRecorder.ts
│   │   │   ├── useAudioWebSocket.ts
│   │   │   └── useVisualizer.ts
│   │   ├── contexts/        # React Context providers
│   │   │   └── UsernameContext.tsx
│   │   ├── config/          # Configuration files
│   │   │   └── api.ts
│   │   └── types.ts         # TypeScript type definitions
│   ├── Dockerfile
│   └── package.json
│
├── server/                  # Go backend application
│   ├── cmd/
│   │   └── api/
│   │       └── main.go      # Application entry point
│   ├── internal/
│   │   ├── handler/         # HTTP/WebSocket handlers
│   │   │   ├── audio.go
│   │   │   ├── health.go
│   │   │   └── module.go
│   │   ├── middleware/      # HTTP middleware
│   │   │   ├── timing.go
│   │   │   └── module.go
│   │   ├── service/         # Business logic
│   │   │   └── module.go
│   │   ├── domain/          # Domain models
│   │   └── config.go
│   ├── Dockerfile
│   └── go.mod
│
├── docker-compose.yml       # Docker orchestration
└── README.md               # This file
```

## 🎯 How It Works

### Audio Recording Flow
1. User grants microphone permission
2. AudioWorklet captures audio in 8192-sample buffers
3. Float32 audio is converted to Int16 PCM format
4. PCM data is Base64 encoded
5. Data is sent via WebSocket to the server
6. Server broadcasts to all connected clients

### Audio Playback Flow
1. Client receives Base64 audio data from WebSocket
2. Base64 is decoded to Int16 PCM
3. PCM is converted back to Float32
4. AudioContext creates and plays the audio buffer
5. Active speaker indicator is shown for 1 second

### WebSocket Message Types
- `join` - User joins the room
- `audio` - Audio data payload
- `users` - User count update

## 🔧 Configuration

### Frontend Configuration
Edit `frontend/src/config/api.ts`:
```typescript
export const API_CONFIG = {
  wsUrl: 'ws://localhost:8080/api/v1/audio/room',
  httpUrl: 'http://localhost:8080/api/v1'
};
```

### Backend Configuration
The server automatically detects the correct ports and protocols based on your environment.

## 🎨 Key Features Explained

### AudioWorklet (Modern Audio Processing)
Replaces the deprecated `ScriptProcessorNode` with a modern, efficient audio processing pipeline that runs on a separate thread for better performance.

### Username Context with LocalStorage
Your username is stored globally using React Context and persists across page refreshes using localStorage.

### Custom React Hooks
- **useAudioWebSocket**: Manages WebSocket connection and reconnection
- **useAudioRecorder**: Handles microphone access and audio streaming
- **useAudioPlayer**: Manages audio playback and active speaker tracking
- **useVisualizer**: Renders real-time audio frequency visualization

### Mobile Support
Uses `dvh` (dynamic viewport height) units for proper mobile layout, especially important for iOS Safari.

## 🐛 Troubleshooting

### Microphone not working
- Ensure you're on HTTPS or localhost (browsers block mic on HTTP)
- Check browser permissions (lock icon in address bar)
- Try incognito mode to reset permissions

### No audio visualization
- Make sure microphone access is granted
- Check that the mic button is enabled (red icon)

### WebSocket connection fails
- Verify backend is running on port 8080
- Check firewall settings
- Ensure ports are not blocked

## 📝 API Endpoints

### WebSocket
- `ws://localhost:8080/api/v1/audio/room` - Main audio room WebSocket

### HTTP
- `GET /api/v1/health` - Health check endpoint

## 🚢 Deployment

### Railway (Recommended)

See [RAILWAY.md](./RAILWAY.md) for detailed Railway deployment instructions.

**Quick Deploy:**
1. Push to GitHub
2. Connect repo to Railway
3. Deploy backend and frontend as separate services
4. Set `NEXT_PUBLIC_BACKEND_URL` environment variable in frontend
5. Done! 🎉

### Docker Compose

### Environment Variables (Optional)
```bash
# Backend
PORT=8080

# Frontend
NODE_ENV=production
```

### Building for Production

#### Frontend
```bash
cd frontend
npm run build
npm start
```

#### Backend
```bash
cd server
go build -o audio-server ./cmd/api
./audio-server
```

## 🤝 Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## 📄 License

This project is open source and available under the MIT License.

## 🙏 Acknowledgments

- Web Audio API for modern audio processing
- Gorilla WebSocket for robust WebSocket implementation
- Next.js team for an amazing React framework
- Tailwind CSS for utility-first styling

---

Built with ❤️ using Go and Next.js