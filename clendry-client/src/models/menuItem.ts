export class MenuItem {
    name!: string; 
    title!: string; 
    short_title!: string; 
    href!: string; 
    link!: Link;
    is_heading!: boolean;
    icon_class!: string; 
    class_name!: string;
    expanded_icon_class!: string;
    collapse_icon_class!: string;
    children!: MenuItem;
}

export class Link {
    name!: string
}
