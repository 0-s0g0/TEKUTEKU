import type { Metadata } from "next";
import { Geist, Geist_Mono,Zen_Kurenaido } from "next/font/google";
import "./globals.css";

//全体で利用したい場合layout.tsxで指定


const ZenKurenaido = Zen_Kurenaido({
  weight: "400",
  variable: "--font-zen-maru-gothic",
  subsets: ["latin"],
});


export const metadata: Metadata = {
  title: "Create Next App",
  description: "Generated by create next app",
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en">
      <body
        className={ZenKurenaido.className}
      >
        {children}
      </body>
    </html>
  );
}
