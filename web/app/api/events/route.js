import { NextResponse } from "next/server";
import { listEvents } from "@/lib/backend";
export async function GET() { try { return NextResponse.json({ events: await listEvents() }); } catch (error) { return NextResponse.json({ error: error.message || "Events unavailable" }, { status: 502 }); } }
