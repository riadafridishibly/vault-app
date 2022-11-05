import { Avatar, Badge } from "@mantine/core";

export function BadgeWithAvatar({ title }: { title: string }) {
	return (
		<>
			<Avatar
				alt='Avatar for badge'
				size={50}
				radius={25}
				mr={5}
				src='https://i.pravatar.cc/150?u=edrx'
			/>
			{/* TODO: Will Come back to this later */}
			<Badge
				sx={{ textTransform: "none" }}
				size='xl'
				radius='xl'
				color='teal'
			>
				{title}
			</Badge>
		</>
	);
}
