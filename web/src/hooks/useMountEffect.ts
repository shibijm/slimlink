import { useEffect } from "react";

export default function useMountEffect(callback: React.EffectCallback): void {
	// eslint-disable-next-line react-hooks/exhaustive-deps
	useEffect(callback, []);
}
