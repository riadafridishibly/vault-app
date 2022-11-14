import { CreatePasswordItemForm } from "@src/shared/components/CreatePassword";
import { data, PasswordManagerTable } from "./Table";

// <DataTable tableData={passwordDummy}></DataTable>
function PasswordManager() {
    return (
        <>
            <CreatePasswordItemForm />
            <PasswordManagerTable />
        </>
    );
}

export default PasswordManager;
