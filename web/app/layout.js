import { Inter, Poppins } from "next/font/google";
import "./globals.css";

const poppins = Poppins({
  subsets: ["latin"],
  weight: ["400", "500", "600", "700"],
  variable: "--font-poppins",
});

const inter = Inter({
  subsets: ["latin"],
  weight: ["400", "500", "600", "700"],
  variable: "--font-inter",
});

export const metadata = {
  title: "Orbit — Fair access to the drops you want",
  description: "A modern event commerce platform for high-demand products.",
};

export default function RootLayout({ children }) {
  return (
    <html lang="en" className={`h-full antialiased ${poppins.variable} ${inter.variable}`}>
      <body className="min-h-full bg-[#06070a] text-white">{children}</body>
    </html>
  );
}
