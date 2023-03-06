import ContentCopyIcon from "@mui/icons-material/ContentCopy";
import { Alert, AlertTitle, Button, Collapse, Link, Stack, Typography } from "@mui/material";
import { Fragment, useRef, useState } from "react";

interface ResultAlertProps {
	showResultAlert: boolean;
	setShowResultAlert: React.Dispatch<React.SetStateAction<boolean>>;
	lastUrl: string;
	shortenedUrl: string;
	error: string;
}

export default function ResultAlert({ showResultAlert, setShowResultAlert, lastUrl, shortenedUrl, error }: ResultAlertProps): JSX.Element {
	const [copyButtonLabel, setCopyButtonLabel] = useState("Copy");
	const copyButtonLabelTimeout = useRef<NodeJS.Timeout | null>(null);

	return (
		<Collapse
			in={showResultAlert}
			onEnter={(): void => {
				setCopyButtonLabel("Copy");
			}}
			sx={{ maxWidth: "100%", overflowWrap: "anywhere" }}
		>
			<Alert
				onClose={(): void => {
					setShowResultAlert(false);
				}}
				severity={error ? "error" : "success"}
			>
				{error ? (
					<Fragment>
						<AlertTitle>Something went wrong</AlertTitle>
						{error}
					</Fragment>
				) : (
					<Fragment>
						<AlertTitle>URL shortened!</AlertTitle>
						{lastUrl}
						<Stack alignItems="flex-start" gap={0.5} marginTop={2}>
							<Link href={shortenedUrl} target="_blank" underline="none">
								<Typography>{shortenedUrl}</Typography>
							</Link>
							<Button
								onClick={(): void => {
									navigator.clipboard.writeText(shortenedUrl);
									setCopyButtonLabel("Copied");
									if (copyButtonLabelTimeout.current !== null) {
										clearTimeout(copyButtonLabelTimeout.current);
									}
									copyButtonLabelTimeout.current = setTimeout(() => {
										copyButtonLabelTimeout.current = null;
										setCopyButtonLabel("Copy");
									}, 1000);
								}}
								size="small"
								startIcon={<ContentCopyIcon />}
							>
								{copyButtonLabel}
							</Button>
						</Stack>
					</Fragment>
				)}
			</Alert>
		</Collapse>
	);
}
