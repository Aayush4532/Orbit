import Link from "next/link";
import { Bell, Menu, Search } from "lucide-react";

export function AppMark({ dark = false }) {
  return <Link href="/" aria-label="Orbit home" className={`inline-flex ${dark ? "text-white" : "text-zinc-900"}`}><span className="grid h-8 w-8 place-items-center rounded-full bg-lime-300 text-sm font-black text-black shadow-lg shadow-blue-500/20">O</span></Link>;
}

export function MarketingNav() {
  return (
    <header className="relative z-20 mx-auto flex max-w-7xl items-center justify-between px-5 py-5 sm:px-8 lg:px-10">
      <AppMark dark />
      <div className="flex items-center gap-2">
        <Link href="/login" className="hidden px-4 py-2 text-sm font-medium text-zinc-300 transition hover:text-white sm:block">
          Log in
        </Link>
        <Link href="/signup" className="rounded-full bg-white px-4 py-2.5 text-sm font-semibold text-black transition hover:bg-lime-300">
          Get started
        </Link>
        <Menu size={20} className="ml-2 text-zinc-300 md:hidden" />
      </div>
    </header>
  );
}

export function DashboardNav({ active = "Home" }) {
  const links = [["Home", "/home"], ["Events", "/events"], ["Orders", "/orders"], ["Saved", "/events"]];
  return <header className="sticky top-0 z-30 border-b border-white/10 bg-[#0a0a0a]/85 backdrop-blur-xl"><div className="mx-auto flex h-16 max-w-7xl items-center justify-between px-5 sm:px-8 lg:px-10"><AppMark dark /><nav className="hidden items-center gap-1 md:flex">{links.map(([name, href]) => <Link key={name} href={href} className={`rounded-lg px-3 py-2 text-sm transition ${active === name ? "bg-white/10 text-white" : "text-zinc-500 hover:text-zinc-200"}`}>{name}</Link>)}</nav><div className="flex items-center gap-3"><button aria-label="Search" className="hidden rounded-full p-2 text-zinc-400 hover:bg-white/10 sm:block"><Search size={19} /></button><button aria-label="Notifications" className="relative rounded-full p-2 text-zinc-400 hover:bg-white/10"><Bell size={19} /><span className="absolute right-2 top-1.5 h-1.5 w-1.5 rounded-full bg-lime-300" /></button><div className="grid h-8 w-8 place-items-center rounded-full bg-gradient-to-br from-lime-300 to-emerald-500 text-xs font-bold text-black">AK</div></div></div></header>;
}
