import { NextResponse } from "next/server";
import { createOrder, findEvent } from "@/lib/backend";
import { getSession } from "@/lib/session";
export async function POST(request) { const user = await getSession(); if (!user) return NextResponse.json({ error: "Unauthorized" }, { status: 401 }); const input = await request.json(); const event = await findEvent(input.eventSlug); const product = event?.products?.find((item) => item.id === input.productId); if (!product) return NextResponse.json({ error: "Product unavailable" }, { status: 404 }); return NextResponse.json({ order: await createOrder({ eventSlug: event.slug, productId: product.id, productName: product.name, price: product.price, buyer: user.email }) }, { status: 201 }); }
