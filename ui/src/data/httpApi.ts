async function apiRequest<T>(method: string, endpoint: string, json?: unknown): Promise<T> {
	const headers: Record<string, string> = {};
	const fetchOptions: RequestInit = { method, headers };
	if (method === "POST") {
		headers["Content-Type"] = "application/json";
		if (json) {
			fetchOptions.body = JSON.stringify(json);
		}
	}
	let fetchResponse;
	try {
		fetchResponse = await fetch(`/api/${endpoint}`, fetchOptions);
	} catch {
		throw new Error("Connection failure");
	}
	let response!: T;
	try {
		response = await fetchResponse.json();
	} catch {
		if (fetchResponse.ok) {
			throw new Error("JSON parse failure");
		}
	}
	if (!fetchResponse.ok) {
		throw new Error(`HTTP ${fetchResponse.status}`);
	}
	return response;
}

export async function apiPost<T>(endpoint: string, json?: unknown): Promise<T> {
	return await apiRequest("POST", endpoint, json);
}
