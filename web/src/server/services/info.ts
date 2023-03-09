import { Info } from "types";

const pageFooterText = process.env.PAGE_FOOTER_TEXT || "";

export function getInfo(): Info {
	return { pageFooterText };
}
