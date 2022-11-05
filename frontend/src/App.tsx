import NavSide from "./components/NavSide/NavSide";
import NavTop from "./components/NavTop/NavTop";
import { MantineProvider, AppShell } from "@mantine/core";
import { Route, Routes } from "react-router-dom";
import { routeMappings } from "./routes";

function App() {
	return (
		<MantineProvider
			theme={{ colorScheme: "dark" }}
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
	);
}

export default App;
