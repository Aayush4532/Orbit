import { NextResponse } from "next/server";

const publicPaths = new Set(["/", "/login", "/signin", "/signup"]);
const protectedPrefixes = ["/home", "/events", "/orders"];

export function proxy(request) {
  const { pathname } = request.nextUrl;
  const token = request.cookies.get("token")?.value;

  const isPublicPath = publicPaths.has(pathname);
  const isProtectedPath = protectedPrefixes.some((prefix) => pathname === prefix || pathname.startsWith(`${prefix}/`));

  if (!token && isProtectedPath) {
    return NextResponse.redirect(new URL("/login", request.url));
  }

  if (token && (pathname === "/login" || pathname === "/signin" || pathname === "/signup")) {
    return NextResponse.redirect(new URL("/home", request.url));
  }

  return NextResponse.next();
}

export const config = {
  matcher: ["/", "/home", "/home/:path*", "/events", "/events/:path*", "/orders", "/orders/:path*", "/login", "/signin", "/signup"],
};
