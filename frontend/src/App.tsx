import NavSide from './components/NavSide/NavSide';
import NavTop from './components/NavTop/NavTop';
import { MantineProvider, AppShell, ColorScheme, ColorSchemeProvider } from '@mantine/core';
import { createHashRouter, Outlet, redirect, RouterProvider } from 'react-router-dom';
import React from 'react';
import { NotificationsProvider } from '@mantine/notifications';
import Auth from './pages/Auth/Auth';
import PasswordGenerator from './pages/PasswordGenerator/PasswordGenerator';
import PasswordManager from './pages/PasswordManager/PasswordManager';
import EncryptFiles from './pages/EncryptFiles/EncryptFiles';
import Settings from './pages/Settings/Settings';
import Keys from './pages/Keys/Keys';
import { PwnChecker } from './pages/PwnChecker/PwnChecker';
import LoginError from './pages/LoginError/LoginError';
import { IsUserLoggedIn } from '@wailsjs/go/main/UserService';

async function loader({}) {
    // login guard
    const isLoggedIn = await IsUserLoggedIn();
    if (!isLoggedIn) {
        return redirect('/');
    }
}

const routes = createHashRouter([
    {
        path: '/',
        element: <Auth />,
    },
    {
        path: '/app',
        element: <AppOutlet />,
        errorElement: <LoginError />,
        loader: loader,
        children: [
            {
                element: <PasswordGenerator />,
                index: true,
            },
            {
                path: 'password_generator',
                element: <PasswordGenerator />,
            },
            {
                path: 'password_manager',
                element: <PasswordManager />,
                errorElement: <LoginError />,
            },
            {
                path: 'encrypt_files',
                element: <EncryptFiles />,
            },
            {
                path: 'settings',
                element: <Settings />,
            },
            {
                path: 'keys',
                element: <Keys />,
            },
            {
                path: 'password-checker',
                element: <PwnChecker />,
            },
        ],
    },
]);

function AppOutlet() {
    const [loggedIn, setIsLoggedIn] = React.useState(false);

    return (
        <AppShell
            padding="md"
            navbar={<NavSide />}
            header={<NavTop />}
            styles={(theme) => ({
                main: {
                    backgroundColor:
                        theme.colorScheme === 'dark' ? theme.colors.dark[8] : theme.colors.gray[0],
                },
            })}
        >
            <Outlet />
        </AppShell>
    );
}

function App() {
    const [colorScheme, setColorScheme] = React.useState<ColorScheme>('dark');
    const toggleColorScheme = (value?: ColorScheme) =>
        setColorScheme(value || (colorScheme === 'dark' ? 'light' : 'dark'));

    return (
        <ColorSchemeProvider colorScheme={colorScheme} toggleColorScheme={toggleColorScheme}>
            <MantineProvider theme={{ colorScheme }} withGlobalStyles withNormalizeCSS>
                <NotificationsProvider>
                    <RouterProvider router={routes} />
                </NotificationsProvider>
            </MantineProvider>
        </ColorSchemeProvider>
    );
}

export default App;
