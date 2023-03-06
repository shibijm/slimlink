import ErrorIcon from "@mui/icons-material/Error";
import { Button, Divider, Stack, Typography } from "@mui/material";
import { PageHeader } from "components";
import Link from "next/link";
import { Fragment } from "react";

export default function NotFoundView(): JSX.Element {
	return (
		<Fragment>
			<PageHeader description="Not found" title="Slimlink - HTTP 404" />
			<Stack alignItems="center" justifyContent="center" spacing={5}>
				<Stack alignItems="center" spacing={2} width={350}>
					<ErrorIcon color="error" fontSize="large" />
					<Typography variant="h4">HTTP 404</Typography>
					<Divider flexItem />
					<Typography>The requested page could not be found</Typography>
				</Stack>
				<Button LinkComponent={Link} href="/" variant="outlined">
					Home
				</Button>
			</Stack>
		</Fragment>
	);
}
