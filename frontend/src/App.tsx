import NavSide from "./components/NavSide/NavSide";
import NavTop from "./components/NavTop/NavTop";
import PasswordGenerator from "./pages/PasswordGenerator/PasswordGenerator";
import { MantineProvider, AppShell } from "@mantine/core";

import styles from "./App.module.css";
import DataTable from "./components/DataTable/DataTable";
import { Route, Routes } from "react-router-dom";
import RandomPasswordGenerator from "./pages/RandomPasswordGenerator/RandomPasswordGenerator";

function App() {
  return (
    <MantineProvider
      theme={{ colorScheme: "dark" }}
      withGlobalStyles
      withNormalizeCSS
    >
      <AppShell
        padding="md"
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
          <Route path={"/"} element={<DataTable />} />
          <Route path={"/wpg"} element={<PasswordGenerator />} />
          <Route
            path={"/password_generator"}
            element={<RandomPasswordGenerator />}
          />
        </Routes>
      </AppShell>
    </MantineProvider>
  );
}

export default App;
