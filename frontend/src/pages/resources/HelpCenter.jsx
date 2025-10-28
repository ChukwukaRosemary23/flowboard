import { Link } from 'react-router-dom';
import { ArrowLeft, Search, HelpCircle, Book, MessageCircle, Mail } from 'lucide-react';
import Navbar from '../../components/Navbar';
import { useState } from 'react';

export default function HelpCenter() {
  const [searchQuery, setSearchQuery] = useState('');

  const categories = [
    {
      icon: <Book className="w-8 h-8" />,
      title: "Getting Started",
      description: "Learn the basics of FlowBoard",
      articles: 12
    },
    {
      icon: <HelpCircle className="w-8 h-8" />,
      title: "Common Questions",
      description: "Answers to frequently asked questions",
      articles: 25
    },
    {
      icon: <MessageCircle className="w-8 h-8" />,
      title: "Troubleshooting",
      description: "Fix common issues",
      articles: 18
    },
    {
      icon: <Mail className="w-8 h-8" />,
      title: "Account & Billing",
      description: "Manage your account and subscription",
      articles: 15
    }
  ];

  const popularArticles = [
    "How do I create a new board?",
    "How to invite team members",
    "Understanding card labels and tags",
    "Setting up automation rules",
    "Exporting and backing up data"
  ];

  return (
    <div className="min-h-screen bg-gray-50">
      <Navbar />
      
      <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-12">
        <div className="mb-8">
          <Link to="/" className="inline-flex items-center text-blue-600 hover:text-blue-700 mb-4">
            <ArrowLeft className="w-4 h-4 mr-2" />
            Back to Home
          </Link>
          <h1 className="text-4xl font-bold text-gray-900 mb-4">Help Center</h1>
          <p className="text-xl text-gray-600 mb-8">
            Find answers and get support
          </p>

          {/* Search Bar */}
          <div className="relative max-w-2xl">
            <Search className="absolute left-4 top-1/2 transform -translate-y-1/2 text-gray-400 w-5 h-5" />
            <input
              type="text"
              placeholder="Search for help..."
              value={searchQuery}
              onChange={(e) => setSearchQuery(e.target.value)}
              className="w-full pl-12 pr-4 py-4 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent text-lg"
            />
          </div>
        </div>

        {/* Categories */}
        <div className="grid md:grid-cols-2 lg:grid-cols-4 gap-6 mb-12">
          {categories.map((category, index) => (
            <div key={index} className="bg-white rounded-lg shadow-md p-6 hover:shadow-xl transition cursor-pointer">
              <div className="text-blue-600 mb-4">{category.icon}</div>
              <h3 className="text-xl font-bold text-gray-900 mb-2">{category.title}</h3>
              <p className="text-gray-600 text-sm mb-3">{category.description}</p>
              <span className="text-blue-600 text-sm font-medium">{category.articles} articles</span>
            </div>
          ))}
        </div>

        {/* Popular Articles */}
        <div className="bg-white rounded-lg shadow-md p-8">
          <h2 className="text-2xl font-bold text-gray-900 mb-6">Popular Articles</h2>
          <ul className="space-y-4">
            {popularArticles.map((article, index) => (
              <li key={index}>
                <a href="#" className="flex items-center text-gray-700 hover:text-blue-600 transition">
                  <HelpCircle className="w-5 h-5 mr-3 text-gray-400" />
                  <span className="text-lg">{article}</span>
                </a>
              </li>
            ))}
          </ul>
        </div>

        {/* Contact Support */}
        <div className="mt-12 bg-gradient-to-r from-blue-600 to-purple-600 rounded-lg p-8 text-center text-white">
          <h2 className="text-2xl font-bold mb-4">Still need help?</h2>
          <p className="mb-6 text-lg">Our support team is here to assist you</p>
          <button className="bg-white text-blue-600 px-8 py-3 rounded-lg font-medium hover:bg-gray-100 transition">
            Contact Support
          </button>
        </div>
      </div>
    </div>
  );
}