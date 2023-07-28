import { getRedisKey, setRedisKey } from "server/data/redis";
import { BadRequestError, NotFoundError, UnexpectedError } from "server/exceptions";
import { generateBase62String } from "server/utils";
import { Link } from "types";
import { isValidUrl } from "utils";

const linkIDLength = (process.env.LINK_ID_LENGTH && parseInt(process.env.LINK_ID_LENGTH)) || 5;

export async function createLink(url: string): Promise<Link> {
	if (!isValidUrl(url)) {
		throw new BadRequestError("Invalid URL");
	}
	let id;
	for (let attempt = 1; ; attempt++) {
		id = generateBase62String(linkIDLength);
		try {
			// eslint-disable-next-line no-await-in-loop
			await getRedisKey(id);
		} catch (e) {
			if (e instanceof NotFoundError) {
				break;
			}
			throw e;
		}
		if (attempt === 10) {
			throw new UnexpectedError("Maximum tries exceeded for link ID generation");
		}
	}
	await setRedisKey(id, url);
	return { id, url: url };
}

export async function getLinkByID(id: string): Promise<Link> {
	const url = await getRedisKey(id);
	return { id, url };
}
