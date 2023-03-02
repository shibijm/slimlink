import { createClient } from "redis";

const client = createClient({
	url: process.env.REDIS_CONNECTION_STRING,
	socket: {
		tls: true,
	},
});

export async function redisGet(key: string): Promise<string> {
	if (!client.isReady) {
		await client.connect();
	}
	const value = await client.get(key);
	return value || "";
}

export async function redisSet(key: string, value: string): Promise<void> {
	if (!client.isReady) {
		await client.connect();
	}
	await client.set(key, value);
}
