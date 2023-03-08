export function logError(e: unknown, source: string): void {
	// eslint-disable-next-line no-console
	console.error(`[${new Date().toISOString().slice(0, 19)}Z]`, "[ERROR]", `[${source}]`, e);
}
