import { CssBaseline, Experimental_CssVarsProvider as CssVarsProvider, GlobalStyles } from "@mui/material";
import { PageFooter, ThemeToggler } from "components";
import type { AppProps } from "next/app";
import Head from "next/head";
import { Fragment } from "react";
import { extendedTheme, globalStyles } from "styles";

export default function App({ Component, pageProps }: AppProps): JSX.Element {
	return (
		<Fragment>
			<Head>
				<meta charSet="UTF-8" />
				<meta content="width=device-width, initial-scale=1.0, shrink-to-fit=no" name="viewport" />
				<meta content="website" property="og:type" />
				<meta content="Slimlink" property="og:site_name" />
				{/* <meta content="https://slimlink.vercel.app/" property="og:url" /> */}
				<meta content="https://raw.githubusercontent.com/shibijm/slimlink/master/web/public/static/media/logo.png" property="og:image" />
				<meta content="Logo" name="og:image:alt" />
				<meta content="summary" name="twitter:card" />
			</Head>
			<CssVarsProvider defaultMode="system" theme={extendedTheme}>
				<CssBaseline />
				<GlobalStyles styles={globalStyles} />
				<ThemeToggler />
				<Component {...pageProps} />
				<PageFooter />
			</CssVarsProvider>
		</Fragment>
	);
}
