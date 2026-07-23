import { NextResponse } from "next/server";

export async function POST(request) {
  const isSecureOrigin = request.url.startsWith("https://");
  const response = NextResponse.json({ ok: true });
  response.cookies.set("token", "", { httpOnly: true, sameSite: "lax", secure: isSecureOrigin, path: "/", maxAge: 0 });
  return response;
}
