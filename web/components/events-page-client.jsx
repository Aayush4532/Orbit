"use client";

import Link from "next/link";
import { useEffect, useMemo, useState } from "react";
import { ArrowUpRight, Clock3, Search, SlidersHorizontal } from "lucide-react";
import { DashboardNav } from "@/components/orbit-nav";

function Product({ color }) {
  return (
    <div className="relative h-44 overflow-hidden rounded-2xl bg-[#222]">
      <div
        className={`absolute left-[12%] top-12 h-24 w-[80%] -rotate-6 rounded-[50%_43%_18%_23%/47%_47%_25%_30%] bg-gradient-to-br ${color} shadow-xl`}
      />
      <div className="absolute left-[29%] top-[58%] h-1 w-[34%] -rotate-6 rounded bg-white/70" />
      <div className="absolute bottom-8 left-[19%] h-2 w-[76%] -rotate-6 rounded-full bg-black/70" />
    </div>
  );
}

export default function EventsPageClient() {
  const [events, setEvents] = useState([]);
  const [tab, setTab] = useState("All");
  const [query, setQuery] = useState("");
  const [error, setError] = useState("");

  useEffect(() => {
    fetch("/api/events")
      .then(async (response) => {
        if (!response.ok) throw new Error("Could not load events");
        return response.json();
      })
      .then(({ events: loaded }) => setEvents(loaded))
      .catch(() => setError("Events are temporarily unavailable."));
  }, []);

  const shown = useMemo(() => {
    const normalizedQuery = query.toLowerCase();

    return events.filter((event) => {
      const matchesTab =
        tab === "All" ||
        event.status === (tab === "Live now" ? "live" : "upcoming");
      const matchesQuery = `${event.name} ${event.edition}`
        .toLowerCase()
        .includes(normalizedQuery);

      return matchesTab && matchesQuery;
    });
  }, [events, query, tab]);

  return (
    <main className="min-h-screen bg-[#0a0a0a] text-white">
      <DashboardNav active="Events" />
      <div className="mx-auto max-w-7xl px-5 py-10 sm:px-8 lg:px-10">
        <p className="text-xs font-semibold uppercase tracking-[0.18em] text-lime-300">Discover</p>
        <div className="mt-3 flex flex-wrap items-end justify-between gap-5">
          <div>
            <h1 className="text-3xl font-semibold tracking-[-0.045em] sm:text-4xl">Events, made fair.</h1>
            <p className="mt-2 text-sm text-zinc-500">
              Every product belongs to an event. Explore what is live now.
            </p>
          </div>
          <Link
            href="/signup"
            className="inline-flex items-center gap-2 rounded-full border border-white/15 px-4 py-2.5 text-sm font-medium hover:border-lime-300 hover:text-lime-300"
          >
            List an event <ArrowUpRight size={16} />
          </Link>
        </div>

        <div className="mt-10 flex flex-col justify-between gap-4 border-y border-white/10 py-4 sm:flex-row sm:items-center">
          <div className="flex gap-1 rounded-xl bg-white/[.04] p-1">
            {[
              "All",
              "Live now",
              "Upcoming",
            ].map((item) => (
              <button
                key={item}
                onClick={() => setTab(item)}
                className={`rounded-lg px-3 py-2 text-sm transition ${
                  tab === item ? "bg-white text-black" : "text-zinc-500 hover:text-white"
                }`}
              >
                {item}
              </button>
            ))}
          </div>

          <label className="flex w-full items-center gap-2 rounded-xl border border-white/10 bg-white/[.03] px-3 py-2 sm:w-64">
            <Search size={16} className="text-zinc-500" />
            <input
              value={query}
              onChange={(event) => setQuery(event.target.value)}
              placeholder="Search events"
              className="w-full bg-transparent text-sm outline-none placeholder:text-zinc-600"
            />
          </label>

          <button aria-label="Filters" className="hidden rounded-xl border border-white/10 p-2.5 text-zinc-400 sm:block">
            <SlidersHorizontal size={17} />
          </button>
        </div>

        <p className="mt-7 text-sm text-zinc-500">
          {error || (
            <>
              <span className="text-white">{shown.length}</span> events available
            </>
          )}
        </p>

        <section className="mt-5 grid gap-5 sm:grid-cols-2 lg:grid-cols-3">
          {shown.map((event) => (
            <article
              key={event.slug}
              className="group rounded-3xl border border-white/10 bg-[#121212] p-4 transition hover:-translate-y-1 hover:border-white/20"
            >
              <Product color={event.color} />
              <div className="mt-4 flex items-start justify-between gap-3">
                <div>
                  <p className="text-[11px] uppercase tracking-[0.22em] text-zinc-500">{event.status}</p>
                  <h2 className="mt-1 text-lg font-semibold text-white">{event.name}</h2>
                  <p className="mt-1 text-sm text-zinc-400">{event.edition}</p>
                </div>
                <span className="rounded-full border border-white/10 px-2.5 py-1 text-[11px] text-zinc-400">
                  {event.products?.length || 0} items
                </span>
              </div>
              <div className="mt-5 flex items-center justify-between gap-3 rounded-2xl border border-white/10 bg-white/[.03] p-3">
                <div className="flex items-center gap-2 text-sm text-zinc-300">
                  <Clock3 size={15} className="text-lime-300" />
                  <span className="font-mono text-white">{event.timeLeft || "00:18:24"}</span>
                </div>
                <Link
                  href={`/events/${event.slug}`}
                  className="rounded-full bg-white px-3 py-1.5 text-xs font-bold text-black transition group-hover:bg-lime-300"
                >
                  View
                </Link>
              </div>
            </article>
          ))}
        </section>
      </div>
    </main>
  );
}
