import { Link } from 'react-router-dom';
import { ArrowLeft, Target, BarChart, Calendar, Users } from 'lucide-react';
import Navbar from '../../components/Navbar';

export default function MarketingTeams() {
  return (
    <div className="min-h-screen bg-gray-50">
      <Navbar />
      
      <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-12">
        <Link to="/" className="inline-flex items-center text-blue-600 hover:text-blue-700 mb-8">
          <ArrowLeft className="w-4 h-4 mr-2" />
          Back to Home
        </Link>

        {/* Hero Section */}
        <div className="text-center mb-16">
          <h1 className="text-5xl font-bold text-gray-900 mb-6">
            FlowBoard for Marketing Teams
          </h1>
          <p className="text-xl text-gray-600 max-w-3xl mx-auto">
            Plan campaigns, manage content, and collaborate seamlessly with your marketing team
          </p>
        </div>

        {/* Features Grid */}
        <div className="grid md:grid-cols-2 lg:grid-cols-4 gap-8 mb-16">
          <div className="bg-white rounded-lg shadow-md p-6 text-center">
            <div className="bg-blue-100 w-16 h-16 rounded-full flex items-center justify-center mx-auto mb-4">
              <Target className="w-8 h-8 text-blue-600" />
            </div>
            <h3 className="text-lg font-bold text-gray-900 mb-2">Campaign Planning</h3>
            <p className="text-gray-600">Organize and track all your marketing campaigns in one place</p>
          </div>

          <div className="bg-white rounded-lg shadow-md p-6 text-center">
            <div className="bg-purple-100 w-16 h-16 rounded-full flex items-center justify-center mx-auto mb-4">
              <BarChart className="w-8 h-8 text-purple-600" />
            </div>
            <h3 className="text-lg font-bold text-gray-900 mb-2">Performance Tracking</h3>
            <p className="text-gray-600">Monitor campaign metrics and ROI with custom dashboards</p>
          </div>

          <div className="bg-white rounded-lg shadow-md p-6 text-center">
            <div className="bg-green-100 w-16 h-16 rounded-full flex items-center justify-center mx-auto mb-4">
              <Calendar className="w-8 h-8 text-green-600" />
            </div>
            <h3 className="text-lg font-bold text-gray-900 mb-2">Content Calendar</h3>
            <p className="text-gray-600">Plan and schedule content across all your channels</p>
          </div>

          <div className="bg-white rounded-lg shadow-md p-6 text-center">
            <div className="bg-orange-100 w-16 h-16 rounded-full flex items-center justify-center mx-auto mb-4">
              <Users className="w-8 h-8 text-orange-600" />
            </div>
            <h3 className="text-lg font-bold text-gray-900 mb-2">Team Collaboration</h3>
            <p className="text-gray-600">Work together with designers, writers, and stakeholders</p>
          </div>
        </div>

        {/* Use Case */}
        <div className="bg-white rounded-lg shadow-xl p-12">
          <h2 className="text-3xl font-bold text-gray-900 mb-8 text-center">How Marketing Teams Use FlowBoard</h2>
          <div className="space-y-8">
            <div className="flex items-start">
              <div className="bg-blue-100 rounded-full p-3 mr-6 flex-shrink-0">
                <span className="text-2xl font-bold text-blue-600">1</span>
              </div>
              <div>
                <h3 className="text-xl font-bold text-gray-900 mb-2">Campaign Launch Board</h3>
                <p className="text-gray-600">Create a board for each campaign with lists for ideation, design, approval, and launch phases.</p>
              </div>
            </div>

            <div className="flex items-start">
              <div className="bg-purple-100 rounded-full p-3 mr-6 flex-shrink-0">
                <span className="text-2xl font-bold text-purple-600">2</span>
              </div>
              <div>
                <h3 className="text-xl font-bold text-gray-900 mb-2">Content Pipeline</h3>
                <p className="text-gray-600">Manage blog posts, social media, and email campaigns with due dates and assignments.</p>
              </div>
            </div>

            <div className="flex items-start">
              <div className="bg-green-100 rounded-full p-3 mr-6 flex-shrink-0">
                <span className="text-2xl font-bold text-green-600">3</span>
              </div>
              <div>
                <h3 className="text-xl font-bold text-gray-900 mb-2">Asset Management</h3>
                <p className="text-gray-600">Attach design files, documents, and assets directly to cards for easy access.</p>
              </div>
            </div>
          </div>
        </div>

        {/* CTA */}
        <div className="mt-16 bg-gradient-to-r from-blue-600 to-purple-600 rounded-lg p-12 text-center text-white">
          <h2 className="text-3xl font-bold mb-4">Ready to transform your marketing workflow?</h2>
          <p className="text-xl mb-8">Join thousands of marketing teams using FlowBoard</p>
          <Link
            to="/register"
            className="inline-block bg-white text-blue-600 px-8 py-4 rounded-lg font-bold text-lg hover:bg-gray-100 transition"
          >
            Start Free Trial
          </Link>
        </div>
      </div>
    </div>
  );
}