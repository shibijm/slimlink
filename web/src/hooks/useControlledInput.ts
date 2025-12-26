import { useReducer, useRef, useState } from "react";

interface ControlledInput {
	bind: {
		value: string;
		onChange: (e: React.ChangeEvent<HTMLInputElement>) => void;
		inputRef: React.RefObject<HTMLInputElement | null>;
	};
	value: string;
	setValue: (value: string) => void;
	ref: React.RefObject<HTMLInputElement | null>;
	focus: () => void;
	blur: () => void;
}

type Values = Record<string, string>;

type ControlledInputs<T> = { [K in keyof T]: ControlledInput };

function composeControlledInput(value: string, setValue: (value: string) => void): ControlledInput {
	// eslint-disable-next-line react-hooks/rules-of-hooks
	const ref = useRef<HTMLInputElement | null>(null);
	return {
		bind: {
			value,
			onChange: (e: React.ChangeEvent<HTMLInputElement>): void => {
				setValue(e.target.value);
			},
			inputRef: ref,
		},
		value,
		setValue,
		ref,
		focus: (): void => {
			if (ref.current) {
				ref.current.focus();
			}
		},
		blur: (): void => {
			if (ref.current) {
				ref.current.blur();
			}
		},
	};
}

function valuesReducer(values: Values, [key, value]: [string, string]): Values {
	return { ...values, [key]: value };
}

export function useControlledInputs<T extends Values>(initialValues: T): ControlledInputs<T> {
	const [values, dispatch] = useReducer(valuesReducer, initialValues);
	return Object.entries(values).reduce(
		(controlledInputs, [key, value]) => ({
			...controlledInputs,
			[key]: composeControlledInput(value, (value) => {
				dispatch([key, value]);
			}),
		}),
		{} as ControlledInputs<T>,
	);
}

export function useControlledInput(initialValue: string): ControlledInput {
	return composeControlledInput(...useState(initialValue));
}
