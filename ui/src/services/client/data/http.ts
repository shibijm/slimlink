export async function apiRequest<T>(method: string, endpoint: string, json?: unknown): Promise<T> {
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
		fetchResponse = await fetch(`${localStorage.apiBase ?? ""}/api/${endpoint}`, fetchOptions);
	} catch {
		throw {
			error: {
				httpStatusCode: 0,
				message: "Connection failure",
			},
		};
	}
	let response!: T;
	try {
		response = await fetchResponse.json();
	} catch {
		if (fetchResponse.ok) {
			throw {
				error: {
					httpStatusCode: 0,
					message: "JSON parse failure",
				},
			};
		}
	}
	if (!fetchResponse.ok) {
		throw {
			error: {
				httpStatusCode: fetchResponse.status,
				message: `HTTP ${fetchResponse.status}`,
			},
		};
	}
	return response;
}
