import { useState, useEffect } from 'react';
import { useParams, useNavigate, Link } from 'react-router-dom';
import { DndContext, DragOverlay, closestCorners, PointerSensor, useSensor, useSensors } from '@dnd-kit/core';
import { SortableContext, verticalListSortingStrategy } from '@dnd-kit/sortable';
import { getBoard, getLists, createList, deleteList, getCards, createCard, updateCard, moveCard } from '../services/api';
import { Plus, ArrowLeft, Trash2, X } from 'lucide-react';
import SortableCard from '../components/board/SortableCard';
import CardModal from '../components/card/CardModal';
import { useWebSocket } from '../hooks/useWebSocket';

function BoardView() {
  const { id } = useParams();
  const navigate = useNavigate();
  const [board, setBoard] = useState(null);
  const [lists, setLists] = useState([]);
  const [cards, setCards] = useState({}); // ✅ Already correct!
  const [loading, setLoading] = useState(true);
  const [newListName, setNewListName] = useState('');
  const [newCardTitles, setNewCardTitles] = useState({});
  const [showNewListInput, setShowNewListInput] = useState(false);
  const [activeCard, setActiveCard] = useState(null);
  const [selectedCardId, setSelectedCardId] = useState(null);

  // WebSocket for real-time updates
  const { isConnected, sendMessage } = useWebSocket(id, (message) => {
    console.log('WebSocket message:', message);
    loadBoardData();
  });

  const sensors = useSensors(
    useSensor(PointerSensor, {
      activationConstraint: {
        distance: 8,
      },
    })
  );

  useEffect(() => {
    if (!id || id === 'undefined') {
      console.error('Invalid board ID');
      navigate('/dashboard');
      return;
    }
    loadBoardData();
  }, [id]);

  const loadBoardData = async () => {
    if (!id || id === 'undefined') {
      console.error('Cannot load board: Invalid ID');
      return;
    }

    try {
      setLoading(true);
      const [boardRes, listsRes] = await Promise.all([
        getBoard(id),
        getLists(id)
      ]);

      // Handle board data
      if (boardRes.data) {
        setBoard(boardRes.data);
      }

      // Handle lists data - response is {"count": 0, "lists": []}
      let listsData = [];
      
      if (Array.isArray(listsRes.data)) {
        listsData = listsRes.data;
      } else if (Array.isArray(listsRes.data?.lists)) {
        listsData = listsRes.data.lists;
      } else if (Array.isArray(listsRes.data?.data)) {
        listsData = listsRes.data.data;
      }

      setLists(listsData);

      // Load cards for each list
      const cardsData = {};
      await Promise.all(
        listsData.map(async (list) => {
          try {
            const listId = list.id || list.list_id;
            const cardsRes = await getCards(listId);
            
            // ✅ FIX: Handle cards response properly
            let cardsList = [];
            if (Array.isArray(cardsRes.data)) {
              cardsList = cardsRes.data;
            } else if (Array.isArray(cardsRes.data?.cards)) {
              cardsList = cardsRes.data.cards;
            } else if (cardsRes.data?.data && Array.isArray(cardsRes.data.data)) {
              cardsList = cardsRes.data.data;
            }
            
            cardsData[listId] = cardsList;
          } catch (err) {
            console.error(`Error loading cards for list ${list.id || list.list_id}:`, err);
            const listId = list.id || list.list_id;
            cardsData[listId] = []; // ✅ Always set to empty array on error
          }
        })
      );
      
      setCards(cardsData);
    } catch (err) {
      console.error('Error loading board:', err);
      alert('Failed to load board: ' + (err.response?.data?.error || err.message));
      navigate('/dashboard');
    } finally {
      setLoading(false);
    }
  };

  const handleCreateList = async (e) => {
    e.preventDefault();
    if (!newListName.trim()) return;

    try {
      const payload = {
        board_id: parseInt(id),
        title: newListName,
        position: lists.length
      };
      
      await createList(payload);
      setNewListName('');
      setShowNewListInput(false);
      loadBoardData();
    } catch (err) {
      console.error('Error creating list:', err);
      alert('Failed to create list: ' + (err.response?.data?.error || err.message));
    }
  };

  const handleDeleteList = async (listId, listName) => {
    if (!window.confirm(`Delete "${listName}"? All cards in this list will be deleted.`)) return;

    try {
      await deleteList(listId);
      loadBoardData();
    } catch (err) {
      console.error('Error deleting list:', err);
      alert('Failed to delete list: ' + (err.response?.data?.error || err.message));
    }
  };

  const handleCreateCard = async (listId) => {
    const title = newCardTitles[listId];
    if (!title?.trim()) return;

    try {
      await createCard({
        list_id: listId,
        title: title,
        position: (cards[listId] || []).length
      });
      setNewCardTitles({ ...newCardTitles, [listId]: '' });
      loadBoardData();
    } catch (err) {
      console.error('Error creating card:', err);
      alert('Failed to create card: ' + (err.response?.data?.error || err.message));
    }
  };

  const handleDragStart = (event) => {
    const { active } = event;
    // ✅ FIX: Find card properly
    let foundCard = null;
    for (const listCards of Object.values(cards)) {
      if (Array.isArray(listCards)) {
        foundCard = listCards.find((c) => c.card_id === active.id || c.id === active.id);
        if (foundCard) break;
      }
    }
    setActiveCard(foundCard);
  };

  const handleDragEnd = async (event) => {
    const { active, over } = event;
    setActiveCard(null);

    if (!over) return;

    const activeCardId = active.id;
    const overId = over.id;

    // Find source list
    let sourceListId = null;
    for (const [listId, listCards] of Object.entries(cards)) {
      if (Array.isArray(listCards) && listCards.find((c) => (c.card_id === activeCardId || c.id === activeCardId))) {
        sourceListId = parseInt(listId);
        break;
      }
    }

    // Determine target list
    let targetListId = null;
    const list = lists.find((l) => (l.id || l.list_id) === overId);
    if (list) {
      targetListId = overId;
    } else {
      for (const [listId, listCards] of Object.entries(cards)) {
        if (Array.isArray(listCards) && listCards.find((c) => (c.card_id === overId || c.id === overId))) {
          targetListId = parseInt(listId);
          break;
        }
      }
    }

    if (!targetListId) return;

    try {
      await moveCard(activeCardId, {
        list_id: targetListId,
        position: 0
      });
      loadBoardData();
    } catch (err) {
      console.error('Error moving card:', err);
      alert('Failed to move card');
    }
  };

  if (loading) {
    return (
      <div className="min-h-screen bg-gradient-to-br from-blue-500 to-purple-600 flex items-center justify-center">
        <div className="text-white text-xl">Loading board...</div>
      </div>
    );
  }

  return (
    <div className="min-h-screen bg-gradient-to-br from-blue-500 to-purple-600">
      {/* Header */}
      <header className="bg-black bg-opacity-20 backdrop-blur-sm">
        <div className="max-w-full px-4 py-4 flex justify-between items-center">
          <div className="flex items-center space-x-4">
            <Link
              to="/dashboard"
              className="text-white hover:bg-white hover:bg-opacity-20 p-2 rounded transition"
            >
              <ArrowLeft className="w-5 h-5" />
            </Link>
            <h1 className="text-2xl font-bold text-white">
              {board?.name || board?.title || 'Board'}
            </h1>
            {isConnected && (
              <span className="flex items-center space-x-2 text-green-300 text-sm">
                <span className="w-2 h-2 bg-green-400 rounded-full animate-pulse"></span>
                <span>Live</span>
              </span>
            )}
          </div>
        </div>
      </header>

      {/* Board Content */}
      <div className="p-4 overflow-x-auto">
        <DndContext
          sensors={sensors}
          collisionDetection={closestCorners}
          onDragStart={handleDragStart}
          onDragEnd={handleDragEnd}
        >
          <div className="flex space-x-4 pb-4">
            {/* Lists */}
            {lists.map((list) => {
              const listId = list.id || list.list_id;
              const listName = list.title || list.name;
              const listCards = cards[listId] || []; // ✅ Always get array
              
              return (
                <div
                  key={listId}
                  className="bg-gray-100 rounded-lg p-4 w-80 flex-shrink-0"
                >
                  {/* List Header */}
                  <div className="flex justify-between items-center mb-4">
                    <h2 className="font-semibold text-gray-900">{listName}</h2>
                    <button
                      onClick={() => handleDeleteList(listId, listName)}
                      className="text-gray-500 hover:text-red-600 transition"
                    >
                      <Trash2 className="w-4 h-4" />
                    </button>
                  </div>

                  {/* Cards */}
                  <SortableContext
                    items={listCards.map((c) => c.card_id || c.id)}
                    strategy={verticalListSortingStrategy}
                  >
                    <div className="space-y-2 mb-3">
                      {listCards.map((card) => (
                        <SortableCard
                          key={card.card_id || card.id}
                          card={card}
                          onClick={() => setSelectedCardId(card.card_id || card.id)}
                        />
                      ))}
                    </div>
                  </SortableContext>

                  {/* Add Card */}
                  <div>
                    <input
                      type="text"
                      placeholder="+ Add a card"
                      value={newCardTitles[listId] || ''}
                      onChange={(e) =>
                        setNewCardTitles({
                          ...newCardTitles,
                          [listId]: e.target.value
                        })
                      }
                      onKeyPress={(e) => {
                        if (e.key === 'Enter') {
                          handleCreateCard(listId);
                        }
                      }}
                      className="w-full px-3 py-2 bg-white border border-gray-300 rounded-lg text-sm focus:outline-none focus:ring-2 focus:ring-blue-500"
                    />
                  </div>
                </div>
              );
            })}

            {/* Add List */}
            <div className="w-80 flex-shrink-0">
              {showNewListInput ? (
                <div className="bg-gray-100 rounded-lg p-4">
                  <form onSubmit={handleCreateList}>
                    <input
                      type="text"
                      placeholder="Enter list name..."
                      value={newListName}
                      onChange={(e) => setNewListName(e.target.value)}
                      className="w-full px-3 py-2 border border-gray-300 rounded-lg mb-2 focus:outline-none focus:ring-2 focus:ring-blue-500"
                      autoFocus
                    />
                    <div className="flex gap-2">
                      <button
                        type="submit"
                        className="flex-1 px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 transition text-sm font-medium"
                      >
                        Add List
                      </button>
                      <button
                        type="button"
                        onClick={() => {
                          setShowNewListInput(false);
                          setNewListName('');
                        }}
                        className="px-3 py-2 text-gray-600 hover:bg-gray-200 rounded-lg transition"
                      >
                        <X className="w-5 h-5" />
                      </button>
                    </div>
                  </form>
                </div>
              ) : (
                <button
                  onClick={() => setShowNewListInput(true)}
                  className="w-full bg-white bg-opacity-20 hover:bg-opacity-30 text-white rounded-lg p-4 flex items-center justify-center space-x-2 transition"
                >
                  <Plus className="w-5 h-5" />
                  <span className="font-medium">Add another list</span>
                </button>
              )}
            </div>
          </div>

          <DragOverlay>
            {activeCard ? (
              <div className="bg-white rounded-lg shadow-xl p-3 cursor-grabbing opacity-90 rotate-3">
                <p className="text-gray-900 font-medium">{activeCard.title}</p>
              </div>
            ) : null}
          </DragOverlay>
        </DndContext>
      </div>

      {/* Card Modal */}
      {selectedCardId && (
        <CardModal
          cardId={selectedCardId}
          onClose={() => setSelectedCardId(null)}
          onUpdate={loadBoardData}
        />
      )}
    </div>
  );
}

export default BoardView;