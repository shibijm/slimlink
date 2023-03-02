import { useEffect, useState } from "react";

export default function useDebouncedValue<T>(value: T, delay = 300): T {
	const [debouncedValue, setDebouncedValue] = useState(value);

	useEffect(() => {
		const timeout = setTimeout(
			() => {
				setDebouncedValue(value);
			},
			value ? delay : 0,
		);
		return (): void => {
			clearTimeout(timeout);
		};
	}, [value, delay]);

	return debouncedValue;
}
