import { createClient } from "redis";
import { NotFoundError } from "server/exceptions";

const client = createClient({ url: process.env.REDIS_CONNECTION_STRING });

export async function getRedisKey(key: string): Promise<string> {
	if (!client.isReady) {
		await client.connect();
	}
	const value = await client.get(key);
	if (!value) {
		throw new NotFoundError("Link not found");
	}
	return value;
}

export async function setRedisKey(key: string, value: string): Promise<void> {
	if (!client.isReady) {
		await client.connect();
	}
	await client.set(key, value);
}
