import { useState, useEffect } from 'react';
import { X, MessageSquare, Tag, Paperclip, Calendar, Trash2, Upload } from 'lucide-react';
import { 
  getCard, 
  getComments, 
  createComment, 
  deleteComment,
  uploadAttachment,
  deleteAttachment,
  downloadAttachment,
  getLabels,
  addLabelToCard,
  removeLabelFromCard
} from '../../services/api';
import { format } from 'date-fns';

export default function CardModal({ cardId, onClose, onUpdate }) {
  const [card, setCard] = useState(null);
  const [comments, setComments] = useState([]);
  const [newComment, setNewComment] = useState('');
  const [loading, setLoading] = useState(true);
  const [uploading, setUploading] = useState(false);
  const [labels, setLabels] = useState([]);
  const [showLabels, setShowLabels] = useState(false);

  useEffect(() => {
    if (cardId) {
      fetchCardDetails();
      fetchComments();
    }
  }, [cardId]);

  const fetchCardDetails = async () => {
    try {
      setLoading(true);
      const response = await getCard(cardId);
      setCard(response.data);
      
      // Fetch labels after we have the card
      if (response.data.board_id) {
        fetchLabels(response.data.board_id);
      }
    } catch (error) {
      console.error('Failed to fetch card:', error);
      alert('Failed to load card details');
    } finally {
      setLoading(false);
    }
  };

  const fetchComments = async () => {
    try {
      const response = await getComments(cardId);
      setComments(response.data.comments || []);
    } catch (error) {
      console.error('Failed to fetch comments:', error);
    }
  };

  const fetchLabels = async (boardId) => {
    try {
      const response = await getLabels(boardId);
      setLabels(response.data.labels || []);
    } catch (error) {
      console.error('Failed to fetch labels:', error);
    }
  };

  const handleAddComment = async (e) => {
    e.preventDefault();
    if (!newComment.trim()) return;

    setLoading(true);
    try {
      await createComment(cardId, newComment);
      setNewComment('');
      fetchComments();
    } catch (error) {
      console.error('Failed to add comment:', error);
      alert('Failed to add comment');
    } finally {
      setLoading(false);
    }
  };

  const handleDeleteComment = async (commentId) => {
    if (!window.confirm('Delete this comment?')) return;

    try {
      await deleteComment(commentId);
      fetchComments();
    } catch (error) {
      console.error('Failed to delete comment:', error);
      alert('Failed to delete comment');
    }
  };

  const handleFileUpload = async (e) => {
    const file = e.target.files[0];
    if (!file) return;

    // Check file size (10MB limit)
    if (file.size > 10 * 1024 * 1024) {
      alert('File size must be less than 10MB');
      return;
    }

    setUploading(true);
    try {
      await uploadAttachment(cardId, file);
      fetchCardDetails();
      if (onUpdate) onUpdate(); // Refresh board view
      alert('File uploaded successfully!');
    } catch (error) {
      console.error('Failed to upload file:', error);
      alert('Failed to upload file: ' + (error.response?.data?.error || error.message));
    } finally {
      setUploading(false);
    }
  };

  const handleDeleteAttachment = async (attachmentId) => {
    if (!window.confirm('Delete this attachment?')) return;

    try {
      await deleteAttachment(attachmentId);
      fetchCardDetails();
      if (onUpdate) onUpdate();
    } catch (error) {
      console.error('Failed to delete attachment:', error);
      alert('Failed to delete attachment');
    }
  };

  const handleDownloadAttachment = async (attachmentId, filename) => {
    try {
      const response = await downloadAttachment(attachmentId);
      const url = window.URL.createObjectURL(new Blob([response.data]));
      const link = document.createElement('a');
      link.href = url;
      link.setAttribute('download', filename);
      document.body.appendChild(link);
      link.click();
      link.remove();
      window.URL.revokeObjectURL(url);
    } catch (error) {
      console.error('Failed to download attachment:', error);
      alert('Failed to download attachment');
    }
  };

  const handleToggleLabel = async (labelId) => {
    const hasLabel = card.labels?.some(l => l.id === labelId);
    
    try {
      if (hasLabel) {
        await removeLabelFromCard(cardId, labelId);
      } else {
        await addLabelToCard(cardId, labelId);
      }
      fetchCardDetails();
      if (onUpdate) onUpdate();
    } catch (error) {
      console.error('Failed to toggle label:', error);
      alert('Failed to update label');
    }
  };

  // Loading state
  if (loading || !card) {
    return (
      <div className="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50">
        <div className="bg-white rounded-lg p-8">
          <div className="text-gray-600">Loading card details...</div>
        </div>
      </div>
    );
  }

  return (
    <div className="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50 p-4 animate-fadeIn">
      <div className="bg-white rounded-lg w-full max-w-3xl max-h-[90vh] overflow-y-auto shadow-2xl animate-slideUp">
        {/* Header */}
        <div className="sticky top-0 bg-white border-b p-4 flex justify-between items-start">
          <div className="flex-1">
            <h2 className="text-2xl font-bold text-gray-800 mb-2">{card.title}</h2>
            <p className="text-sm text-gray-600">in list {card.list?.title || 'Unknown'}</p>
          </div>
          <button
            onClick={onClose}
            className="text-gray-500 hover:text-gray-700 p-2 hover:bg-gray-100 rounded transition"
          >
            <X className="w-6 h-6" />
          </button>
        </div>

        <div className="p-6 space-y-6">
          {/* Description */}
          {card.description && (
            <div>
              <h3 className="font-semibold text-gray-800 mb-2">Description</h3>
              <p className="text-gray-700 whitespace-pre-wrap">{card.description}</p>
            </div>
          )}

          {/* Due Date */}
          {card.due_date && (
            <div className="flex items-center gap-2 text-sm">
              <Calendar className="w-4 h-4 text-gray-600" />
              <span className="font-medium">Due:</span>
              <span className="bg-yellow-100 px-3 py-1 rounded">
                {format(new Date(card.due_date), 'MMM d, yyyy h:mm a')}
              </span>
            </div>
          )}

          {/* Labels */}
          <div>
            <div className="flex items-center justify-between mb-2">
              <h3 className="font-semibold text-gray-800 flex items-center gap-2">
                <Tag className="w-4 h-4" />
                Labels
              </h3>
              <button
                onClick={() => setShowLabels(!showLabels)}
                className="text-sm text-blue-600 hover:underline"
              >
                {showLabels ? 'Hide' : 'Edit'}
              </button>
            </div>

            {card.labels?.length > 0 && (
              <div className="flex flex-wrap gap-2 mb-2">
                {card.labels.map((label) => (
                  <span
                    key={label.id}
                    className="px-3 py-1 rounded text-white font-medium text-sm"
                    style={{ backgroundColor: label.color }}
                  >
                    {label.name}
                  </span>
                ))}
              </div>
            )}

            {showLabels && (
              <div className="bg-gray-50 p-3 rounded-lg space-y-2 animate-fadeIn">
                {labels.length > 0 ? (
                  labels.map((label) => {
                    const isActive = card.labels?.some(l => l.id === label.id);
                    return (
                      <button
                        key={label.id}
                        onClick={() => handleToggleLabel(label.id)}
                        className={`w-full text-left px-3 py-2 rounded font-medium text-white transition ${
                          isActive ? 'ring-2 ring-blue-500' : ''
                        }`}
                        style={{ backgroundColor: label.color }}
                      >
                        {label.name} {isActive && '✓'}
                      </button>
                    );
                  })
                ) : (
                  <p className="text-gray-500 text-sm">No labels available</p>
                )}
              </div>
            )}
          </div>

          {/* Attachments */}
          <div>
            <h3 className="font-semibold text-gray-800 mb-3 flex items-center gap-2">
              <Paperclip className="w-4 h-4" />
              Attachments
            </h3>

            {card.attachments && card.attachments.length > 0 && (
              <div className="space-y-2 mb-3">
                {card.attachments.map((attachment) => (
                  <div
                    key={attachment.id}
                    className="flex items-center justify-between bg-gray-50 p-3 rounded hover:bg-gray-100 transition"
                  >
                    <div className="flex items-center gap-3">
                      <Paperclip className="w-4 h-4 text-gray-600" />
                      <div>
                        <p className="font-medium text-sm">{attachment.filename}</p>
                        <p className="text-xs text-gray-600">
                          {(attachment.file_size / 1024).toFixed(1)} KB
                        </p>
                      </div>
                    </div>
                    <div className="flex gap-2">
                      <button
                        onClick={() => handleDownloadAttachment(attachment.id, attachment.filename)}
                        className="text-blue-600 hover:underline text-sm"
                      >
                        Download
                      </button>
                      <button
                        onClick={() => handleDeleteAttachment(attachment.id)}
                        className="text-red-600 hover:underline text-sm"
                      >
                        Delete
                      </button>
                    </div>
                  </div>
                ))}
              </div>
            )}

            <label className="flex items-center justify-center gap-2 bg-blue-50 hover:bg-blue-100 text-blue-600 px-4 py-3 rounded-lg cursor-pointer transition border-2 border-dashed border-blue-300">
              <Upload className="w-5 h-5" />
              <span className="font-medium">
                {uploading ? 'Uploading...' : 'Choose File to Upload'}
              </span>
              <input
                type="file"
                onChange={handleFileUpload}
                disabled={uploading}
                className="hidden"
              />
            </label>
            <p className="text-xs text-gray-500 mt-1">Max file size: 10MB</p>
          </div>

          {/* Comments */}
          <div>
            <h3 className="font-semibold text-gray-800 mb-3 flex items-center gap-2">
              <MessageSquare className="w-4 h-4" />
              Comments
            </h3>

            {/* Add Comment */}
            <form onSubmit={handleAddComment} className="mb-4">
              <textarea
                value={newComment}
                onChange={(e) => setNewComment(e.target.value)}
                placeholder="Write a comment..."
                className="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent resize-none"
                rows={3}
              />
              <button
                type="submit"
                disabled={loading || !newComment.trim()}
                className="mt-2 px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 transition disabled:opacity-50 disabled:cursor-not-allowed"
              >
                {loading ? 'Posting...' : 'Post Comment'}
              </button>
            </form>

            {/* Comments List */}
            <div className="space-y-3">
              {comments.map((comment) => (
                <div
                  key={comment.id}
                  className="bg-gray-50 p-3 rounded-lg animate-fadeIn"
                >
                  <div className="flex justify-between items-start mb-2">
                    <div>
                      <span className="font-medium text-sm">{comment.user?.username || 'Unknown'}</span>
                      <span className="text-xs text-gray-600 ml-2">
                        {format(new Date(comment.created_at), 'MMM d, h:mm a')}
                      </span>
                    </div>
                    <button
                      onClick={() => handleDeleteComment(comment.id)}
                      className="text-gray-500 hover:text-red-500 p-1"
                    >
                      <Trash2 className="w-4 h-4" />
                    </button>
                  </div>
                  <p className="text-gray-700 text-sm">{comment.content}</p>
                </div>
              ))}

              {comments.length === 0 && (
                <p className="text-gray-500 text-sm text-center py-4">
                  No comments yet. Be the first to comment!
                </p>
              )}
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}