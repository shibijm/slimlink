import { version } from "@/config";
import { useMenu, useModal } from "@/hooks";
import BrightnessMediumIcon from "@mui/icons-material/BrightnessMedium";
import CheckIcon from "@mui/icons-material/Check";
import DarkModeIcon from "@mui/icons-material/DarkMode";
import InfoIcon from "@mui/icons-material/Info";
import LightModeIcon from "@mui/icons-material/LightMode";
import MenuIcon from "@mui/icons-material/Menu";
import {
	Button,
	Dialog,
	DialogActions,
	DialogContent,
	DialogTitle,
	Divider,
	Fade,
	IconButton,
	Link,
	ListItemIcon,
	ListItemText,
	Menu,
	MenuItem,
	Stack,
	Tooltip,
	Typography,
	useColorScheme,
} from "@mui/material";
import { Fragment } from "react";

const themes = {
	system: { label: "System Theme", icon: <BrightnessMediumIcon /> },
	light: { label: "Light Theme", icon: <LightModeIcon /> },
	dark: { label: "Dark Theme", icon: <DarkModeIcon /> },
};

const links = {
	GitHub: "https://github.com/shibijm/slimlink",
	"Third-Party Notices": "/static/notice.txt",
};

export default function PageMenu() {
	const pageMenu = useMenu();
	const aboutDialog = useModal();
	const { mode, setMode } = useColorScheme();

	return (
		<Fragment>
			<Fade in>
				<Tooltip enterDelay={500} placement="right" title="Menu">
					<IconButton
						onClick={pageMenu.handleOpen}
						sx={{
							position: "fixed",
							left: (theme) => theme.spacing(2),
							top: (theme) => theme.spacing(2),
						}}
					>
						<MenuIcon />
					</IconButton>
				</Tooltip>
			</Fade>
			<Menu anchorEl={pageMenu.anchorElement} onClose={pageMenu.handleClose} open={pageMenu.isOpen} slotProps={{ list: { sx: { width: "220px" } } }}>
				<Fragment>
					{Object.entries(themes).map(([theme, { label, icon }]) => (
						<MenuItem
							key={theme}
							onClick={(): void => {
								setMode(theme as "system" | "light" | "dark");
							}}
						>
							<ListItemIcon>{icon}</ListItemIcon>
							<ListItemText>{label}</ListItemText>
							{mode === theme && <CheckIcon color="primary" />}
						</MenuItem>
					))}
					<Divider />
				</Fragment>
				<MenuItem
					onClick={(): void => {
						aboutDialog.handleOpen();
						pageMenu.handleClose();
					}}
				>
					<ListItemIcon>
						<InfoIcon />
					</ListItemIcon>
					<ListItemText>About</ListItemText>
				</MenuItem>
			</Menu>
			<Dialog onClose={aboutDialog.handleClose} open={aboutDialog.isOpen}>
				<DialogTitle>Slimlink</DialogTitle>
				<DialogContent>
					<Stack alignItems="flex-start" gap={3}>
						<Typography>Version: {version}</Typography>
						<Stack alignItems="flex-start" gap={1}>
							{Object.entries(links).map(([label, href]) => (
								<Link href={href} key={label} target="_blank" underline="none">
									{label}
								</Link>
							))}
						</Stack>
					</Stack>
				</DialogContent>
				<DialogActions>
					<Button onClick={aboutDialog.handleClose}>Close</Button>
				</DialogActions>
			</Dialog>
		</Fragment>
	);
}
