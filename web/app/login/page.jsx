import { LoginForm } from "@/components/login-form"

export default function Page() {
  return (
    <main className="relative flex min-h-svh w-full items-center justify-center overflow-hidden bg-[#090909] p-6 md:p-10">
      <div className="pointer-events-none absolute -left-32 top-1/4 h-96 w-96 rounded-full bg-lime-400/10 blur-[120px]" />
      <div className="relative w-full max-w-md">
        <LoginForm />
      </div>
    </main>
  );
}
