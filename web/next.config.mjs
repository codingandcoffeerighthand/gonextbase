/** @type {import('next').NextConfig} */
const nextConfig = {
    reactStrictMode: false,
    output: "export",
    basePath: "/",
    skipTrailingSlashRedirect: false,
    distDir: process.env.DIST_DIR ?? "out",
};

export default nextConfig;
