import { eventSummary, events, getEvent } from "@/lib/catalog";
import { backendRequest } from "@/lib/session";

const baseUrl = process.env.BACKEND_API_URL?.replace(/\/$/, "");
async function upstream(path, init) {
  const response = await backendRequest(path, { ...init, headers: { "Content-Type": "application/json", ...(init?.headers || {}) } });
  if (!response.ok) throw new Error(`Backend request failed (${response.status})`);
  return response.json();
}
export async function listEvents() { if (!baseUrl) return events.map(eventSummary); const payload = await upstream("/api/buyer/events"); return payload.events || payload; }
export async function findEvent(slug) { if (!baseUrl) return getEvent(slug); const payload = await upstream(`/api/buyer/event/${slug}`); return payload.event || payload; }
export async function createOrder(input) {
  if (baseUrl) return upstream(`/api/buyer/event/${input.eventId}/purchase/${input.productId}`, { method: "POST" });
  return { id: `ord_demo_${Date.now()}`, status: "confirmed", ...input, createdAt: new Date().toISOString() };
}
