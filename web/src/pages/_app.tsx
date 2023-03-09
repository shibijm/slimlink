import { CssBaseline, GlobalStyles, ThemeProvider } from "@mui/material";
import { AboutButton } from "components";
import type { AppProps } from "next/app";
import Head from "next/head";
import { Fragment } from "react";
import { globalStyles } from "styles/global";
import { lightTheme } from "styles/themes";

export default function App({ Component, pageProps }: AppProps): JSX.Element {
	return (
		<Fragment>
			<Head>
				<meta charSet="UTF-8" />
				<meta content="width=device-width, initial-scale=1.0, shrink-to-fit=no" name="viewport" />
				<meta content="website" property="og:type" />
				<meta content="Slimlink" property="og:site_name" />
				{/* <meta content="https://example.com/" property="og:url" /> */}
				<meta content="https://raw.githubusercontent.com/shibijm/slimlink/master/web/src/assets/logo.svg" property="og:image" />
				<meta content="Logo" name="og:image:alt" />
				<meta content="summary" name="twitter:card" />
			</Head>
			<ThemeProvider theme={lightTheme}>
				<CssBaseline />
				<GlobalStyles styles={globalStyles} />
				<Component {...pageProps} />
				<AboutButton />
			</ThemeProvider>
		</Fragment>
	);
}
