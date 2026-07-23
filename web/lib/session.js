import { cookies } from "next/headers";

const backendUrl = () => (process.env.BACKEND_API_URL || "http://localhost:9132").replace(/\/$/, "");

export async function getSession() {
  const token = (await cookies()).get("token")?.value;
  if (!token) return null;
  try {
    const response = await fetch(`${backendUrl()}/api/auth/check`, { headers: { cookie: `token=${token}` }, cache: "no-store" });
    if (!response.ok) return null;
    return (await response.json()).user || null;
  } catch { return null; }
}

export async function backendRequest(path, init = {}) {
  const token = (await cookies()).get("token")?.value;
  const headers = new Headers(init.headers);
  if (token) headers.set("cookie", `token=${token}`);
  return fetch(`${backendUrl()}${path}`, { ...init, headers, cache: "no-store" });
}