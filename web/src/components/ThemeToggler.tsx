import DarkModeIcon from "@mui/icons-material/DarkMode";
import LightModeIcon from "@mui/icons-material/LightMode";
import { Fade, IconButton, Tooltip, useColorScheme } from "@mui/material";
import { useMountStatus } from "hooks";

export default function ThemeToggler(): JSX.Element | null {
	const isMounted = useMountStatus();
	const { mode, systemMode, setMode } = useColorScheme();

	if (!isMounted) {
		return null;
	}

	const theme = mode === "system" ? systemMode : mode;
	const nextTheme = theme === "dark" ? "light" : "dark";

	return (
		<Fade in>
			<Tooltip placement="left" title={`Switch to ${nextTheme} theme`}>
				<IconButton
					onClick={(): void => {
						setMode(nextTheme);
					}}
					sx={{
						position: "fixed",
						right: 16,
						top: 16,
					}}
				>
					{theme === "dark" ? <LightModeIcon /> : <DarkModeIcon />}
				</IconButton>
			</Tooltip>
		</Fade>
	);
}
