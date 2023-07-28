import { useState } from "react";

export default function useMenu(): {
	isOpen: boolean;
	anchorElement: HTMLButtonElement | null;
	handleOpen: (e: React.MouseEvent<HTMLButtonElement>) => void;
	handleClose: () => void;
} {
	const [anchorElement, setAnchorElement] = useState<HTMLButtonElement | null>(null);

	function handleOpen(e: React.MouseEvent<HTMLButtonElement>): void {
		setAnchorElement(e.currentTarget);
	}

	function handleClose(): void {
		setAnchorElement(null);
	}

	return { isOpen: !!anchorElement, anchorElement, handleOpen, handleClose };
}
