import { AppRouterCacheProvider } from "@mui/material-nextjs/v15-appRouter";
import { ThemeProvider } from "@mui/material/styles";
import type { Metadata } from "next";
import { Roboto } from "next/font/google";
import thema from "../thema";
import "./globals.css";

export const metadata: Metadata = {
	title: "検証用タスク管理アプリケーション",
	description: "検証用タスク管理アプリケーション",
    robots: {
        index: false,
        googleBot: {
            index: false
        }
    }
};

const roboto = Roboto({
	weight: ["300", "400", "500", "700"],
	subsets: ["latin"],
	display: "swap",
	variable: "--font-roboto",
});

export default function RootLayout({
	children,
}: Readonly<{
	children: React.ReactNode;
}>) {
	return (
		<html lang="ja">
			<body className={roboto.variable}>
				<AppRouterCacheProvider>
					<ThemeProvider theme={thema}>{children}</ThemeProvider>
				</AppRouterCacheProvider>
			</body>
		</html>
	);
}
