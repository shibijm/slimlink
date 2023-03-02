const nextConfig = {
	reactStrictMode: true,
	rewrites: () => [
		{
			source: "/:path((?!api/).*)",
			destination: "/api/links/:path/redirect",
		},
	],
};

export default nextConfig;
