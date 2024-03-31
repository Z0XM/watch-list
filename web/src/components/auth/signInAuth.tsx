"use client";

// import * as React from "react"

import { cn } from "@/lib/utils";
import { Icons } from "@/components/assets/icons";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { useState, useRef } from "react";
import { createClientComponentClient } from "@supabase/auth-helpers-nextjs";
import Link from "next/link";
import { toast } from "sonner";
import { useRouter } from "next/navigation";

interface UserAuthFormProps extends React.HTMLAttributes<HTMLDivElement> {}

export function UserAuthFormSignIn({ className, ...props }: UserAuthFormProps) {
    // const [isLoading, setIsLoading] = useState<boolean>(false);
    const [disableInput, setDisableInput] = useState<boolean>(false);
    const userEmailRef = useRef<HTMLInputElement | null>(null);
    const userPasswordRef = useRef<HTMLInputElement | null>(null);
    const router = useRouter();
    const supabase = createClientComponentClient();

    const onSubmit = async () => {
        setDisableInput(true);
        const email = userEmailRef.current?.value;
        const password = userPasswordRef.current?.value;

        if (!email) {
            toast.error("Please enter your email");
            setDisableInput(false);
            return;
        }
        if (!password) {
            toast.error("Please enter your password");
            setDisableInput(false);
            return;
        }

        const { data, error } = await supabase.auth.signInWithPassword({
            email,
            password,
        });

        if (error) {
            toast.error(error.message);
            setDisableInput(false);
            return;
        }
        router.refresh();
    };

    return (
        <div className={cn("grid gap-6", className)} {...props}>
            <form
                onSubmit={(e) => {
                    e.preventDefault();
                    onSubmit();
                }}
            >
                <div className="grid gap-2">
                    <div className="grid gap-1">
                        <Input
                            id="email"
                            placeholder="name@example.com"
                            type="email"
                            autoCapitalize="none"
                            autoComplete="email"
                            autoCorrect="off"
                            ref={userEmailRef}
                            disabled={disableInput}
                            required
                        />
                    </div>
                    <div className="grid gap-1">
                        <Input
                            id="email"
                            placeholder="Password"
                            type="password"
                            autoCapitalize="none"
                            autoComplete="email"
                            autoCorrect="off"
                            ref={userPasswordRef}
                            disabled={disableInput}
                            required
                        />
                    </div>  
                    <Button type="submit">Sign In with Email</Button>
                </div>
            </form>
            <div className="relative">
                <div className="absolute inset-0 flex items-center">
                    <span className="w-full border-t" />
                </div>
                <div className="relative flex justify-center text-xs uppercase">
                    <span className="bg-background px-2 text-muted-foreground">
                        Or continue with
                    </span>
                </div>
            </div>
            <div className="lg:p-8">
                <div className="mx-auto flex w-full flex-col justify-center space-y-6 sm:w-[350px]">
                    <div className="flex flex-col space-y-2 text-center">
                        <h1 className="text-xl font-semibold tracking-tight">
                            Don't have an account?
                        </h1>
                        <button className=" bg-slate-200 rounded p-2">
                            <Link href={"/signup"}>Go to Sign Up page</Link>
                        </button>
                    </div>
                </div>
            </div>
        </div>
    );
}
