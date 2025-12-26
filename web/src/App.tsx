import { PageFooter, PageMenu } from "@/components";
import MainView from "@/components/views/MainView";
import { defaultTheme, globalStyles } from "@/styles";
import { CssBaseline, GlobalStyles, ThemeProvider } from "@mui/material";
import Fade from "@mui/material/Fade";

export default function App() {
	return (
		<ThemeProvider defaultMode="system" theme={defaultTheme}>
			<CssBaseline />
			<GlobalStyles styles={globalStyles} />
			<PageMenu />
			<Fade in>
				<div>
					<MainView />
				</div>
			</Fade>
			<PageFooter />
		</ThemeProvider>
	);
}
