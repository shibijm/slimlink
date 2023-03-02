import Head from "next/head";

interface PageHeaderProps {
	title: string;
	description: string;
}

export default function PageHeader({ title, description }: PageHeaderProps): JSX.Element {
	return (
		<Head>
			<title>{title}</title>
			<meta content={title} name="title" />
			<meta content={description} name="description" />
			<meta content={title} property="og:title" />
			<meta content={description} property="og:description" />
		</Head>
	);
}
