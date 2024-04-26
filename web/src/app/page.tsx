"use client";
import { createClientComponentClient } from "@supabase/auth-helpers-nextjs";
import { useRouter } from "next/navigation";
import Link from "next/link";

export default function homepage() {
    const supabase = createClientComponentClient();
    const router = useRouter();
    const handleSignOut = async () => {
        await supabase.auth.signOut();
        router.refresh();
    };

    return (
        <>
            <div className="w-screen h-screen bg-slate-400 flex flex-col gap-2 justify-center items-center">
                <p className="text-stone-300">Hello from Home Page</p>
                <div className="flex flex-col gap-2 justify-center items-center">
                    <button className=" bg-zinc-200 rounded p-2">
                        <Link href={"/signup"}>Sign Up</Link>
                    </button>
                    <button className=" bg-zinc-200 rounded p-2">
                        <Link href={"/login"}>Login</Link>
                    </button>
                </div>
                <button
                    onClick={handleSignOut}
                    className=" bg-slate-200 rounded p-2"
                >
                    {" "}
                    Sign Out
                </button>
            </div>
        </>
    );
}
