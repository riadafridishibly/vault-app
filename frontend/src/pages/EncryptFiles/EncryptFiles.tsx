import { Button, Center } from "@mantine/core";
import React from "react";
import { OpenDialog } from "../../../wailsjs/go/main/App";

function EncryptFiles() {
  const [files, setFiles] = React.useState<string[]>([]);
  const openFileDialog = () => {
    OpenDialog()
      .then((files) => setFiles((curr) => [...curr, ...files]))
      .catch((err) => console.error(err));
  };
  return (
    <>
      <Center>{JSON.stringify(files)}</Center>
      <Button
        onClick={(e: React.MouseEvent) => {
          e.preventDefault();
          openFileDialog();
        }}
      >
        Open Files
      </Button>
    </>
  );
}

export default EncryptFiles;
