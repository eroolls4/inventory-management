export interface Inventory {
    id: number;
    name: string;
    description: string;
    created_at: string;
    updated_at: string;
}

export interface Item {
    id: number;
    inventory_id: number;
    inventory: Inventory;
    name: string;
    description: string;
    quantity: number;
    last_restock_at: string;
    created_at: string;
    updated_at: string;
}

export interface RestockHistory {
    id: number;
    item_id: number;
    restock_amount: number;
    restock_timestamp: string;
}

export interface User {
    id: number;
    Email: string;
}

export interface NewItem {
    inventory_id: string;
    name: string;
    description: string;
    quantity: string;
}