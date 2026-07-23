import { redirect } from "next/navigation";
import EventsPageClient from "@/components/events-page-client";
import { getSession } from "@/lib/session";

export default async function EventsPage() {
  if (!(await getSession())) redirect("/login");

  return <EventsPageClient />;
}
