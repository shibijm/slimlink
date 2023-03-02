import InfoIcon from "@mui/icons-material/Info";
import { Button, Dialog, DialogActions, DialogContent, DialogTitle, Link, Stack } from "@mui/material";
import { useModal } from "hooks";
import { Fragment } from "react";

const links = {
	GitHub: "https://github.com/shibijm/slimlink",
	Copyright: "/static/copyright.txt",
	License: "/static/license.txt",
	"Third-Party Notices": "/static/notice.txt",
};

export default function AboutButton(): JSX.Element {
	const { isOpen, handleOpen, handleClose } = useModal();

	return (
		<Fragment>
			<Button endIcon={<InfoIcon />} onClick={handleOpen} sx={{ position: "fixed", right: 32, bottom: 32 }} variant="outlined">
				About
			</Button>
			<Dialog onClose={handleClose} open={isOpen}>
				<DialogTitle>Slimlink</DialogTitle>
				<DialogContent>
					<Stack alignItems="flex-start" gap={1}>
						{Object.entries(links).map(([label, href]) => (
							<Link href={href} key={label} target="_blank" underline="none">
								{label}
							</Link>
						))}
					</Stack>
				</DialogContent>
				<DialogActions>
					<Button onClick={handleClose}>Close</Button>
				</DialogActions>
			</Dialog>
		</Fragment>
	);
}
