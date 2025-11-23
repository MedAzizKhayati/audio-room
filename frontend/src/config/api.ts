// API Configuration
const getApiConfig = () => {
  if (typeof window === 'undefined') {
    return {
      wsUrl: '',
      httpUrl: ''
    };
  }

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
