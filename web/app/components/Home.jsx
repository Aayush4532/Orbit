import Link from "next/link";
import {
  ArrowRight,
  ArrowUpRight,
  BadgeCheck,
  Clock3,
  Flame,
  ShieldCheck,
  Sparkles,
  TrendingUp,
} from "lucide-react";
import { AppMark, MarketingNav } from "@/components/orbit-nav";

const drops = [
  {
    name: "Jordan 1 Retro High OG",
    detail: "Chicago Reimagined · US 8–12",
    price: "$180",
    left: "6 pairs left",
    time: "00:18:24",
    image: "https://images.unsplash.com/photo-1542291026-7eec264c27ff?auto=format&fit=crop&w=900&q=80",
    accent: "from-red-500 via-rose-500 to-orange-300",
  },
  {
    name: "Jordan 4 Retro SE",
    detail: "Craft Medium Olive · US 7–13",
    price: "$210",
    left: "24 pairs left",
    time: "02:18:04",
    image: "https://images.unsplash.com/photo-1600185365483-26d7a4cc7519?auto=format&fit=crop&w=900&q=80",
    accent: "from-emerald-400 via-teal-400 to-lime-300",
  },
  {
    name: "Nike Air Max Pulse",
    detail: "Midnight Volt · US 6–11",
    price: "$180",
    left: "14 pairs left",
    time: "04:05:31",
    image: "https://images.unsplash.com/photo-1605348532760-6753d2c43329?auto=format&fit=crop&w=900&q=80",
    accent: "from-slate-100 via-zinc-400 to-slate-700",
  },
];

const pillars = [
  {
    icon: ShieldCheck,
    title: "Verified inventory",
    copy: "Every item is reviewed before the event opens so buyers move with confidence.",
  },
  {
    icon: Flame,
    title: "Built for demand",
    copy: "Timed checkouts create a calmer, faster, and more transparent release flow.",
  },
  {
    icon: BadgeCheck,
    title: "Instant confirmation",
    copy: "Order status and confirmation appear in seconds, so the drop stays frictionless.",
  },
];

const steps = [
  {
    num: "01",
    title: "Create a buyer pool",
    copy: "Join the event and get seated with a fair queue based on real-time inventory.",
  },
  {
    num: "02",
    title: "Watch the countdown",
    copy: "See the timer, exact stock visibility, and keep each release transparent.",
  },
  {
    num: "03",
    title: "Check out cleanly",
    copy: "Secure the item with one streamlined flow and get confirmation instantly.",
  },
];

function DropTile({ drop }) {
  return (
    <article className="group overflow-hidden rounded-[28px] border border-white/10 bg-white/[0.04] transition duration-300 hover:-translate-y-1 hover:border-white/20 hover:bg-white/[0.06]">
      <div className="relative overflow-hidden border-b border-white/10 bg-[#121212]">
        <div className={`absolute inset-0 bg-gradient-to-br ${drop.accent} opacity-35 blur-3xl`} />
        <img
          src={drop.image}
          alt={drop.name}
          className="relative h-64 w-full object-cover object-center transition duration-500 group-hover:scale-[1.02]"
        />
      </div>
      <div className="p-4 sm:p-5">
        <div className="flex items-center justify-between gap-3 text-[11px] uppercase tracking-[0.22em] text-zinc-400">
          <span>{drop.left}</span>
          <span className="font-mono text-zinc-500">{drop.time}</span>
        </div>
        <h3 className="mt-3 text-lg font-semibold text-white">{drop.name}</h3>
        <p className="mt-1 text-sm text-zinc-400">{drop.detail}</p>
        <div className="mt-5 flex items-center justify-between gap-3">
          <span className="text-lg font-semibold text-lime-300">{drop.price}</span>
          <Link
            href="/events"
            className="rounded-full bg-white px-3.5 py-2 text-xs font-bold text-black transition group-hover:bg-lime-300"
          >
            Enter
          </Link>
        </div>
      </div>
    </article>
  );
}

