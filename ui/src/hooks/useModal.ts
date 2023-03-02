import { useState } from "react";

export default function useModal(): {
	isOpen: boolean;
	setOpen: React.Dispatch<React.SetStateAction<boolean>>;
	handleOpen: () => void;
	handleClose: () => void;
} {
	const [isOpen, setOpen] = useState(false);

	function handleOpen(): void {
		setOpen(true);
	}

	function handleClose(): void {
		setOpen(false);
	}

	return { isOpen, setOpen, handleOpen, handleClose };
}
