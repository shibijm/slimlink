import { apiGet } from "data/httpApi";
import { Info } from "types";

export async function getInfo(): Promise<Info> {
	return await apiGet<Info>("info");
}
