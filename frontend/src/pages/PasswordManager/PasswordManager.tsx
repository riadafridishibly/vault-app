import DataTable from '@src/components/DataTable/DataTable';
import CreatePassword from '@src/shared/components/CreatePassword';
import { passwordDummy } from '@src/shared/data/dummy-data.data';

function PasswordManager() {
    return (
        <>
            <CreatePassword />
            <DataTable tableData={passwordDummy}></DataTable>
        </>
    );
}

export default PasswordManager;
