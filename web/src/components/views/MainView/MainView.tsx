import logo from "@/assets/logo.svg";
import { useControlledInput, useDebouncedValue, useDelayedLoading, useMountEffect } from "@/hooks";
import { createLink } from "@/services/link";
import { isValidUrl } from "@/utils";
import ClearIcon from "@mui/icons-material/Clear";
import ShortcutIcon from "@mui/icons-material/Shortcut";
import { Button, IconButton, InputAdornment, Stack, TextField, Typography } from "@mui/material";
import { useState } from "react";
import ResultAlert from "./ResultAlert";

export default function MainView() {
	const url = useControlledInput("");
	const [lastUrl, setLastUrl] = useState("");
	const [shortenedUrl, setShortenedUrl] = useState("");
	const [error, setError] = useState("");
	const [showResultAlert, setShowResultAlert] = useState(false);
	const { showLoading, setLoading } = useDelayedLoading(false);
	const urlDebounced = useDebouncedValue(url.value);
	const urlFieldError = urlDebounced && !isValidUrl(urlDebounced) ? "Invalid URL" : "";

	useMountEffect(url.focus);

	function shortenUrl(): void {
		if (!isValidUrl(url.value)) {
			url.focus();
			return;
		}
		setLoading(true);
		setShowResultAlert(false);
		createLink(url.value)
			.then((link) => {
				setLastUrl(link.url);
				setShortenedUrl(`${location.origin}/${link.id}`);
				setError("");
			})
			.catch((error: Error) => {
				setShortenedUrl("");
				setError(error.message);
			})
			.finally(() => {
				setShowResultAlert(true);
				setLoading(false);
			});
	}

	const urlFieldConditionalInputProps = {
		endAdornment: (
			<InputAdornment position="end">
				<IconButton
					disabled={showLoading}
					edge="end"
					onClick={(): void => {
						url.setValue("");
						url.focus();
					}}
					size="small"
					title="Clear"
				>
					<ClearIcon />
				</IconButton>
			</InputAdornment>
		),
	};

	return (
		<Stack alignItems="center" gap={3} height="100%">
			<Stack alignItems="center" gap={1}>
				<img alt="Logo" height={32} src={logo} width={32} />
				<Stack alignItems="center">
					<Typography variant="h5">Slimlink</Typography>
					<Typography variant="body2">URL Shortener</Typography>
				</Stack>
			</Stack>
			<TextField
				disabled={showLoading}
				error={!!urlFieldError}
				fullWidth
				helperText={urlFieldError}
				label="URL"
				onKeyDown={(e): void => {
					if (e.key === "Enter") {
						url.blur();
						shortenUrl();
					} else if (e.key === "Escape") {
						url.setValue("");
					}
				}}
				slotProps={{
					formHelperText: { sx: { margin: 0, height: 0 } },
					input: url.value ? urlFieldConditionalInputProps : undefined,
				}}
				sx={{ maxWidth: 600 }}
				variant="filled"
				{...url.bind}
			/>
			<Button loading={showLoading} loadingPosition="start" onClick={shortenUrl} startIcon={<ShortcutIcon />} variant="contained">
				Shorten
			</Button>
			<ResultAlert {...{ showResultAlert, setShowResultAlert, lastUrl, shortenedUrl, error }} />
		</Stack>
	);
}
