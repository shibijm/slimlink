import ClearIcon from "@mui/icons-material/Clear";
import ShortcutIcon from "@mui/icons-material/Shortcut";
import { LoadingButton } from "@mui/lab";
import { IconButton, InputAdornment, Stack, TextField, Typography } from "@mui/material";
import logo from "assets/logo.svg";
import { PageHead } from "components";
import { useControlledInput, useDebouncedValue, useDelayedLoading, useMountEffect } from "hooks";
import { Fragment, useEffect, useState } from "react";
import { createLink } from "services/link";
import { isValidUrl } from "utils";
import ResultAlert from "./ResultAlert";

export default function MainView(): JSX.Element {
	const url = useControlledInput("");
	const [urlFieldError, setUrlFieldError] = useState("");
	const [lastUrl, setLastUrl] = useState("");
	const [shortenedUrl, setShortenedUrl] = useState("");
	const [error, setError] = useState("");
	const [showResultAlert, setShowResultAlert] = useState(false);
	const { showLoading, setLoading } = useDelayedLoading(false);
	const urlDebounced = useDebouncedValue(url.value);

	useMountEffect(url.focus);

	useEffect(() => {
		if (urlDebounced && !isValidUrl(urlDebounced)) {
			setUrlFieldError("Invalid URL");
		} else {
			setUrlFieldError("");
		}
	}, [urlDebounced]);

	function shortenUrl(): void {
		if (!isValidUrl(url.value)) {
			setUrlFieldError("Invalid URL");
			url.focus();
			return;
		}
		setLoading(true);
		setShowResultAlert(false);
		setUrlFieldError("");
		createLink(url.value)
			.then((link) => {
				setLastUrl(link.url);
				setShortenedUrl(`${location.origin}/${link.id}`);
				setError("");
			})
			.catch((error) => {
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
		<Fragment>
			<PageHead description="URL shortener" title="Slimlink" />
			<Stack alignItems="center" gap={3} height="100%">
				<Stack alignItems="center" gap={1}>
					<img alt="Logo" height={32} src={logo.src} width={32} />
					<Stack alignItems="center">
						<Typography variant="h5">Slimlink</Typography>
						<Typography variant="body2">URL Shortener</Typography>
					</Stack>
				</Stack>
				<TextField
					FormHelperTextProps={{ sx: { margin: 0, height: 0 } }}
					InputProps={url.value ? urlFieldConditionalInputProps : undefined}
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
					sx={{ maxWidth: 600 }}
					variant="filled"
					{...url.bind}
				/>
				<LoadingButton loading={showLoading} loadingPosition="start" onClick={shortenUrl} startIcon={<ShortcutIcon />} variant="contained">
					Shorten
				</LoadingButton>
				<ResultAlert {...{ showResultAlert, setShowResultAlert, lastUrl, shortenedUrl, error }} />
			</Stack>
		</Fragment>
	);
}
