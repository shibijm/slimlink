import type { NextApiRequest, NextApiResponse } from "next";
import { BadRequestError, NotFoundError } from "server/exceptions";
import { logError } from "server/logging/console";
import { addLink, getLinkByID } from "server/services/link";
import { Link } from "types";

export async function addLinkHandler(req: NextApiRequest, res: NextApiResponse): Promise<void> {
	if (req.method !== "POST") {
		res.status(405).end();
		return;
	}
	let link = null;
	try {
		link = await addLink(req.body.url);
	} catch (e) {
		if (e instanceof BadRequestError) {
			res.status(400).end();
		} else {
			logError(e, "controllers.link.addLinkHandler");
			res.status(500).end();
		}
		return;
	}
	res.status(200).json(link);
}

async function runWithLink(req: NextApiRequest, res: NextApiResponse, func: (link: Link | null) => void): Promise<void> {
	if (req.method !== "GET") {
		res.status(405).end();
		return;
	}
	if (typeof req.query.id !== "string") {
		res.status(400).end();
		return;
	}
	let link = null;
	try {
		link = await getLinkByID(req.query.id);
	} catch (e) {
		if (!(e instanceof NotFoundError)) {
			logError(e, "controllers.link.runWithLink");
			res.status(500).end();
			return;
		}
	}
	func(link);
}

export async function getLinkHandler(req: NextApiRequest, res: NextApiResponse): Promise<void> {
	await runWithLink(req, res, (link) => {
		if (!link) {
			res.status(404).end();
			return;
		}
		res.status(200).json(link);
	});
}

export async function redirectToLinkUrlHandler(req: NextApiRequest, res: NextApiResponse): Promise<void> {
	await runWithLink(req, res, (link) => {
		if (!link) {
			res.redirect("/404");
			return;
		}
		res.redirect(link.url);
	});
}
