import { CssBaseline, Experimental_CssVarsProvider as CssVarsProvider, GlobalStyles } from "@mui/material";
import Fade from "@mui/material/Fade";
import { PageFooter, PageMenu } from "components";
import type { AppProps } from "next/app";
import Head from "next/head";
import { Fragment } from "react";
import { defaultTheme, globalStyles } from "styles";

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
			<CssVarsProvider defaultMode="system" theme={defaultTheme}>
				<CssBaseline />
				<GlobalStyles styles={globalStyles} />
				<PageMenu />
				<Fade in>
					<div>
						<Component {...pageProps} />
					</div>
				</Fade>
				<PageFooter />
			</CssVarsProvider>
		</Fragment>
	);
}
