const base62Characters = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz";

export function generateBase62String(length: number): string {
	let output = "";
	while (output.length < length) {
		output += base62Characters.charAt(Math.random() * 61);
	}
	return output;
}
