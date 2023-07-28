import { useRef, useState } from "react";
import useMountEffect from "./useMountEffect";

export default function useDelayedLoading(
	initialLoading = true,
	delay = 100,
): { loading: boolean; showLoading: boolean; setLoading: (loading: boolean) => void } {
	const timeoutID = useRef(0);
	const [loading, setLoading] = useState(initialLoading);
	const [showLoading, setShowLoading] = useState(false);

	function setLoadingTimeout(): void {
		if (timeoutID.current) {
			window.clearTimeout(timeoutID.current);
		}
		timeoutID.current = window.setTimeout(() => {
			timeoutID.current = 0;
			setShowLoading(true);
		}, delay);
	}

	useMountEffect(() => {
		if (initialLoading) {
			setLoadingTimeout();
		}
		return () => {
			if (timeoutID.current) {
				window.clearTimeout(timeoutID.current);
			}
		};
	});

	function setLoadingWrapper(loading: boolean): void {
		if (loading) {
			setLoading(true);
			setLoadingTimeout();
		} else {
			if (timeoutID.current) {
				window.clearTimeout(timeoutID.current);
				timeoutID.current = 0;
			}
			setLoading(false);
			setShowLoading(false);
		}
	}

	return { loading, showLoading, setLoading: setLoadingWrapper };
}
