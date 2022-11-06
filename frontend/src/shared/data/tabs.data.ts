import {
	IconBellRinging,
	IconDatabaseImport,
	IconFilePower,
	IconKey,
	IconLock,
	IconLockAccess,
	IconNotes,
	IconSettings,
	IconShieldLock,
} from "@tabler/icons";

export const tabs = {
	account: [
		{
			link: "/notifications",
			label: "Notifications",
			icon: IconBellRinging,
		},
		{ link: "/notes", label: "Notes", icon: IconNotes },
		{
			link: "/password_manager",
			label: "Password Manager",
			icon: IconShieldLock,
		},
		{
			link: "/password_generator",
			label: "Password Generator",
			icon: IconLock,
		},
		{ link: "/encrypt_files", label: "Encrypt Files", icon: IconFilePower },
		{ link: "/settings", label: "Settings", icon: IconSettings },
		{ link: "/keys", label: "Keys", icon: IconKey },
	],
	general: [{ link: "/", label: "Databases", icon: IconDatabaseImport }],
};
