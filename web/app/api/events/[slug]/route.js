import { NextResponse } from "next/server";
import { findEvent } from "@/lib/backend";
export async function GET(_request, { params }) { try { const event = await findEvent((await params).slug); return event ? NextResponse.json({ event }) : NextResponse.json({ error: "Event not found" }, { status: 404 }); } catch (error) { return NextResponse.json({ error: error.message || "Event unavailable" }, { status: 502 }); } }
