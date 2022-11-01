import { useLocation, useParams } from "react-router-dom";

function NewNote() {
  const { state } = useLocation();
  return (
    <>
      <div>Create New Note</div>
      <div>{JSON.stringify(state)}</div>
    </>
  );
}

export default NewNote;
