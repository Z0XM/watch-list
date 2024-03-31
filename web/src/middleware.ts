import { createMiddlewareClient } from "@supabase/auth-helpers-nextjs";
import { NextResponse } from "next/server";
import type { NextRequest } from "next/server";

export async function middleware(req: NextRequest) {
    console.log("Inside Middleware");
    const res = NextResponse.next();
    const supabase = createMiddlewareClient({ req, res });
    const sessionResponse = await supabase.auth.getSession();
    const currentSession = sessionResponse.data.session;

    if (currentSession) {
        if (req.nextUrl.pathname.startsWith("/signup") || req.nextUrl.pathname.startsWith("/login") ) {
            console.log("inside current Session");
            return NextResponse.rewrite(new URL("/", req.url)); 
        }
        return res;
    }

    return res;
}

export const config = {
    matcher: [
        /*
         * Match all request paths except for the ones starting with:
         * - _next/static (static files)
         * - _next/image (image optimization files)
         * - favicon.ico (favicon file)
         */
        "/((?!_next/static|_next/image|favicon.ico).*)",
    ],
};
