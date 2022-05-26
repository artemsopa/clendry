export default class FileServ {
    id!: string;
    c_type!: string;
    created_at!: string;
    folder_id!: string;
    is_fav!: boolean;
    is_trash!: boolean;
    size!: number;
    title!: string;
    type!: string;
    url!: string;
    isEdit: boolean = false;
    png: string = "";
    member_id?: string;
}