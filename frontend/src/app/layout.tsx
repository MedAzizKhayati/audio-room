import type { Metadata } from "next";
import { Geist, Geist_Mono } from "next/font/google";
import "./globals.css";

const geistSans = Geist({
  variable: "--font-geist-sans",
  subsets: ["latin"],
});

const geistMono = Geist_Mono({
  variable: "--font-geist-mono",
  subsets: ["latin"],
});

export const metadata: Metadata = {
  title: "Audio Room - Real-time Voice Chat",
  description: "Connect and communicate with others in real-time using voice chat. Built with Next.js and Go for seamless audio streaming.",
  keywords: ["audio chat", "voice chat", "real-time communication", "webrtc", "websocket"],
  authors: [{ name: "Mohamed Aziz Khayati" }],
  openGraph: {
    title: "Audio Room - Real-time Voice Chat",
    description: "Connect and communicate with others in real-time using voice chat.",
    type: "website",
  },
  icons: {
    icon: "/favicon.svg",
  },
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en">
      <body
        className={`${geistSans.variable} ${geistMono.variable} antialiased`}
      >
        {children}
      </body>
    </html>
  );
}
