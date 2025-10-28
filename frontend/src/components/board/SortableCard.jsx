import { useSortable } from '@dnd-kit/sortable';
import { CSS } from '@dnd-kit/utilities';
import { Calendar, MessageSquare, Paperclip, Trash2 } from 'lucide-react';
import { format } from 'date-fns';

export default function SortableCard({ card, onClick, onDelete }) {
  const {
    attributes,
    listeners,
    setNodeRef,
    transform,
    transition,
    isDragging,
  } = useSortable({ id: card.card_id }); // ← Changed from card.id to card.card_id

  const style = {
    transform: CSS.Transform.toString(transform),
    transition,
    opacity: isDragging ? 0.5 : 1,
  };

  return (
    <div
      ref={setNodeRef}
      style={style}
      {...attributes}
      {...listeners}
      onClick={onClick}
      className="bg-white rounded-lg shadow hover:shadow-md transition-all duration-200 p-3 cursor-pointer group relative transform hover:scale-102"
    >
      <h4 className="text-sm font-medium text-gray-800 mb-2">{card.title}</h4>

      {/* Badges */}
      <div className="flex flex-wrap gap-2 text-xs text-gray-600">
        {card.due_date && (
          <div className="flex items-center gap-1 bg-yellow-100 px-2 py-1 rounded">
            <Calendar className="w-3 h-3" />
            {format(new Date(card.due_date), 'MMM d')}
          </div>
        )}
        
        {card.comment_count > 0 && ( // ← Changed from card.comments?.length
          <div className="flex items-center gap-1 bg-gray-100 px-2 py-1 rounded">
            <MessageSquare className="w-3 h-3" />
            {card.comment_count}
          </div>
        )}

        {card.attachments?.length > 0 && (
          <div className="flex items-center gap-1 bg-gray-100 px-2 py-1 rounded">
            <Paperclip className="w-3 h-3" />
            {card.attachments.length}
          </div>
        )}
      </div>

      {/* Labels */}
      {card.labels?.length > 0 && (
        <div className="flex flex-wrap gap-1 mt-2">
          {card.labels.map((label) => (
            <span
              key={label.label_id} // ← Changed from label.id
              className="px-2 py-1 rounded text-xs text-white font-medium"
              style={{ backgroundColor: label.color }}
            >
              {label.name}
            </span>
          ))}
        </div>
      )}

      {/* Delete Button */}
      {onDelete && (
        <button
          onClick={(e) => {
            e.stopPropagation();
            onDelete();
          }}
          className="absolute top-2 right-2 opacity-0 group-hover:opacity-100 transition-opacity bg-red-500 text-white p-1 rounded hover:bg-red-600"
        >
          <Trash2 className="w-3 h-3" />
        </button>
      )}
    </div>
  );
}