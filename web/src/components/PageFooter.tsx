import { useMountEffect } from "@/hooks";
import { getInfo } from "@/services/info";
import type { Info } from "@/types";
import { Fade, Typography } from "@mui/material";
import { useState } from "react";

export default function PageFooter() {
	const [info, setInfo] = useState<Info>();

	useMountEffect(() => {
		getInfo()
			.then(setInfo)
			.catch((error: Error) => {
				setInfo({ pageFooterText: `Failed to fetch info - ${error.message}` });
			});
	});

	return info ? (
		<Fade in>
			<Typography
				sx={{
					position: "fixed",
					left: (theme) => theme.spacing(3),
					bottom: (theme) => theme.spacing(3),
				}}
			>
				{info.pageFooterText}
			</Typography>
		</Fade>
	) : null;
}
