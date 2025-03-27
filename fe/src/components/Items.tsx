import { useState, useEffect } from 'react';
import { useNavigate } from 'react-router-dom';
import {apiServices} from "../apiServices/apiServices.ts";
import {Inventory, Item, NewItem, RestockHistory, User} from "../IProps/interfaces.ts";

const Items: React.FC = () => {
    const [items, setItems] = useState<Item[]>([]);
    const [history, setHistory] = useState<RestockHistory[]>([]);
    const [inventories, setInventories] = useState<Inventory[]>([]);
    const [user, setUser] = useState<User | null>(null);
    const [showHistoryModal, setShowHistoryModal] = useState(false);
    const [showRestockModal, setShowRestockModal] = useState(false);
    const [showCreateModal, setShowCreateModal] = useState(false);
    const [selectedItemId, setSelectedItemId] = useState<number | null>(null);
    const [restockAmount, setRestockAmount] = useState('');
    const [newItem, setNewItem] = useState<NewItem>({ inventory_id: '', name: '', description: '', quantity: '' });
    const [errorMessage, setErrorMessage] = useState<string | null>(null);
    const [restockAttempts, setRestockAttempts] = useState<{ [key: number]: number }>({});
    const navigate = useNavigate();

    useEffect(() => {
        fetchUser();
        fetchItems();
        fetchInventories();
    }, []);

    const fetchUser = async () => {
        try {
            const response = await apiServices.validateUser();
            setUser(response.data.user as User);
        } catch (error) {
            console.error('Error fetching user:', error);
            navigate('/login');
        }
    };

    const fetchItems = async () => {
        try {
            const response = await apiServices.fetchItems();
            setItems(response.data);
        } catch (error) {
            console.error('Error fetching items:', error);
        }
    };

    const fetchInventories = async () => {
        try {
            const response = await apiServices.fetchInventories();
            setInventories(response.data);
        } catch (error) {
            console.error('Error fetching inventories:', error);
        }
    };

    const fetchHistory = async (itemId: number) => {
        try {
            const response = await apiServices.fetchRestockHistory(itemId);
            setHistory(response.data);
            setSelectedItemId(itemId);
            setShowHistoryModal(true);
        } catch (error) {
            console.error('Error fetching history:', error);
        }
    };

    const createItem = async () => {
        try {
            await apiServices.createItem({
                inventory_id: Number(newItem.inventory_id),
                name: newItem.name,
                description: newItem.description,
                quantity: Number(newItem.quantity),
            });
            setNewItem({ inventory_id: '', name: '', description: '', quantity: '' });
            setShowCreateModal(false);
            fetchItems();
        } catch (error) {
            console.error('Error creating item:', error);
        }
    };

    const restockItem = async (id: number, amount: number) => {
        try {
            await apiServices.restockItem(id, amount);
            setShowRestockModal(false);
            setRestockAmount('');
            setRestockAttempts(prev => ({ ...prev, [id]: 0 }));
            setErrorMessage(null);
            fetchItems();
        } catch (error: any) {
            const attempts = (restockAttempts[id] || 0) + 1;
            setRestockAttempts(prev => ({ ...prev, [id]: attempts }));

            if (attempts > 3 || error.response?.data?.error) {
                setShowRestockModal(false);
                setErrorMessage(
                    error.response?.data?.error || 'Failed to restock item after multiple attempts'
                );
                setRestockAmount('');
            }
            console.error('Error restocking item:', error);
        }
    };

    const deleteItem = async (id: number) => {
        try {
            await apiServices.deleteItem(id);
            setItems(items.filter(item => item.id !== id));
        } catch (error) {
            console.error('Error deleting item:', error);
        }
    };

    const handleSignout = async () => {
        try {
            await apiServices.signout();
            navigate('/login');
        } catch (error) {
            console.error('Error signing out:', error);
        }
    };

    const openRestockModal = (itemId: number) => {
        setSelectedItemId(itemId);
        setShowRestockModal(true);
        setErrorMessage(null);
    };

    return (
        <div className="p-6">
            <h1 className="text-center text-xl font-bold">Welcome to Item Inventory Management!</h1>
            <div className="flex justify-between mb-6 items-center">
                <h1 className="text-3xl font-bold">
                    {user ? `Hello, ${user.Email}` : 'Inventory Items'}
                </h1>
                <div className="flex gap-4">
                    <button
                        className="bg-blue-500 text-white p-2 rounded"
                        onClick={() => setShowCreateModal(true)}
                    >
                        Add Item
                    </button>
                    <button className="bg-red-500 text-white p-2 rounded" onClick={handleSignout}>
                        Sign Out
                    </button>
                </div>
            </div>

            {errorMessage && (
                <div className="text-red-500 text-center mb-4">{errorMessage}</div>
            )}
            {/* Items Table */}
            <table className="w-full bg-white shadow-md rounded">
                <thead>
                <tr className="bg-gray-200">
                    <th className="p-3 text-center">Item Name</th>
                    <th className="p-3 text-center">Description</th>
                    <th className="p-3 text-center">Inventory Name</th>
                    <th className="p-3 text-center">Inventory Desc</th>
                    <th className="p-3 text-center">Quantity</th>
                    <th className="p-3 text-center">Created At</th>
                    <th className="p-3 text-center">Actions</th>
                </tr>
                </thead>
                <tbody>
                {items.map(item => (
                    <tr key={item.id} className="border-b">
                        <td className="p-3 text-center">{item.name}</td>
                        <td className="p-3 text-center">{item.description}</td>
                        <td className="p-3 text-center">{item.inventory.name}</td>
                        <td className="p-3 text-center">{item.inventory.description}</td>
                        <td className="p-3 text-center">{item.quantity}</td>
                        <td className="p-3 text-center">{new Date(item.created_at).toLocaleString()}</td>
                        <td className="p-3 flex gap-2 justify-center">
                            <button
                                className="bg-yellow-500 text-white p-1 rounded"
                                onClick={() => openRestockModal(item.id)}
                            >
                                Restock
                            </button>
                            <button
                                className="bg-red-500 text-white p-1 rounded"
                                onClick={() => deleteItem(item.id)}
                            >
                                Delete
                            </button>
                            <button
                                className="bg-green-500 text-white p-1 rounded"
                                onClick={() => fetchHistory(item.id)}
                            >
                                History
                            </button>
                        </td>
                    </tr>
                ))}
                </tbody>
            </table>
            {/* Create Item Modal */}
            {showCreateModal && (
                <div className="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center">
                    <div className="bg-white p-6 rounded shadow-lg w-1/3">
                        <h2 className="text-xl font-bold mb-4 text-center">Create New Item</h2>
                        <select
                            className="w-full border p-2 rounded mb-4"
                            value={newItem.inventory_id}
                            onChange={e => setNewItem({ ...newItem, inventory_id: e.target.value })}
                        >
                            <option value="">Select Inventory</option>
                            {inventories.map(inventory => (
                                <option key={inventory.id} value={inventory.id}>
                                    {inventory.name}
                                </option>
                            ))}
                        </select>
                        <input
                            className="w-full border p-2 rounded mb-4"
                            placeholder="Name"
                            value={newItem.name}
                            onChange={e => setNewItem({ ...newItem, name: e.target.value })}
                        />
                        <input
                            className="w-full border p-2 rounded mb-4"
                            placeholder="Description"
                            value={newItem.description}
                            onChange={e => setNewItem({ ...newItem, description: e.target.value })}
                        />
                        <input
                            type="number"
                            className="w-full border p-2 rounded mb-4"
                            placeholder="Quantity"
                            value={newItem.quantity}
                            onChange={e => setNewItem({ ...newItem, quantity: e.target.value })}
                        />
                        <div className="flex gap-4 justify-center">
                            <button
                                className="bg-blue-500 text-white p-2 rounded"
                                onClick={createItem}
                                disabled={!newItem.inventory_id || !newItem.name || !newItem.quantity || Number(newItem.quantity) < 0}
                            >
                                Create
                            </button>
                            <button
                                className="bg-gray-500 text-white p-2 rounded"
                                onClick={() => {
                                    setShowCreateModal(false);
                                    setNewItem({ inventory_id: '', name: '', description: '', quantity: '' });
                                }}
                            >
                                Cancel
                            </button>
                        </div>
                    </div>
                </div>
            )}

            {/* Restock Modal */}
            {showRestockModal && (
                <div className="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center">
                    <div className="bg-white p-6 rounded shadow-lg w-1/3">
                        <h2 className="text-xl font-bold mb-4 text-center">Restock Item {selectedItemId}</h2>
                        <input
                            type="number"
                            className="w-full border p-2 rounded mb-4"
                            placeholder="Enter restock amount"
                            value={restockAmount}
                            onChange={e => setRestockAmount(e.target.value)}
                        />
                        <div className="flex gap-4 justify-center">
                            <button
                                className="bg-blue-500 text-white p-2 rounded"
                                onClick={() => restockItem(selectedItemId!, Number(restockAmount))}
                                disabled={!restockAmount || Number(restockAmount) <= 0}
                            >
                                Restock
                            </button>
                            <button
                                className="bg-gray-500 text-white p-2 rounded"
                                onClick={() => {
                                    setShowRestockModal(false);
                                    setRestockAmount('');
                                    setErrorMessage(null);
                                }}
                            >
                                Cancel
                            </button>
                        </div>
                    </div>
                </div>
            )}

            {/* Restock History Modal */}
            {showHistoryModal && (
                <div className="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center">
                    <div className="bg-white p-6 rounded shadow-lg w-1/2">
                        <h2 className="text-xl font-bold mb-4 text-center">Restock History for Item {selectedItemId}</h2>
                        <table className="w-full">
                            <thead>
                            <tr className="bg-gray-200">
                                <th className="p-2 text-center">Restock Amount</th>
                                <th className="p-2 text-center">Timestamp</th>
                            </tr>
                            </thead>
                            <tbody>
                            {history.map(entry => (
                                <tr key={entry.id} className="border-b">
                                    <td className="p-2 text-center">{entry.restock_amount}</td>
                                    <td className="p-2 text-center">{new Date(entry.restock_timestamp).toLocaleString()}</td>
                                </tr>
                            ))}
                            </tbody>
                        </table>
                        <button
                            className="mt-4 bg-gray-500 text-white p-2 rounded mx-auto block"
                            onClick={() => setShowHistoryModal(false)}
                        >
                            Close
                        </button>
                    </div>
                </div>
            )}
        </div>
    );
};

export default Items;