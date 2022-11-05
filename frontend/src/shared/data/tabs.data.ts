import {
	IconBellRinging,
	IconDatabaseImport,
	IconFileAnalytics,
	IconFilePower,
	IconKey,
	IconLicense,
	IconLockAccess,
	IconMessage2,
	IconMessages,
	IconNotes,
	IconReceiptRefund,
	IconSettings,
	IconShoppingCart,
	IconUsers,
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
			icon: IconLockAccess,
		},
		{
			link: "/password_generator",
			label: "Password Generator",
			icon: IconKey,
		},
		{ link: "/encrypt_files", label: "Encrypt Files", icon: IconFilePower },
		{ link: "/settings", label: "Settings", icon: IconSettings },
	],
	general: [{ link: "/", label: "Databases", icon: IconDatabaseImport }],
};
