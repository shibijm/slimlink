import { useRef, useState } from "react";
import useMountEffect from "./useMountEffect";

export default function useDelayedLoading(
	initialLoading = true,
	delay = 100,
): { loading: boolean; showLoading: boolean; setLoading: (loading: boolean) => void } {
	const loadingTimeout = useRef<NodeJS.Timeout | null>(null);
	const [loading, setLoading] = useState(initialLoading);
	const [showLoading, setShowLoading] = useState(false);

	function setLoadingTimeout(): void {
		if (loadingTimeout.current !== null) {
			clearTimeout(loadingTimeout.current);
		}
		loadingTimeout.current = setTimeout(() => {
			loadingTimeout.current = null;
			setShowLoading(true);
		}, delay);
	}

	useMountEffect(() => {
		if (initialLoading) {
			setLoadingTimeout();
		}
		return () => {
			if (loadingTimeout.current !== null) {
				clearTimeout(loadingTimeout.current);
			}
		};
	});

	function setLoadingWrapper(loading: boolean): void {
		if (loading) {
			setLoading(true);
			setLoadingTimeout();
		} else {
			if (loadingTimeout.current !== null) {
				clearTimeout(loadingTimeout.current);
				loadingTimeout.current = null;
			}
			setLoading(false);
			setShowLoading(false);
		}
	}

	return { loading, showLoading, setLoading: setLoadingWrapper };
}
