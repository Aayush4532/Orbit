import { NextResponse } from "next/server";
const backendUrl = () => (process.env.BACKEND_API_URL || "http://localhost:9132").replace(/\/$/, "");
export async function POST(request) {
  const input = await request.json();
  const isSecureOrigin = request.url.startsWith("https://");

  try {
    const upstream = await fetch(`${backendUrl()}/api/auth/signin`, { method: "POST", headers: { "Content-Type": "application/json" }, body: JSON.stringify({ emailId: input.emailId, password: input.password }), cache: "no-store" });
    const body = await upstream.json();
    if (!upstream.ok) return NextResponse.json(body, { status: upstream.status });
    const token = upstream.headers.get("set-cookie")?.match(/token=([^;]+)/)?.[1];
    if (!token) return NextResponse.json({ error: "Backend did not return a session token." }, { status: 502 });
    const response = NextResponse.json(body);
    response.cookies.set("token", token, { httpOnly: true, sameSite: "lax", secure: isSecureOrigin, path: "/", maxAge: 7200 });
    return response;
  } catch { return NextResponse.json({ error: "Orbit backend is unavailable. Start the Go server and try again." }, { status: 503 }); }
}
