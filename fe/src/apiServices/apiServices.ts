import axios from 'axios';

const apiBaseUrl = `${import.meta.env.VITE_APP_API_URL}/api`;


export const apiServices = {
    validateUser: () =>
        axios.get(`${import.meta.env.VITE_APP_API_URL}/validate`, { withCredentials: true }),

    fetchItems: () =>
        axios.get(`${apiBaseUrl}/items`, { withCredentials: true }),

    fetchInventories: () =>
        axios.get(`${apiBaseUrl}/inventory`, { withCredentials: true }),

    fetchRestockHistory: (itemId: number) =>
        axios.get(`${apiBaseUrl}/items/${itemId}/restock-history`, { withCredentials: true }),

    createItem: (itemData: {
        inventory_id: number;
        name: string;
        description: string;
        quantity: number;
    }) =>
        axios.post(`${apiBaseUrl}/items`, itemData, { withCredentials: true }),

    restockItem: (id: number, amount: number) =>
        axios.post(`${apiBaseUrl}/items/${id}/restock`, { amount }, { withCredentials: true }),

    deleteItem: (id: number) =>
        axios.delete(`${apiBaseUrl}/items/${id}`, { withCredentials: true }),

    signout: () =>
        axios.post(`${import.meta.env.VITE_APP_API_URL}/signout`, {}, { withCredentials: true }),
};