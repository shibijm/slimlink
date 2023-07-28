import { Fade, Typography } from "@mui/material";
import { useMountEffect } from "hooks";
import { useState } from "react";
import { getInfo } from "services/info";
import { Info } from "types";

export default function PageFooter(): JSX.Element | null {
	const [info, setInfo] = useState<Info>();

	useMountEffect(() => {
		getInfo()
			.then(setInfo)
			.catch((error) => {
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
