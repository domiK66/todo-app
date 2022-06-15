import '../styles/globals.css'
import type { AppProps } from 'next/app'
import Head from 'next/head'
import { Navbar } from '../components/Navbar';
import { Footer } from '../components/Footer';

function MyApp({ Component, pageProps }: AppProps) {
  return (
    <>
    <Head>
      <title>todo-frontend</title>
      <meta name="description" content="a simple todo app" />
      <link rel="icon" href="/favicon.ico" />
    </Head>

    <Navbar />
    <div className="max-w-7xl mx-auto px-2 sm:px-6 lg:px-8">
      <Component {...pageProps} />
    </div>
    </>
  );
}

export default MyApp
