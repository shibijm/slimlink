import { createTheme } from "@mui/material";

export const lightTheme = createTheme({
	typography: {
		fontFamily: ["Inter", "Roboto", "Segoe UI", "Arial", "sans-serif"].join(","),
	},
	palette: {
		mode: "light",
	},
});
