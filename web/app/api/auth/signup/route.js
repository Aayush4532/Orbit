import { NextResponse } from "next/server";
const backendUrl = () => (process.env.BACKEND_API_URL || "http://localhost:9132").replace(/\/$/, "");
export async function POST(request) {
  const input = await request.json();
  try {
    const upstream = await fetch(`${backendUrl()}/api/auth/signup`, { method: "POST", headers: { "Content-Type": "application/json" }, body: JSON.stringify(input), cache: "no-store" });
    const body = await upstream.json();
    if (!upstream.ok) return NextResponse.json(body, { status: upstream.status });
    const token = upstream.headers.get("set-cookie")?.match(/token=([^;]+)/)?.[1];
    if (!token) return NextResponse.json({ error: "Backend did not return a session token." }, { status: 502 });
    const response = NextResponse.json(body, { status: 201 });
    response.cookies.set("token", token, { httpOnly: true, sameSite: "lax", secure: process.env.NODE_ENV === "production", path: "/", maxAge: 7200 });
    return response;
  } catch { return NextResponse.json({ error: "Orbit backend is unavailable. Start the Go server and try again." }, { status: 503 }); }
}
