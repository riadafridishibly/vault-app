import DeletePasswordItemModal from '@src/shared/components/DeletePasswordItemModal';
import { PasswordItemForm } from '@src/shared/components/PasswordItemModal';
import PwnCheckModal from '@src/shared/components/PwnCheckModal';
import { PasswordManagerTable } from './Table';

// <DataTable tableData={passwordDummy}></DataTable>
function PasswordManager() {
    return (
        <>
            <PwnCheckModal />
            <PasswordItemForm />
            <DeletePasswordItemModal />
            <PasswordManagerTable />
        </>
    );
}

export default PasswordManager;
