import { Stack, Typography } from "@mui/material";
import { AboutButton } from "components";
import { useMountEffect } from "hooks";
import { useState } from "react";
import { getInfo } from "services/info";
import { Info } from "types";

export default function PageFooter(): JSX.Element {
	const [info, setInfo] = useState<Info>();

	useMountEffect(() => {
		getInfo()
			.then(setInfo)
			.catch((error) => {
				setInfo({ pageFooterText: `Failed to fetch info - ${error.message}` });
			});
	});

	return (
		<Stack
			alignItems="center"
			direction="row"
			gap={3}
			justifyContent="space-between"
			sx={{
				position: "fixed",
				left: 0,
				bottom: 0,
				width: "100%",
				padding: 3,
			}}
		>
			<Typography>{info?.pageFooterText}</Typography>
			<AboutButton />
		</Stack>
	);
}
