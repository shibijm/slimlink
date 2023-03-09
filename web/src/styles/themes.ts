import { experimental_extendTheme as extendTheme } from "@mui/material";

const commonTheme = {
	typography: {
		fontFamily: "Inter, Roboto, 'Segoe UI', Arial, sans-serif",
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
};

export const extendedTheme = extendTheme(commonTheme);
