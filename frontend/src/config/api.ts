// API Configuration
const getApiConfig = () => {
  if (typeof window === 'undefined') {
    return {
      wsUrl: '',
      httpUrl: ''
    };
  }

  // Check if we're in production (Railway deployment)
  const isProduction = process.env.NODE_ENV === 'production';
  const backendUrl = process.env.NEXT_PUBLIC_BACKEND_URL;

  if (isProduction && backendUrl) {
    // Railway production environment
    const wsProtocol = backendUrl.startsWith('https') ? 'wss:' : 'ws:';
    const httpProtocol = backendUrl.startsWith('https') ? 'https:' : 'http:';
    const url = backendUrl.replace(/^https?:\/\//, '');
    
    return {
      wsUrl: `${wsProtocol}//${url}/api/v1/audio/room`,
      httpUrl: `${httpProtocol}//${url}/api/v1`
    };
  }

  // Development environment
  const protocol = window.location.protocol === 'https:' ? 'wss:' : 'ws:';
  const httpProtocol = window.location.protocol;
  const host = window.location.hostname;
  const port = '8080';

  return {
    wsUrl: `${protocol}//${host}:${port}/api/v1/audio/room`,
    httpUrl: `${httpProtocol}//${host}:${port}/api/v1`
  };
};

export const API_CONFIG = getApiConfig();
