import { apiRequest } from "services/client/data/http";
import { Link, LinkRequestDTO } from "types";

export async function addLink(linkRequestDTO: LinkRequestDTO): Promise<Link> {
	return await apiRequest<Link>("POST", "links", linkRequestDTO);
}
