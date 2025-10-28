import { useState, useEffect } from 'react';
import { Link, useNavigate } from 'react-router-dom';
import { useAuth } from '../contexts/AuthContext';
import { getBoards, createBoard, deleteBoard } from '../services/api';
import { Plus, Trash2, LogOut } from 'lucide-react';

function Dashboard() {
  const [boards, setBoards] = useState([]);
  const [loading, setLoading] = useState(true);
  const [showCreateModal, setShowCreateModal] = useState(false);
  const [newBoardName, setNewBoardName] = useState('');
  const [creating, setCreating] = useState(false);
  const { user, logout } = useAuth();
  const navigate = useNavigate();

  useEffect(() => {
    loadBoards();
  }, []);

  const loadBoards = async () => {
    try {
      setLoading(true);
      const response = await getBoards();
      console.log('API Response:', response); // Debug log
      
      // FIXED: Your Go API returns {boards: [...], count: N}
      const boardsData = response.data;
      
      // Ensure it's always an array
      if (boardsData && Array.isArray(boardsData.boards)) {
        console.log('Boards loaded:', boardsData.boards);
        setBoards(boardsData.boards);
      } else if (Array.isArray(boardsData)) {
        console.log('Boards array:', boardsData);
        setBoards(boardsData);
      } else {
        console.warn('Unexpected response format:', boardsData);
        setBoards([]);
      }
    } catch (err) {
      console.error('Error loading boards:', err);
      setBoards([]); // Set to empty array on error to prevent .map() errors
    } finally {
      setLoading(false);
    }
  };

  const handleCreateBoard = async (e) => {
    e.preventDefault();
    if (!newBoardName.trim()) return;

    try {
      setCreating(true);
      console.log('Sending board data:', { name: newBoardName }); // Debug log
      const response = await createBoard({ name: newBoardName });
      console.log('Board created successfully:', response); // Debug log
      setNewBoardName('');
      setShowCreateModal(false);
      loadBoards();
    } catch (err) {
      console.error('Error creating board:', err);
      console.error('Error response:', err.response?.data); // Debug log
      alert(`Failed to create board: ${err.response?.data?.error || err.message}`);
    } finally {
      setCreating(false);
    }
  };

  const handleDeleteBoard = async (boardId, boardName) => {
    if (!window.confirm(`Delete "${boardName}"? This cannot be undone.`)) return;

    try {
      await deleteBoard(boardId);
      loadBoards();
    } catch (err) {
      console.error('Error deleting board:', err);
      alert('Failed to delete board');
    }
  };

  const handleLogout = () => {
    logout();
    navigate('/login');
  };

  if (loading) {
    return (
      <div className="min-h-screen bg-gradient-to-br from-blue-50 to-purple-50 flex items-center justify-center">
        <div className="text-xl text-gray-600">Loading...</div>
      </div>
    );
  }

  return (
    <div className="min-h-screen bg-gradient-to-br from-blue-50 to-purple-50">
      {/* Header */}
      <header className="bg-white shadow-sm">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-4">
          <div className="flex justify-between items-center">
            <div className="flex items-center space-x-4">
              <Link to="/" className="flex items-center space-x-2">
                <div className="bg-gradient-to-r from-blue-600 to-purple-600 p-2 rounded-lg">
                  <svg className="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M9 5H7a2 2 0 00-2 2v12a2 2 0 002 2h10a2 2 0 002-2V7a2 2 0 00-2-2h-2M9 5a2 2 0 002 2h2a2 2 0 002-2M9 5a2 2 0 012-2h2a2 2 0 012 2" />
                  </svg>
                </div>
                <span className="text-2xl font-bold text-gray-900">FlowBoard</span>
              </Link>
            </div>
            <div className="flex items-center space-x-4">
              <span className="text-gray-700">Hello, <span className="font-semibold">{user?.username}</span>!</span>
              <button
                onClick={handleLogout}
                className="flex items-center space-x-2 px-4 py-2 text-gray-700 hover:bg-gray-100 rounded-lg transition"
              >
                <LogOut className="w-4 h-4" />
                <span>Logout</span>
              </button>
            </div>
          </div>
        </div>
      </header>

      {/* Main Content */}
      <main className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
        <div className="mb-8">
          <h1 className="text-3xl font-bold text-gray-900 mb-2">Your Boards</h1>
          <p className="text-gray-600">Create and manage your project boards</p>
        </div>

        {/* Boards Grid */}
        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 gap-6">
          {/* Create New Board Card */}
          <button
            onClick={() => setShowCreateModal(true)}
            className="bg-gradient-to-br from-blue-100 to-purple-100 rounded-xl p-6 hover:from-blue-200 hover:to-purple-200 transition-all duration-300 transform hover:scale-105 border-2 border-dashed border-blue-300 hover:border-blue-400 min-h-[150px] flex flex-col items-center justify-center"
          >
            <Plus className="w-12 h-12 text-blue-600 mb-2" />
            <span className="text-blue-600 font-semibold">Create New Board</span>
          </button>

          {/* Existing Boards - FIXED: Check multiple possible ID field names */}
          {Array.isArray(boards) && boards.map((board) => {
            // Handle different possible ID field names from Go backend
            const boardId = board.board_id || board.id || board.ID;
            const boardName = board.name || board.title || board.Name || board.Title;
            const createdAt = board.created_at || board.createdAt || board.CreatedAt;
            
            console.log('Board:', { boardId, boardName, board }); // Debug log
            
            return (
              <div
                key={boardId}
                className="bg-white rounded-xl shadow-md hover:shadow-xl transition-all duration-300 transform hover:-translate-y-1 overflow-hidden group"
              >
                <Link
                  to={`/board/${boardId}`}
                  className="block p-6 min-h-[150px] flex flex-col justify-between"
                >
                  <h3 className="text-xl font-bold text-gray-900 mb-2 group-hover:text-blue-600 transition">
                    {boardName}
                  </h3>
                  <p className="text-sm text-gray-500">
                    Created {createdAt ? new Date(createdAt).toLocaleDateString() : 'Unknown'}
                  </p>
                </Link>
                <div className="bg-gray-50 px-6 py-3 flex justify-end border-t border-gray-100">
                  <button
                    onClick={() => handleDeleteBoard(boardId, boardName)}
                    className="text-red-600 hover:text-red-700 hover:bg-red-50 p-2 rounded transition"
                    title="Delete board"
                  >
                    <Trash2 className="w-4 h-4" />
                  </button>
                </div>
              </div>
            );
          })}
        </div>

        {boards.length === 0 && (
          <div className="text-center py-16">
            <div className="text-gray-400 mb-4">
              <svg className="w-24 h-24 mx-auto" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={1.5} d="M9 5H7a2 2 0 00-2 2v12a2 2 0 002 2h10a2 2 0 002-2V7a2 2 0 00-2-2h-2M9 5a2 2 0 002 2h2a2 2 0 002-2M9 5a2 2 0 012-2h2a2 2 0 012 2" />
              </svg>
            </div>
            <h3 className="text-xl font-semibold text-gray-700 mb-2">No boards yet</h3>
            <p className="text-gray-500 mb-6">Create your first board to get started!</p>
          </div>
        )}
      </main>

      {/* Create Board Modal */}
      {showCreateModal && (
        <div className="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50 p-4">
          <div className="bg-white rounded-2xl p-8 max-w-md w-full shadow-2xl transform transition-all">
            <h2 className="text-2xl font-bold text-gray-900 mb-6">Create New Board</h2>
            <form onSubmit={handleCreateBoard}>
              <div className="mb-6">
                <label className="block text-sm font-medium text-gray-700 mb-2">
                  Board Name
                </label>
                <input
                  type="text"
                  value={newBoardName}
                  onChange={(e) => setNewBoardName(e.target.value)}
                  placeholder="e.g., Project Management"
                  className="w-full px-4 py-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent"
                  autoFocus
                  required
                />
              </div>
              <div className="flex gap-3">
                <button
                  type="button"
                  onClick={() => {
                    setShowCreateModal(false);
                    setNewBoardName('');
                  }}
                  className="flex-1 px-4 py-3 border border-gray-300 text-gray-700 rounded-lg hover:bg-gray-50 transition font-medium"
                >
                  Cancel
                </button>
                <button
                  type="submit"
                  disabled={creating || !newBoardName.trim()}
                  className="flex-1 px-4 py-3 bg-gradient-to-r from-blue-600 to-purple-600 text-white rounded-lg hover:from-blue-700 hover:to-purple-700 transition font-medium disabled:opacity-50 disabled:cursor-not-allowed"
                >
                  {creating ? 'Creating...' : 'Create Board'}
                </button>
              </div>
            </form>
          </div>
        </div>
      )}
    </div>
  );
}

export default Dashboard;