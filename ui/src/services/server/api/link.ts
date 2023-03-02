import { redisGet, redisSet } from "services/server/data/redis";
import { Link, LinkRequestDTO } from "types";
import { generateBase62String } from "utils/base62";
import { isValidUrl } from "utils/url";

const linkIDLength = (process.env.LINK_ID_LENGTH && parseInt(process.env.LINK_ID_LENGTH)) || 5;

export async function addLink(linkRequestDTO: LinkRequestDTO): Promise<Link> {
	if (!isValidUrl(linkRequestDTO.url)) {
		throw new Error("Invalid URL");
	}
	let id;
	let url;
	let tries = 0;
	do {
		id = generateBase62String(linkIDLength);
		// eslint-disable-next-line no-await-in-loop
		url = await redisGet(id);
	} while (url && ++tries < 10);
	if (url) {
		throw new Error("Maximum tries exceeded for link ID generation");
	}
	await redisSet(id, linkRequestDTO.url);
	return { id, url: linkRequestDTO.url };
}

export async function getLinkByID(id: string): Promise<Link> {
	const url = await redisGet(id);
	if (!url) {
		throw new Error("Link not found");
	}
	return { id, url };
}
