import {
  createStyles,
  Header,
  Autocomplete,
  Group,
  Burger,
  Button,
} from "@mantine/core";
import { useDisclosure } from "@mantine/hooks";
import { IconAbacus, IconPlus, IconSearch } from "@tabler/icons";
import React from "react";
import { useLocation, useNavigate } from "react-router-dom";

const useStyles = createStyles((theme) => ({
  header: {
    paddingLeft: theme.spacing.md,
    paddingRight: theme.spacing.md,
  },

  inner: {
    height: 56,
    display: "flex",
    justifyContent: "space-between",
    alignItems: "center",
  },

  links: {
    [theme.fn.smallerThan("md")]: {
      display: "none",
    },
  },

  search: {
    [theme.fn.smallerThan("xs")]: {
      display: "none",
    },
  },

  link: {
    display: "block",
    lineHeight: 1,
    padding: "8px 12px",
    borderRadius: theme.radius.sm,
    textDecoration: "none",
    color:
      theme.colorScheme === "dark"
        ? theme.colors.dark[0]
        : theme.colors.gray[7],
    fontSize: theme.fontSizes.sm,
    fontWeight: 500,

    "&:hover": {
      backgroundColor:
        theme.colorScheme === "dark"
          ? theme.colors.dark[6]
          : theme.colors.gray[0],
    },
  },
}));

interface HeaderSearchProps {
  links: { link: string; label: string }[];
}

function HeaderSearch({ links }: HeaderSearchProps) {
  const [opened, { toggle }] = useDisclosure(false);
  const { classes } = useStyles();
  const loc = useLocation();
  const navigate = useNavigate();

  const items = links.map((link) => (
    <a
      key={link.label}
      href={link.link}
      className={classes.link}
      onClick={(event) => event.preventDefault()}
    >
      {link.label}
    </a>
  ));

  return (
    <Header height={56} className={classes.header} mb={120} fixed>
      <div className={classes.inner}>
        <Group>
          <Burger opened={opened} onClick={toggle} size="sm" />
          <IconAbacus size={28} />
        </Group>

        <Group>
          <Group ml={50} spacing={5} className={classes.links}>
            {items}
          </Group>
          <Autocomplete
            className={classes.search}
            placeholder="Search"
            icon={<IconSearch size={16} stroke={1.5} />}
            data={[
              "React",
              "Angular",
              "Vue",
              "Next.js",
              "Riot.js",
              "Svelte",
              "Blitz.js",
            ]}
          />
          {loc && loc.pathname === "/notes" && (
            <Button
              onClick={(e: React.MouseEvent) => {
                e.preventDefault();
                navigate("/new_note", { state: { hello: "world" } });
              }}
              leftIcon={<IconPlus />}
            >
              New Note
            </Button>
          )}
        </Group>
      </div>
    </Header>
  );
}

function NavTop() {
  const loc = useLocation();
  React.useEffect(() => {
    console.log("location:", loc);
  }, [loc]);
  return (
    <HeaderSearch
      links={[
        {
          link: "/about",
          label: "Features",
        },
        {
          link: "/pricing",
          label: "Pricing",
        },
        {
          link: "/learn",
          label: "Learn",
        },
        {
          link: "/community",
          label: "Community",
        },
      ]}
    />
  );
}

export default NavTop;
