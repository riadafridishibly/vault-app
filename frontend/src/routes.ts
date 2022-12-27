import { PwnChecker } from './pages/PwnChecker/PwnChecker';
import {
    IconAlertOctagon,
    IconBellRinging,
    IconDatabaseImport,
    IconFilePower,
    IconKey,
    IconLock,
    IconNotes,
    IconSettings,
    IconShieldLock,
    TablerIcon,
} from '@tabler/icons';
import EncryptFiles from './pages/EncryptFiles/EncryptFiles';
import NewNote from './pages/Notes/NewNote';
import Notes from './pages/Notes/Notes';
import Notifications from './pages/Notifications/Notifications';
import PasswordManager from './pages/PasswordManager/PasswordManager';
import PasswordGenerator from './pages/PasswordGenerator/PasswordGenerator';
import Settings from './pages/Settings/Settings';
import Keys from './pages/Keys/Keys';
import { RouteType } from './shared/enums/route-type.enum';
import { RouteData } from './shared/interfaces/route-data.interface';
import Auth from './pages/Auth/Auth';

const nullTablerIcon = null as any as TablerIcon;

export const routeMappings: RouteData[] = [
    // {
    //     path: '/',
    //     component: PasswordGenerator,
    //     routeType: RouteType.General,
    //     label: 'Databases',
    //     icon: IconDatabaseImport,
    // },
    // {
    //     path: '/notifications',
    //     component: Notifications,
    //     routeType: RouteType.Account,
    //     label: 'Notifications',
    //     icon: IconBellRinging,
    // },
    // {
    //     path: '/notes',
    //     component: Notes,
    //     routeType: RouteType.Account,
    //     label: 'Notes',
    //     icon: IconNotes,
    // },
    {
        path: '/app/password_manager',
        component: PasswordManager,
        routeType: RouteType.Account,
        label: 'Password Manager',
        icon: IconShieldLock,
    },
    {
        path: '/app/encrypt_files',
        component: EncryptFiles,
        routeType: RouteType.Account,
        label: 'Encrypt Files',
        icon: IconFilePower,
    },
    {
        path: '/app/password_generator',
        component: PasswordGenerator,
        routeType: RouteType.Account,
        label: 'Password Generator',
        icon: IconLock,
    },
    {
        path: '/app/settings',
        component: Settings,
        routeType: RouteType.Account,
        label: 'Settings',
        icon: IconSettings,
    },
    // {
    //     path: '/notes/new',
    //     component: NewNote,
    //     routeType: RouteType.Hidden,
    //     icon: nullTablerIcon,
    // },
    {
        path: '/app/keys',
        component: Keys,
        routeType: RouteType.Account,
        label: 'Keys',
        icon: IconKey,
    },
    {
        path: '/app/password-checker',
        component: PwnChecker,
        routeType: RouteType.Account,
        label: 'Password Breach',
        icon: IconAlertOctagon,
    },
    {
        path: '/app/keys/new',
        component: Keys,
        routeType: RouteType.Hidden,
        label: 'Keys',
        icon: IconKey,
    },
];
