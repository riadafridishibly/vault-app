import { Code } from "@mantine/core";
import { useRecoilValue } from "recoil";
import { selectedFiles } from "./EncryptFilesControls";

function EncryptFiles() {
  const files = useRecoilValue(selectedFiles)
  return (
    <Code>
      {JSON.stringify(files, null, 4)}
    </Code>
  );
}

export default EncryptFiles;
