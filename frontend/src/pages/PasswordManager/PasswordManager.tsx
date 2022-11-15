import { PasswordItemForm } from '@src/shared/components/PasswordItemModal';
import PwnCheckModal from '@src/shared/components/PwnCheckModal';
import { data, PasswordManagerTable } from './Table';

// <DataTable tableData={passwordDummy}></DataTable>
function PasswordManager() {
    return (
        <>
            <PwnCheckModal />
            <PasswordItemForm />
            <PasswordManagerTable />
        </>
    );
}

export default PasswordManager;
