export interface WSMessage {
  type: 'join' | 'audio' | 'users';
  sender?: string;
  payload?: string; // Base64 audio
  count?: number;
}
