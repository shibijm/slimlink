import { getInitColorSchemeScript } from "@mui/material";
import { Head, Html, Main, NextScript } from "next/document";

export default function Document(): JSX.Element {
	return (
		<Html lang="en">
			<Head>
				<link href="https://fonts.googleapis.com" rel="preconnect" />
				<link crossOrigin="anonymous" href="https://fonts.gstatic.com" rel="preconnect" />
				<link href="https://fonts.googleapis.com/css2?family=Inter:wght@300;400;500;700&display=swap" rel="stylesheet" />
			</Head>
			<body>
				{getInitColorSchemeScript({ defaultMode: "system" })}
				<Main />
				<NextScript />
			</body>
		</Html>
	);
}
