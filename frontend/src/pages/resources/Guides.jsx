import { Link } from 'react-router-dom';
import { ArrowLeft, BookOpen, Clock } from 'lucide-react';
import Navbar from '../../components/Navbar';

export default function Guides() {
  const guides = [
    {
      id: 1,
      title: "Getting Started with FlowBoard",
      description: "Everything you need to know to start using FlowBoard effectively",
      duration: "10 min read",
      level: "Beginner"
    },
    {
      id: 2,
      title: "Advanced Board Management",
      description: "Master advanced features like automation, custom fields, and more",
      duration: "15 min read",
      level: "Advanced"
    },
    {
      id: 3,
      title: "Team Collaboration Best Practices",
      description: "Learn how to work effectively with your team using FlowBoard",
      duration: "12 min read",
      level: "Intermediate"
    },
    {
      id: 4,
      title: "Keyboard Shortcuts Guide",
      description: "Speed up your workflow with these essential keyboard shortcuts",
      duration: "5 min read",
      level: "All Levels"
    }
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
          <h1 className="text-4xl font-bold text-gray-900 mb-4">FlowBoard Guides</h1>
          <p className="text-xl text-gray-600">
            Step-by-step guides to help you master FlowBoard
          </p>
        </div>

        <div className="grid md:grid-cols-2 gap-6">
          {guides.map((guide) => (
            <div key={guide.id} className="bg-white rounded-lg shadow-md p-6 hover:shadow-xl transition">
              <div className="flex items-start justify-between mb-4">
                <BookOpen className="w-8 h-8 text-blue-600" />
                <span className="bg-blue-100 text-blue-700 text-xs font-medium px-3 py-1 rounded-full">
                  {guide.level}
                </span>
              </div>
              <h2 className="text-2xl font-bold text-gray-900 mb-2">{guide.title}</h2>
              <p className="text-gray-600 mb-4">{guide.description}</p>
              <div className="flex items-center text-sm text-gray-500">
                <Clock className="w-4 h-4 mr-2" />
                <span>{guide.duration}</span>
              </div>
              <button className="mt-4 text-blue-600 hover:text-blue-700 font-medium">
                Read guide →
              </button>
            </div>
          ))}
        </div>
      </div>
    </div>
  );
}