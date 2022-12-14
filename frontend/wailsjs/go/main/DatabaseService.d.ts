// Cynhyrchwyd y ffeil hon yn awtomatig. PEIDIWCH Â MODIWL
// This file is automatically generated. DO NOT EDIT
import {ent} from '../models';

export function ActivateKey(arg1:number):Promise<ent.Key>;

export function CreateFile(arg1:ent.File):Promise<ent.File>;

export function CreateKey():Promise<ent.Key>;

export function CreateMasterPasswordHash(arg1:string):Promise<Error>;

export function CreatePasswordItem(arg1:ent.PasswordItem):Promise<ent.PasswordItem>;

export function DeactivateKey(arg1:number):Promise<ent.Key>;

export function DeleteFileByID(arg1:number):Promise<Error>;

export function DeleteKeyByID(arg1:number):Promise<Error>;

export function DeletePasswordItem(arg1:number):Promise<ent.PasswordItem>;

export function ReadAllActiveKeys():Promise<Array<ent.Key>>;

export function ReadAllFiles():Promise<Array<ent.File>>;

export function ReadAllKeys():Promise<Array<ent.Key>>;

export function ReadAllPasswordItems():Promise<Array<ent.PasswordItem>>;

export function ReadMasterPasswordHash():Promise<string>;

export function ReadSingleFile(arg1:number):Promise<ent.File>;

export function ReadSingleKey(arg1:number):Promise<ent.Key>;

export function ReadSinglePasswordItems(arg1:number):Promise<ent.PasswordItem>;

export function UpdatePasswordItem(arg1:ent.PasswordItem):Promise<ent.PasswordItem>;
