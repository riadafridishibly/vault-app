import DataTable from '@src/components/DataTable/DataTable';
import CreatePassword from '@src/shared/components/CreatePassword';
import { passwordDummy } from '@src/shared/data/dummy-data.data';
import { data, PasswordManagerTable } from './Table';

// <DataTable tableData={passwordDummy}></DataTable>
function PasswordManager() {
    return (
        <>
            <CreatePassword />
            <PasswordManagerTable data={data.data} />
        </>
    );
}

export default PasswordManager;
