/** @type {import('next').NextConfig} */
const nextConfig = {
  reactStrictMode: true,
  async rewrites() {
    return [
      // Rewriting to an external URL
      {
        source: '/api/:slug*',
        destination: 'http://127.0.0.1:3001/api/:slug*',
      },
    ];
  },
}

module.exports = nextConfig
