import { apiPost } from "data/httpApi";
import { Link } from "types";

export async function addLink(url: string): Promise<Link> {
	return await apiPost<Link>("links", { url });
}
