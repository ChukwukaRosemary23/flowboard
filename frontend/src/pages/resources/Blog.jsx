import { Link } from 'react-router-dom';
import { ArrowLeft, Calendar, User } from 'lucide-react';
import Navbar from '../../components/Navbar';

export default function Blog() {
  const posts = [
    {
      id: 1,
      title: "10 Ways to Boost Team Productivity with FlowBoard",
      excerpt: "Discover proven strategies to help your team work more efficiently using FlowBoard's powerful features.",
      date: "Oct 25, 2024",
      author: "FlowBoard Team",
      image: "https://images.unsplash.com/photo-1552664730-d307ca884978?w=800"
    },
    {
      id: 2,
      title: "How to Organize Your Projects Like a Pro",
      excerpt: "Learn the best practices for structuring boards, lists, and cards to keep your projects organized.",
      date: "Oct 20, 2024",
      author: "FlowBoard Team",
      image: "https://images.unsplash.com/photo-1454165804606-c3d57bc86b40?w=800"
    },
    {
      id: 3,
      title: "Remote Team Collaboration Made Easy",
      excerpt: "Tips and tricks for managing remote teams effectively using FlowBoard's real-time features.",
      date: "Oct 15, 2024",
      author: "FlowBoard Team",
      image: "https://images.unsplash.com/photo-1522071820081-009f0129c71c?w=800"
    }
  ];

  return (
    <div className="min-h-screen bg-gray-50">
      <Navbar />
      
      <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-12">
        {/* Header */}
        <div className="mb-8">
          <Link to="/" className="inline-flex items-center text-blue-600 hover:text-blue-700 mb-4">
            <ArrowLeft className="w-4 h-4 mr-2" />
            Back to Home
          </Link>
          <h1 className="text-4xl font-bold text-gray-900 mb-4">FlowBoard Blog</h1>
          <p className="text-xl text-gray-600">
            Tips, updates, and insights to help you get the most out of FlowBoard
          </p>
        </div>

        {/* Blog Posts Grid */}
        <div className="grid md:grid-cols-2 lg:grid-cols-3 gap-8">
          {posts.map((post) => (
            <div key={post.id} className="bg-white rounded-lg shadow-md overflow-hidden hover:shadow-xl transition">
              <img src={post.image} alt={post.title} className="w-full h-48 object-cover" />
              <div className="p-6">
                <div className="flex items-center text-sm text-gray-600 mb-3">
                  <Calendar className="w-4 h-4 mr-2" />
                  <span>{post.date}</span>
                  <span className="mx-2">•</span>
                  <User className="w-4 h-4 mr-2" />
                  <span>{post.author}</span>
                </div>
                <h2 className="text-xl font-bold text-gray-900 mb-2">{post.title}</h2>
                <p className="text-gray-600 mb-4">{post.excerpt}</p>
                <button className="text-blue-600 hover:text-blue-700 font-medium">
                  Read more →
                </button>
              </div>
            </div>
          ))}
        </div>
      </div>
    </div>
  );
}