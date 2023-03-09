import { useRef } from "react";
import useMountEffect from "./useMountEffect";

export default function useMountStatus(): boolean {
	const isMounted = useRef(false);

	useMountEffect(() => {
		isMounted.current = true;
		return () => {
			isMounted.current = false;
		};
	});

	return isMounted.current;
}
