import Head from "next/head";

interface PageHeadProps {
	title: string;
	description: string;
}

export default function PageHead({ title, description }: PageHeadProps): JSX.Element {
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
