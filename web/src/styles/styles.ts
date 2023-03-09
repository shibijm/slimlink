import { experimental_extendTheme as extendTheme } from "@mui/material";

export const defaultTheme = extendTheme({
	typography: {
		fontFamily: "Inter, Roboto, 'Segoe UI', Arial, sans-serif",
	},
	transitions: {
		duration: {
			enteringScreen: 250,
		},
	},
	components: {
		MuiButton: {
			styleOverrides: {
				root: {
					minWidth: "auto",
				},
			},
		},
	},
});

export const globalStyles = {
	body: {
		padding: defaultTheme.spacing(4),
	},
};
