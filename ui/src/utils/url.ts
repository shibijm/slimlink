export function isValidUrl(url: string): boolean {
	let urlObj;
	try {
		urlObj = new URL(url);
	} catch {
		return false;
	}
	return urlObj.protocol === "http:" || urlObj.protocol === "https:";
}
