import { LoginForm } from "@/components/login-form"

export default function Page() {
  return (
    <main className="relative flex min-h-svh w-full items-center justify-center overflow-hidden bg-[#090909] p-6 md:p-10">
      <div className="relative w-full max-w-md">
        <LoginForm />
      </div>
    </main>
  )
}
