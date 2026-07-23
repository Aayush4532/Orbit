export const events = [
  { slug: "jordan-1-chicago-reimagined", brand: "NIKE", name: "Jordan 1 Retro High OG", edition: "Chicago Reimagined", price: 180, stock: 6, totalStock: 10, endsAt: "2026-07-18T18:00:00.000Z", status: "live", color: "from-red-600 via-rose-500 to-orange-200", products: [{ id: "j1-8", name: "Jordan 1 Retro High OG", size: "US 8", price: 180 }, { id: "j1-9", name: "Jordan 1 Retro High OG", size: "US 9", price: 180 }, { id: "j1-10", name: "Jordan 1 Retro High OG", size: "US 10", price: 180 }] },
  { slug: "jordan-4-craft-olive", brand: "NIKE", name: "Jordan 4 Retro SE", edition: "Craft Medium Olive", price: 210, stock: 24, totalStock: 30, endsAt: "2026-07-18T21:00:00.000Z", status: "live", color: "from-lime-400 via-emerald-500 to-teal-200", products: [{ id: "j4-8", name: "Jordan 4 Retro SE", size: "US 8", price: 210 }, { id: "j4-9", name: "Jordan 4 Retro SE", size: "US 9", price: 210 }, { id: "j4-10", name: "Jordan 4 Retro SE", size: "US 10", price: 210 }] },
  { slug: "jordan-3-white-cement", brand: "NIKE", name: "Jordan 3 Retro", edition: "White Cement Reimagined", price: 200, stock: 16, totalStock: 20, endsAt: "2026-07-19T10:00:00.000Z", status: "live", color: "from-zinc-100 via-zinc-400 to-zinc-700", products: [{ id: "j3-9", name: "Jordan 3 Retro", size: "US 9", price: 200 }, { id: "j3-10", name: "Jordan 3 Retro", size: "US 10", price: 200 }] },
  { slug: "nb-990v6-grey-day", brand: "NEW BALANCE", name: "990v6 Made in USA", edition: "Grey Day", price: 220, stock: 30, totalStock: 30, endsAt: "2026-07-20T11:00:00.000Z", status: "upcoming", color: "from-slate-100 via-slate-400 to-slate-700", products: [{ id: "nb-9", name: "990v6 Made in USA", size: "US 9", price: 220 }] },
];

export function eventSummary(event) {
  const { products, ...summary } = event;
  return summary;
}

export function getEvent(slug) { return events.find((event) => event.slug === slug); }