export default function Home() {
  return (
    <main className="relative isolate min-h-screen overflow-hidden bg-[#05070b] text-white selection:bg-lime-300 selection:text-black">
      <div className="pointer-events-none absolute inset-0 bg-[radial-gradient(circle_at_top_left,rgba(163,230,53,0.14),transparent_20%),radial-gradient(circle_at_top_right,rgba(56,189,248,0.12),transparent_28%),radial-gradient(circle_at_bottom_left,rgba(168,85,247,0.14),transparent_24%),linear-gradient(180deg,#05070b_0%,#090b0f_100%)]" />

      <MarketingNav />

      <section className="relative mx-auto grid max-w-7xl gap-6 px-5 pb-10 pt-3 sm:px-8 lg:grid-cols-[1.05fr_.95fr] lg:items-center lg:px-10 lg:pb-12 lg:pt-4">
        <div className="pointer-events-none absolute left-[-4rem] top-8 h-72 w-72 rounded-full bg-lime-300/15 blur-[120px]" />
        <div className="pointer-events-none absolute bottom-8 right-5 h-72 w-72 rounded-full bg-sky-500/10 blur-[120px]" />

        <div className="relative z-10 max-w-2xl">
          <div className="mb-6 inline-flex items-center gap-2 rounded-full border border-lime-300/30 bg-lime-300/10 px-3 py-1 text-[11px] font-semibold uppercase tracking-[0.24em] text-lime-300">
            <Sparkles size={14} /> Commerce, in the moment
          </div>

          <h1 className="text-5xl font-semibold leading-[0.93] tracking-[-0.08em] text-white sm:text-6xl lg:text-[5.8rem]">
            Your next drop,
            <br />
            without the noise.
          </h1>

          <p className="mt-4 max-w-xl text-base leading-7 text-zinc-400 sm:text-lg">
            Orbit gives serious buyers a clean, fair, and high-trust way to enter live releases with real-time inventory, smart timing, and one fast checkout.
          </p>

          <div className="mt-6 flex flex-wrap gap-3">
            <Link
              href="/events"
              className="inline-flex items-center gap-2 rounded-full bg-lime-300 px-6 py-3.5 text-sm font-bold text-black transition hover:bg-lime-200"
            >
              Explore live drops <ArrowUpRight size={17} />
            </Link>
            <Link
              href="/signup"
              className="rounded-full border border-white/15 bg-white/[0.05] px-6 py-3.5 text-sm font-semibold text-white transition hover:border-white/35 hover:bg-white/[0.08]"
            >
              Become a seller
            </Link>
          </div>

          <div className="mt-6 grid max-w-lg grid-cols-3 gap-4 border-t border-white/10 pt-4">
            {[
              ["99.8%", "successful checkouts"],
              ["4.9/5", "buyer rating"],
              ["24/7", "verified support"],
            ].map(([stat, label]) => (
              <div key={stat}>
                <p className="text-xl font-semibold tracking-tight text-white">{stat}</p>
                <p className="mt-1 text-xs leading-4 text-zinc-500">{label}</p>
              </div>
            ))}
          </div>
        </div>

        <div className="relative mx-auto w-full max-w-xl">
          <div className="absolute -right-4 top-8 h-56 w-56 rounded-full bg-fuchsia-500/20 blur-[110px]" />
          <div className="relative rounded-[2rem] border border-white/10 bg-white/[0.04] p-3 shadow-[0_24px_80px_rgba(0,0,0,0.55)] backdrop-blur-xl">
            <div className="rounded-[1.5rem] bg-[#101217] p-3 sm:p-4">
              <div className="flex items-center justify-between gap-4">
                <span className="rounded-full bg-lime-300/15 px-3 py-1 text-[11px] font-bold uppercase tracking-[0.24em] text-lime-300">
                  Live event
                </span>
                <span className="flex items-center gap-2 text-xs text-zinc-400">
                  <span className="h-2 w-2 animate-pulse rounded-full bg-red-400" />
                  342 watching
                </span>
              </div>

              <div className="mt-3 overflow-hidden rounded-[1.5rem] border border-white/10">
                <img
                  src="https://images.unsplash.com/photo-1542291026-7eec264c27ff?auto=format&fit=crop&w=1200&q=80"
                  alt="Featured sneaker drop"
                  className="h-56 w-full object-cover object-center"
                />
              </div>

              <div className="mt-5 flex items-end justify-between gap-4">
                <div>
                  <p className="text-[11px] uppercase tracking-[0.24em] text-zinc-500">Nike · Jordan 1</p>
                  <h2 className="mt-1 text-2xl font-semibold tracking-tight text-white">Chicago Reimagined</h2>
                </div>
                <p className="text-2xl font-semibold text-lime-300">$180</p>
              </div>

              <div className="mt-5 flex items-center justify-between rounded-2xl border border-white/10 bg-black/30 p-3.5">
                <div className="flex items-center gap-2 text-sm text-zinc-200">
                  <Clock3 size={16} className="text-lime-300" />
                  Ends in
                  <span className="ml-1 font-mono font-semibold tracking-[0.22em] text-white">00:18:24</span>
                </div>
                <span className="text-xs text-zinc-500">6 / 10 left</span>
              </div>
            </div>
          </div>

          <div className="absolute -bottom-7 -left-6 hidden rounded-2xl border border-white/10 bg-[#16181d] p-4 shadow-xl sm:block">
            <div className="flex items-center gap-3">
              <div className="flex h-9 w-9 items-center justify-center rounded-full bg-lime-300 text-black">
                <BadgeCheck size={17} strokeWidth={3} />
              </div>
              <div>
                <p className="text-xs font-semibold text-white">Fair access, always</p>
                <p className="text-[11px] text-zinc-500">No queues. No games.</p>
              </div>
            </div>
          </div>
        </div>
      </section>

      <section className="border-y border-white/10 bg-white/[0.03]">
        <div className="mx-auto max-w-7xl px-5 py-5 sm:px-8 lg:px-10">
          <p className="text-center text-[11px] font-semibold uppercase tracking-[0.32em] text-zinc-500">
            Trusted by collectors who move fast
          </p>
        </div>
      </section>

      <section className="mx-auto max-w-7xl px-5 py-10 sm:px-8 lg:px-10">
        <div className="mb-5 flex items-end justify-between gap-5">
          <div>
            <p className="text-xs font-semibold uppercase tracking-[0.24em] text-lime-300">Happening now</p>
            <h2 className="mt-3 text-3xl font-semibold tracking-[-0.05em] text-white sm:text-4xl">
              Catch the next live drop.
            </h2>
          </div>
          <Link href="/events" className="hidden items-center gap-1 text-sm text-zinc-300 hover:text-lime-300 sm:flex">
            View all events <ArrowUpRight size={16} />
          </Link>
        </div>

        <div className="grid gap-5 md:grid-cols-3">
          {drops.map((drop) => (
            <DropTile key={drop.name} drop={drop} />
          ))}
        </div>
      </section>

      <section className="mx-auto max-w-7xl px-5 pb-10 sm:px-8 lg:px-10">
        <div className="grid gap-5 lg:grid-cols-3">
          {pillars.map(({ icon: Icon, title, copy }) => (
            <div key={title} className="rounded-[26px] border border-white/10 bg-white/[0.04] p-6">
              <Icon className="mb-8 text-lime-300" size={22} />
              <h3 className="text-lg font-semibold text-white">{title}</h3>
              <p className="mt-2 text-sm leading-6 text-zinc-500">{copy}</p>
            </div>
          ))}
        </div>
      </section>

      <section className="mx-auto max-w-7xl px-5 pb-6 sm:px-8 lg:px-10">
        <div className="rounded-[30px] border border-white/10 bg-white/[0.03] p-5 sm:p-6">
          <p className="text-xs font-semibold uppercase tracking-[0.24em] text-lime-300">How it works</p>
          <div className="mt-6 grid gap-4 md:grid-cols-3">
            {steps.map((step) => (
              <div key={step.num} className="rounded-[24px] border border-white/10 bg-[#101114] p-5">
                <div className="text-xs font-semibold uppercase tracking-[0.24em] text-zinc-500">{step.num}</div>
                <h3 className="mt-3 text-lg font-semibold text-white">{step.title}</h3>
                <p className="mt-2 text-sm leading-6 text-zinc-500">{step.copy}</p>
              </div>
            ))}
          </div>
        </div>
      </section>

      <section className="mx-auto max-w-7xl px-5 pb-8 sm:px-8 lg:px-10">
        <div className="rounded-[30px] border border-white/10 bg-gradient-to-br from-white/[0.06] to-transparent p-5 sm:p-6">
          <div className="flex flex-col gap-5 lg:flex-row lg:items-center lg:justify-between">
            <div>
              <p className="text-xs font-semibold uppercase tracking-[0.24em] text-lime-300">Why Orbit</p>
              <h2 className="mt-3 text-2xl font-semibold tracking-tight text-white sm:text-3xl">
                Better drops, cleaner access, sharper momentum.
              </h2>
            </div>
            <div className="flex items-center gap-2 rounded-full border border-white/10 bg-black/20 px-4 py-2 text-sm text-zinc-300">
              <TrendingUp size={16} className="text-lime-300" />
              Real-time event visibility
            </div>
          </div>
        </div>
      </section>

      <section className="mx-auto max-w-7xl px-5 pb-10 sm:px-8 lg:px-10">
        <div className="flex flex-col gap-4 rounded-[30px] border border-white/10 bg-[#0b0c11] px-5 py-6 sm:px-6 lg:flex-row lg:items-center lg:justify-between">
          <div>
            <p className="text-xs font-semibold uppercase tracking-[0.24em] text-lime-300">Launch with confidence</p>
            <h3 className="mt-2 text-2xl font-semibold tracking-tight text-white">Turn demand into a premium event experience.</h3>
          </div>
          <Link
            href="/signup"
            className="inline-flex items-center gap-2 self-start rounded-full bg-white px-5 py-3 text-sm font-bold text-black transition hover:bg-lime-300"
          >
            Start building your drop <ArrowRight size={16} />
          </Link>
        </div>
      </section>

      <footer className="border-t border-white/10 py-6">
        <div className="mx-auto flex max-w-7xl flex-wrap items-center justify-between gap-4 px-5 text-xs text-zinc-500 sm:px-8 lg:px-10">
          <AppMark />
          <p>© 2026 Orbit. Better drops, on repeat.</p>
          <div className="flex gap-5">
            <span>Privacy</span>
            <span>Terms</span>
            <span>Support</span>
          </div>
        </div>
      </footer>
    </main>
  );
}
