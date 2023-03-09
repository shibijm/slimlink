import type { NextApiRequest, NextApiResponse } from "next";
import { getInfo } from "server/services/info";

export function getInfoHandler(req: NextApiRequest, res: NextApiResponse): void {
	if (req.method !== "GET") {
		res.status(405).end();
		return;
	}
	const info = getInfo();
	res.status(200).json(info);
}
