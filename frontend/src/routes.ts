import { TablerIcon } from '@tabler/icons';
import DataTable from './components/DataTable/DataTable';
import EncryptFiles from './pages/EncryptFiles/EncryptFiles';
import NewNote from './pages/Notes/NewNote';
import Notes from './pages/Notes/Notes';
import Notifications from './pages/Notifications/Notifications';
import PasswordManager from './pages/PasswordManager/PasswordManager';
import PasswordGenerator from './pages/PasswordGenerator/PasswordGenerator';
import Settings from './pages/Settings/Settings';

// TODO: route type may change depending on authentication logic
enum RouteType {
    Default,
    Account,
    General,
    Hidden,
}

export interface RouteData {
    path: string;
    component: () => JSX.Element;
    label?: string;
    icon?: TablerIcon;
    routeType?: RouteType;
}

export const routeMappings: RouteData[] = [
    { path: '/', component: PasswordGenerator, routeType: RouteType.General },
    {
        path: '/password_generator',
        component: PasswordGenerator,
        routeType: RouteType.Account,
    },
    {
        path: '/notifications',
        component: Notifications,
        routeType: RouteType.Account,
    },
    { path: '/settings', component: Settings, routeType: RouteType.Account },
    {
        path: '/password_manager',
        component: PasswordManager,
        routeType: RouteType.Account,
    },
    {
        path: '/encrypt_files',
        component: EncryptFiles,
        routeType: RouteType.Account,
    },
    { path: '/notes', component: Notes, routeType: RouteType.Account },
    { path: '/notes/new', component: NewNote },
];
