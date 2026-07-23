import { NextResponse } from "next/server";
import { backendRequest } from "@/lib/session";
async function proxy(request, { params }) {
  const path = `/${(await params).path.join("/")}`;
  const body = ["GET", "HEAD"].includes(request.method) ? undefined : await request.text();
  try {
    const upstream = await backendRequest(`/api${path}`, { method: request.method, body, headers: body ? { "Content-Type": request.headers.get("content-type") || "application/json" } : {} });
    return new NextResponse(await upstream.text(), { status: upstream.status, headers: { "content-type": upstream.headers.get("content-type") || "application/json" } });
  } catch { return NextResponse.json({ error: "Orbit backend is unavailable." }, { status: 503 }); }
}
export const GET = proxy; export const POST = proxy; export const PUT = proxy; export const DELETE = proxy;
