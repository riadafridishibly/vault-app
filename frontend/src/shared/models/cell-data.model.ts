export class CellData {
    value: string = '';
    isCopyable: boolean = false;
    isTags: boolean = false;
    constructor(val: string, isCopy: boolean = false, isTags: boolean = false) {
        this.value = val;
        this.isCopyable = isCopy;
        this.isTags = isTags;
    }
}
