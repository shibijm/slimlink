import { apiPost } from "data/httpApi";
import { Link } from "types";

export async function createLink(url: string): Promise<Link> {
	return await apiPost<Link>("links", { url });
}
