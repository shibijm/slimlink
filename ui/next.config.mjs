const nextConfig = {
	reactStrictMode: true,
	rewrites: () => [
		{
			source: "/:path",
			destination: "/api/links/:path/redirect",
		},
	],
};

export default nextConfig;
