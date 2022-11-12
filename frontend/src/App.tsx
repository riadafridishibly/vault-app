import NavSide from "./components/NavSide/NavSide";
import NavTop from "./components/NavTop/NavTop";
import { MantineProvider, AppShell, ColorScheme, ColorSchemeProvider } from "@mantine/core";
import { Route, Routes } from "react-router-dom";
import { routeMappings } from "./routes";
import React from "react";

function App() {
	const [colorScheme, setColorScheme] = React.useState<ColorScheme>('dark');
	const toggleColorScheme = (value?: ColorScheme) =>
		setColorScheme(value || (colorScheme === 'dark' ? 'light' : 'dark'));

	return (
		<ColorSchemeProvider colorScheme={colorScheme} toggleColorScheme={toggleColorScheme}>
			<MantineProvider
				theme={{ colorScheme }}
				withGlobalStyles
				withNormalizeCSS
			>
				<AppShell
					padding='md'
					navbar={<NavSide />}
					header={<NavTop />}
					styles={(theme) => ({
						main: {
							backgroundColor:
								theme.colorScheme === "dark"
									? theme.colors.dark[8]
									: theme.colors.gray[0],
						},
					})}
				>
					<Routes>
						{routeMappings.map((route, index) => {
							return (
								<Route
									key={index}
									path={route.path}
									element={route.component()}
								/>
							);
						})}
					</Routes>
				</AppShell>
			</MantineProvider>
		</ColorSchemeProvider>

	);
}

export default App;
