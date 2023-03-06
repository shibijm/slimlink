export function isValidUrl(url: string): boolean {
	if (url.length > 2048) {
		return false;
	}
	let urlObj;
	try {
		urlObj = new URL(url);
	} catch {
		return false;
	}
	return urlObj.protocol === "http:" || urlObj.protocol === "https:";
}
