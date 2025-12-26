import { apiGet } from "@/data/httpApi";
import type { Info } from "@/types";

export async function getInfo(): Promise<Info> {
	return await apiGet<Info>("info");
}
