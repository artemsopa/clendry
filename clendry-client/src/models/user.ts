export default interface User {
    id: string;
    nick: string;
    email: string;
    password: string;
    downloads: number;
    uploads: number;
    memory: number;
}