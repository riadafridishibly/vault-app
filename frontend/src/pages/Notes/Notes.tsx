import { useState } from "react";
import { RichTextEditor } from "@mantine/rte";

const initialValue =
  "<p>Your initial <b>html value</b> or an empty string to init editor without value</p>";

export default function Notes() {
  const [value, onChange] = useState(initialValue);
  return (
    <RichTextEditor
      sticky={true}
      // stickyOffset={"100px"}
      controls={[]}
      value={value}
      onChange={onChange}
      id="rte"
    />
  );
}
